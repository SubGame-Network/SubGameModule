package service

import (
	"github.com/itering/substrate-api-rpc/rpc"
	"github.com/shopspring/decimal"
)

func (svc *MDService) SubGameGetExtrinsicFee(encodeExtrinsic string) (fee decimal.Decimal, err error) {
	paymentInfo, err := rpc.GetPaymentQueryInfo(svc.subgameWsConn, encodeExtrinsic)
	if paymentInfo != nil {
		return paymentInfo.PartialFee, nil
	}
	return decimal.Zero, err
}
