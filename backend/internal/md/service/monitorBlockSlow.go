package service

import (
	"fmt"

	"github.com/centrifuge/go-substrate-rpc-client/v3/types"
	"github.com/pkg/errors"
)

func (svc *MDService) MonitorBlockSlow() error {
	// // 拿Best區塊資料
	// signedBlock, err := svc.subgameApi.RPC.Chain.GetBlockLatest()
	// if err != nil {
	// 	return errors.WithStack(err)
	// }
	// chainBlockNumber := int64(signedBlock.Block.Header.Number)

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
	finalizedBlockNum := int64(finalizedBlock.Block.Header.Number)

	record, err := svc.MDRepo.SubGameGetLastBlockLog()
	if err != nil {
		return errors.WithStack(err)
	}
	if record == nil {
		return nil
	}
	dbBlcokNumber := int64(record.Num)

	blockNumAway := abs(finalizedBlockNum - dbBlcokNumber)
	if blockNumAway >= 50 {
		svc.notifySvc.NotifySlackError(fmt.Sprintf(
			"SubGame Module 區塊監聽落後\n鏈上區塊：%d\n資料庫區塊：%d\n相差塊數：%d",
			finalizedBlockNum,
			dbBlcokNumber,
			blockNumAway))
	}

	return nil
}

func abs(input int64) int64 {
	if input < 0 {
		return -input
	}
	return input
}
