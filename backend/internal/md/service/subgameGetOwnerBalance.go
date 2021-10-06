package service

import (
	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/itering/substrate-api-rpc/rpc"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

func (svc *MDService) SubGameGetOwnerBalance() (decimal.Decimal, error) {
	var balance decimal.Decimal

	accountRaw, err := rpc.ReadStorage(nil, "system", "account", "", svc.config.SubGame.MDOwnerPubkey)
	if err != nil {
		return balance, errors.WithStack(err)
	}
	accountData := new(domain.SubGameAccountDecode)
	accountRaw.ToAny(accountData)

	balance = accountData.Data.Free.Add(accountData.Data.Reserved)
	balance = SubGameDotToUnit(balance)

	return balance, nil
}
