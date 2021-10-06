package service

import (
	"fmt"

	"github.com/pkg/errors"
)

func (svc *MDService) StakeTimeoutNotify() error {
	records, err := svc.MDRepo.StakeGetTimeout()
	if err != nil {
		return errors.WithStack(err)
	}

	for _, val := range records {
		notifyText := fmt.Sprintf("發送交易Pending逾時:")
		notifyText += fmt.Sprintf("\nDB ID: `%v`", val.ID)
		notifyText += fmt.Sprintf("\nUser: `%v`", val.UserName)
		notifyText += fmt.Sprintf("\nAmount: `%v`", val.StakeSGB)
		notifyText += fmt.Sprintf("\nTxHash: `%v`", val.TxHash)
		svc.notifySvc.NotifySlackError(notifyText)

		err := svc.MDRepo.StakeUpdateNotifyTime(val.ID)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
