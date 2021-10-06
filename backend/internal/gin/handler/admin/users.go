package admin

import (
	"bytes"
	"io/ioutil"

	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/SubGame-Network/SubGameModuleService/internal/gin/handler"
	SubGameHelp "github.com/SubGame-Network/SubGameModuleService/internal/md/service"
	"github.com/SubGame-Network/SubGameModuleService/internal/provider"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type UserRequest struct {
	Row     int    `form:"row"`
	Page    int    `form:"page"`
	Account string `form:"userAccount"`
	Email   string `form:"email"`
	Address string `form:"address"`
}
type UsersResponse struct {
	Count int64               `json:"count"`
	List  []UsersResponseItem `json:"list"`
}

type UsersResponseItem struct {
	ID               uint64                        `json:"id"`
	Account          string                        `json:"userAccount"`
	Email            string                        `json:"email"`
	Address          string                        `json:"address"`
	UnBorrowMt       decimal.Decimal               `json:"unBorrowMt"`       // 未借貸Mt
	BorrowMt         decimal.Decimal               `json:"borrowMt"`         // 借貸Mt
	BenefitBalance   decimal.Decimal               `json:"benefitBalance"`   // 未提領SGB獎勵餘額
	TotalProfit      decimal.Decimal               `json:"totalProfit"`      // 累計SGB獎勵
	ActiveStatus     domain.UserActiveStatusType   `json:"activeStatus"`     // 啟動1, 凍結0 (前台借貸功能disable)
	WithdrawStatus   domain.UserWithdrawStatusType `json:"withdrawStatus"`   // 可提領1, 禁止提領0
	ProfitPercentage decimal.Decimal               `json:"profitPercentage"` // 收益比（總收益/(未借貸Mt + 借貸Mt)）
}

func Users(c *gin.Context) {
	var req UserRequest
	reqBody, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(reqBody))
	if err := c.BindQuery(&req); err != nil {
		zap.S().Infof("err: %v, req: %v", err, string(reqBody))
		handler.Failed(c, domain.ErrorBadRequest, err.Error())
		return
	}

	repo, err := provider.NewMDRepo()
	if err != nil {
		zap.S().Warn("", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}
	if req.Address != "" {
		// address to pubkey
		pubkey, err := SubGameHelp.SubGameAddressToPubkey(req.Address)
		if err == nil {
			req.Address = pubkey
		}
	}

	users, count, err := repo.GetAllUserByPage(req.Row, req.Page, req.Account, req.Email, req.Address)
	if err != nil {
		zap.S().Warn("", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}

	output := UsersResponse{
		Count: count,
		List:  []UsersResponseItem{},
	}

	config := provider.NewConfig()
	for _, data := range users {
		// 借貸總額
		totalMt := data.UnBorrowMt.Add(data.BorrowMt)
		profitPercentage := domain.GetProfitPercentage(data.TotalProfit, totalMt)

		address, _ := SubGameHelp.SubGamePubkeyToAddress(config, data.Address)
		output.List = append(output.List, UsersResponseItem{
			ID:               data.ID,
			Account:          data.Account,
			Email:            data.Email,
			Address:          address,
			UnBorrowMt:       data.UnBorrowMt,
			BorrowMt:         data.BorrowMt,
			BenefitBalance:   data.BenefitBalance,
			TotalProfit:      data.TotalProfit,
			ActiveStatus:     data.ActiveStatus,
			WithdrawStatus:   data.WithdrawStatus,
			ProfitPercentage: profitPercentage.Mul(decimal.NewFromInt(100)),
		})
	}

	handler.Success(c, output)
}
