package service

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/centrifuge/go-substrate-rpc-client/v3/types"
	"github.com/itering/substrate-api-rpc"
	"github.com/itering/substrate-api-rpc/metadata"
	"github.com/itering/substrate-api-rpc/rpc"
	"github.com/itering/substrate-api-rpc/websocket"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

const NoHashError = "no hash"

// 監聽區塊方法入口
func (svc *MDService) SubGameEventListener() error {
	fmt.Println("=== SubGame Event Listen Start ===")

	if svc.subgameApi == nil {
		return nil
	}

	var blockNum uint64 = 0
	var runCount int64 = 0

	// 從資料庫log找最後block
	blockLog, _ := svc.MDRepo.SubGameGetLastBlockLog()
	if blockLog != nil {
		blockNum = blockLog.Num
		if blockLog.Done == true {
			blockNum++
		}
	}

	// 間隔1秒去掃區塊
	for ; true; <-time.Tick(1 * time.Second) {
		err := svc.SubGameEventRunner(blockNum)
		if err != nil {
			switch err.Error() {
			case NoHashError:
				continue
			default:
				// 紀錄 blcok_log 失敗
				inputBlockLog := &domain.SubGameBlockLog{
					Num:        uint64(blockNum),
					Done:       false,
					ErrorCount: 1,
				}
				svc.MDRepo.SubGameInsertUpdateBlockLog(inputBlockLog)

				// 寫失敗log
				zap.S().Errorw("SubGameEventListenerError", err)
			}
		}

		blockNum++
		runCount++
	}

	return nil
}

