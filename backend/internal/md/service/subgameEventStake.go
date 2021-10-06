package service

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

func (svc *MDService) SubGameEventStake(event *domain.SubGameChainEvent, gas decimal.Decimal, blockTimestamp int) error {

	// (AccountId, ProgramId, PalletId, u64, u64, u64, NftId, BalanceOf)
	paramsByte, _ := json.Marshal(event.Params)
	var paramsData []map[string]interface{}
	json.Unmarshal(paramsByte, &paramsData)

	fromPubkey := fmt.Sprintf("%v", paramsData[0]["value"])
	fromPubkey = strings.TrimPrefix(fromPubkey, "0x")

	programIdStr := fmt.Sprintf("%v", paramsData[1]["value"])
	programId, err := strconv.ParseUint(programIdStr, 10, 64)
	if err != nil {
		return errors.WithStack(err)
	}

	palletIdStr := fmt.Sprintf("%v", paramsData[2]["value"])
	palletId, err := strconv.ParseUint(palletIdStr, 10, 64)
	if err != nil {
		return errors.WithStack(err)
	}

	periodOfUseDayStr := fmt.Sprintf("%v", paramsData[3]["value"])
	periodOfUseDay, err := strconv.ParseInt(periodOfUseDayStr, 10, 64)
	if err != nil {
		return errors.WithStack(err)
	}

	startTimestampStr := fmt.Sprintf("%v", paramsData[4]["value"])
	startTimestamp, err := strconv.ParseInt(startTimestampStr[:10], 10, 64)
	if err != nil {
		return errors.WithStack(err)
	}

	endTimestampStr := fmt.Sprintf("%v", paramsData[5]["value"])
	endTimestamp, err := strconv.ParseInt(endTimestampStr[:10], 10, 64)
	if err != nil {
		return errors.WithStack(err)
	}

	NftHash := fmt.Sprintf("%v", paramsData[6]["value"])

	amountStr := fmt.Sprintf("%v", paramsData[7]["value"])
	amount, err := decimal.NewFromString(amountStr)
	if err != nil {
		return errors.WithStack(err)
	}
	amount = SubGameDotToUnit(amount)

	var TxStatus uint8 = 1
	if strings.EqualFold(event.EventId, "ExtrinsicFailed") == true {
		TxStatus = 0
	}

	// 除去小數位
	fee := SubGameDotToUnit(gas)

	doneAt := time.Unix(int64(blockTimestamp), 0)

	// check user exist
	user, err := svc.MDRepo.GetUserByAddress(fromPubkey)
	if err != nil {
		return errors.WithStack(err)
	}

	if user == nil {
		return errors.New("user not found")
	}

	startTime := time.Unix(startTimestamp, 0)
	endTime := time.Unix(endTimestamp, 0)

	// 新增提領記錄
	err = svc.MDRepo.StakeInsertRecord(domain.StakeRecord{
		UserId:           user.ID,
		UserName:         user.NickName,
		ModuleId:         palletId,
		ProgramId:        programId,
		StakeSGB:         amount,
		PeriodOfUseMonth: int(periodOfUseDay),
		StartTime:        startTime,
		EndTime:          endTime,
		NFTHash:          NftHash,
		Address:          fromPubkey,
		TxHash:           event.ExtrinsicHash,
		BlockNum:         event.BlockNum,
		Fee:              fee,
		TxStatus:         TxStatus,
		DoneAt:           &doneAt,
		Nonce:            "",
	})
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