// 修復漏聽的區塊
func (svc *MDService) SubGameFixBlockGap() error {
	logs, err := svc.MDRepo.SubGameGetBlockGapLog()
	if err != nil {
		return errors.WithStack(err)
	}

	var fixErr error
	for _, val := range logs {
		// if val.ErrorCount >= 5 {
		// 	continue
		// }

		blockNum := val.Num
		err := svc.SubGameEventRunner(blockNum)
		if err == nil || err.Error() == NoHashError {
			continue
		}

		errorCount := val.ErrorCount + 1

		// 紀錄 blcok_log 失敗
		inputBlockLog := &domain.SubGameBlockLog{
			Num:        uint64(blockNum),
			Done:       false,
			ErrorCount: errorCount,
		}
		svc.MDRepo.SubGameInsertUpdateBlockLog(inputBlockLog)

		// 寫失敗log
		zap.S().Errorw("SubGameEventListenerError", err)
		fixErr = errors.WithStack(err)

		// 發通知
		if errorCount >= 20 {
			notifyText := fmt.Sprintf("SubGame Event Error: %v", err.Error())
			svc.notifySvc.NotifySlackError(notifyText)
		}

		// 出錯就中斷
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return fixErr
}

// 實際執行監聽的方法
func (svc *MDService) SubGameEventRunner(num uint64) error {
	// startTime := time.Now()

	blockNum := int(num)
	jsonRpcResult := &rpc.JsonRpcResult{}

	// Last Finalized Block Hash
	var finalizedHashStr string
	err := svc.subgameApi.Client.Call(&finalizedHashStr, "chain_getFinalizedHead")
	if err != nil {
		return errors.WithStack(err)
	}
	if finalizedHashStr == "" {
		return errors.New(NoHashError)
	}
	// Finalized Block
	finalizedHash, err := types.NewHashFromHexString(finalizedHashStr)
	if err != nil {
		return errors.WithStack(err)
	}
	finalizedBlock, err := svc.subgameApi.RPC.Chain.GetBlock(finalizedHash)
	if err != nil {
		return errors.WithStack(err)
	}
	finalizedBlockNum := int(finalizedBlock.Block.Header.Number)
	if finalizedBlockNum < blockNum {
		return errors.New(NoHashError)
	}

	// Block Hash
	err = websocket.SendWsRequest(svc.subgameWsConn, jsonRpcResult, rpc.ChainGetBlockHash(domain.ChainWsBlockHash, blockNum))
	if err != nil {
		return errors.WithStack(err)
	}
	blockHash, _ := jsonRpcResult.ToString()

	if blockHash == "" {
		return errors.New(NoHashError)
	}

	// Block
	err = websocket.SendWsRequest(svc.subgameWsConn, jsonRpcResult, rpc.ChainGetBlock(domain.ChainWsBlock, blockHash))
	if err != nil {
		return errors.WithStack(err)
	}
	rpcBlock := jsonRpcResult.ToBlock()

	// Runtime
	err = websocket.SendWsRequest(svc.subgameWsConn, jsonRpcResult, rpc.ChainGetRuntimeVersion(domain.ChainWsSpec, blockHash))
	if err != nil {
		return errors.WithStack(err)
	}

	// chain_spec version
	runtimeVersion := jsonRpcResult.ToRuntimeVersion()
	spec := 0
	if runtimeVersion != nil {
		spec = runtimeVersion.SpecVersion
	}

	// Pallet移除，Metadata拿前一個區塊的hash
	preBlockHash := blockHash
	if blockNum > 1 {
		err = websocket.SendWsRequest(svc.subgameWsConn, jsonRpcResult, rpc.ChainGetBlockHash(domain.ChainWsBlockHash, blockNum-1))
		if err != nil {
			return errors.WithStack(err)
		}
		preBlockHash, _ = jsonRpcResult.ToString()
	}
	// Metadata
	metaRaw, err := rpc.GetMetadataByHash(nil, preBlockHash)
	if err != nil {
		return errors.WithStack(err)
	}
	metaRuntimeRaw := &metadata.RuntimeRaw{
		Spec: spec,
		Raw:  metaRaw,
	}
	metadataInstant := metadata.Process(metaRuntimeRaw)

	// validator list
	validatorList := []string{}
	validatorsRaw, _ := rpc.ReadStorage(nil, "Session", "Validators", blockHash)
	for _, addr := range validatorsRaw.ToStringSlice() {
		validatorList = append(validatorList, strings.TrimPrefix(addr, "0x"))
	}

	// Events
	err = websocket.SendWsRequest(svc.subgameWsConn, jsonRpcResult, rpc.StateGetStorage(domain.ChainWsEvent, svc.config.SubGame.EventKey, blockHash))
	if err != nil {
		return errors.WithStack(err)
	}
	events, _ := jsonRpcResult.ToString()

	// Decode Extrinsic
	decodeExtrinsic, err := substrate.DecodeExtrinsic(rpcBlock.Block.Extrinsics, metadataInstant, spec)
	if err != nil {
		// 新增pallet，Metadata拿當下區塊的hash
		metaRaw, err := rpc.GetMetadataByHash(nil, blockHash)
		if err != nil {
			return errors.WithStack(err)
		}
		metaRuntimeRaw := &metadata.RuntimeRaw{
			Spec: spec,
			Raw:  metaRaw,
		}
		metadataInstant := metadata.Process(metaRuntimeRaw)

		decodeExtrinsic, err = substrate.DecodeExtrinsic(rpcBlock.Block.Extrinsics, metadataInstant, spec)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	// Decode Events
	decodeEvents, err := substrate.DecodeEvent(events, metadataInstant, spec)
	if err != nil {
		return errors.WithStack(err)
	}

	// === 資料整理 ===
	// 整理Event
	domainEvents := []*domain.SubGameChainEvent{}
	decodeByte, _ := json.Marshal(decodeEvents)
	json.Unmarshal(decodeByte, &domainEvents)
	var eventCount int
	eventMap := make(map[string][]*domain.SubGameChainEvent)
	for index, event := range domainEvents {
		event.ModuleId = strings.ToLower(event.ModuleId)
		event.EventIndex = fmt.Sprintf("%d-%d", blockNum, index)
		event.BlockNum = blockNum

		extrinsicIndex := fmt.Sprintf("%d-%d", blockNum, event.ExtrinsicIdx)
		eventMap[extrinsicIndex] = append(eventMap[extrinsicIndex], event)

		eventCount++
	}

	// 整理Extrinsic
	extrinsicFeeMap := make(map[string]decimal.Decimal)
	extrinsicHashMap := make(map[string]string)
	var extrinsicsCount int
	var blockTimestamp int

	domainExtrinsic := []*domain.SubGameChainExtrinsic{}
	decodeExtrinsicByte, _ := json.Marshal(decodeExtrinsic)
	json.Unmarshal(decodeExtrinsicByte, &domainExtrinsic)
	for index, extrinsic := range domainExtrinsic {
		extrinsic.CallModule = strings.ToLower(extrinsic.CallModule)
		extrinsic.BlockNum = blockNum
		extrinsic.ExtrinsicIndex = fmt.Sprintf("%d-%d", extrinsic.BlockNum, index)
		extrinsic.IsSigned = extrinsic.Signature != ""
		extrinsic.ExtrinsicHash = SubGameCheckStringToHex(extrinsic.ExtrinsicHash)

		extrinsic.Success = true
		for _, event := range eventMap[extrinsic.ExtrinsicIndex] {
			if strings.EqualFold(event.ModuleId, "system") {
				extrinsic.Success = !strings.EqualFold(event.EventId, "ExtrinsicFailed")
			}
		}

		if tp := svc.SubGameGetExtrinsicTimestamp(extrinsic); tp > 0 {
			blockTimestamp = tp
		}
		extrinsic.BlockTimestamp = blockTimestamp

		if extrinsic.ExtrinsicHash != "" {
			fee, _ := svc.SubGameGetExtrinsicFee(rpcBlock.Block.Extrinsics[index])
			extrinsic.Fee = fee

			extrinsicFeeMap[extrinsic.ExtrinsicIndex] = fee
			extrinsicHashMap[extrinsic.ExtrinsicIndex] = extrinsic.ExtrinsicHash
		}

		extrinsicsCount++
	}

	// 補上Event缺少extrinsic的資料
	for _, event := range domainEvents {
		event.ExtrinsicHash = extrinsicHashMap[fmt.Sprintf("%d-%d", blockNum, event.ExtrinsicIdx)]
	}
	// === 資料整理 ===

	for _, event := range domainEvents {
		moduleId := strings.ToLower(event.ModuleId)
		eventId := strings.ToLower(event.EventId)

		if moduleId != "subgamestakenft" {
			continue
		}

		gas := extrinsicFeeMap[fmt.Sprintf("%d-%d", blockNum, event.ExtrinsicIdx)]

		switch eventId {
		case "stake":
			err := svc.SubGameEventStake(event, gas, blockTimestamp)
			if err != nil {
				return errors.WithStack(err)
			}
		}
	}

	// 紀錄 blcok_log
	inputBlockLog := &domain.SubGameBlockLog{
		Num:  uint64(blockNum),
		Done: true,
	}
	svc.MDRepo.SubGameInsertUpdateBlockLog(inputBlockLog)

	// zap.S().Infow("區塊監聽完成", "BlockNum", blockNum, "Time", time.Now().Sub(startTime).Seconds())

	return nil
}
