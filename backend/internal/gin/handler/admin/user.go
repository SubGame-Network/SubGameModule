package admin

import (
	"strconv"
	"time"

	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/SubGame-Network/SubGameModuleService/internal/gin/handler"
	SubGameHelp "github.com/SubGame-Network/SubGameModuleService/internal/md/service"
	"github.com/SubGame-Network/SubGameModuleService/internal/provider"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type UserResponse struct {
	Basic          UserResponseItem             `json:"basic"`
	BorrowRecord   []*BorrowResponseItem        `json:"borrowRecord"`
	ProfitRecord   []*ProfitResponseItem        `json:"profitRecord"`
	WithdrawRecord []*WithdrawResponseItem      `json:"withdrawRecord"`
	UserEditRecord []*UserEditAddressRecordList `json:"userEditAddressRecord"`
}

type UserResponseItem struct {
	Id               uint64                        `json:"id"`
	Account          string                        `json:"userAccount"`
	Email            string                        `json:"email"`
	VerifyingEmail   string                        `json:"VerifyingEmail"`
	Address          string                        `json:"address"`
	UnBorrowMt       decimal.Decimal               `json:"unBorrowMt"`       // 未借貸Mt
	BorrowMt         decimal.Decimal               `json:"borrowMt"`         // 借貸Mt
	BenefitBalance   decimal.Decimal               `json:"benefitBalance"`   // 未提領SGB獎勵餘額
	TotalProfit      decimal.Decimal               `json:"totalProfit"`      // 累計SGB獎勵
	ActiveStatus     domain.UserActiveStatusType   `json:"activeStatus"`     // 啟動1, 凍結0 (前台借貸功能disable)
	WithdrawStatus   domain.UserWithdrawStatusType `json:"withdrawStatus"`   // 可提領1, 禁止提領0
	ProfitPercentage decimal.Decimal               `json:"profitPercentage"` // 收益比（總收益/(未借貸Mt + 借貸Mt)）
	IsCanReset       bool                          `json:"isCanReset"`       // 可以重設user
	IsCanEditData    bool                          `json:"isCanEditData"`
	IsHaveTotp       bool                          `json:"isHaveTotp"`
	IpAddress        string                        `json:"ipAddress"`
}
type BorrowResponseItem struct {
	Account     string          `json:"userAccount"`
	CreatedAt   string          `json:"createdTime"`
	BorrowMt    decimal.Decimal `json:"borrowMt"`
	TotalProfit decimal.Decimal `json:"profitUSDT"`
	DonedAt     *string         `json:"completedTime"`
}
type ProfitResponseItem struct {
	CreatedTime       string `json:"createdTime"`
	BorrowMt          string `json:"borrowMt"`
	Rate              string `json:"rate"`
	CoinRate          string `json:"coinRate"`
	BeforeTotalProfit string `json:"beforeUSDT"`
	AfterTotalProfit  string `json:"afterUSDT"`
	ProfitUSDT        string `json:"profitUSDT"`
	ProfitSGB         string `json:"profitSGB"`
}
type WithdrawResponseItem struct {
	BeforeBalance string `json:"beforeBalance"`
	AfterBalance  string `json:"afterBalance"`
	Amount        string `json:"amount"`
	CreatedTime   string `json:"createdTime"`
	CompletedTime string `json:"completedTime"`
	Hash          string `json:"hash"`
	Status        uint8  `json:"status"`
}

type UserEditAddressRecordList struct {
	Account       string `json:"userAccount"`
	BeforeAddress string `json:"beforeAddress"`
	AfterAddress  string `json:"afterAddress"`
	CreatedAt     string `json:"createdTime"`
}

func User(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	repo, err := provider.NewMDRepo()
	if err != nil {
		zap.S().Warn("", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}
	user, err := repo.GetUserById(id)
	if err != nil {
		zap.S().Warn("", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}

	if user == nil {
		handler.Failed(c, domain.ErrorUserNotFound, "")
		return
	}

	redis, err := provider.NewRedis()
	if err != nil {
		zap.S().Warnw("", "err", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}

	key := domain.GetRedisKeyAdminResetUserEmailEmail(user.Account)
	verifyingEmail, _ := redis.Get(key)

	// 借貸總額
	totalMt := user.UnBorrowMt.Add(user.BorrowMt)

	profitPercentage := domain.GetProfitPercentage(user.TotalProfit, totalMt)

	config := provider.NewConfig()
	address, _ := SubGameHelp.SubGamePubkeyToAddress(config, user.Address)

	isHaveTotp := user.TotpSecret != "" && user.IsOkTotpSecret

	isCanEditData := true
	w, err := repo.WithdrawRecordByAccountAndTime(user.Account, time.Time{}, time.Time{})
	if err != nil {
		zap.S().Warnw("", "err", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}
	b, err := repo.GetBorrowRecordsByAccount(user.Account)
	if err != nil {
		zap.S().Warnw("", "err", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}
	if len(w) > 0 || len(b) > 0 {
		isCanEditData = false
	}

	// 可以重設
	isCanReset := false
	userResetRecord, _ := repo.GetUserResetRecordByUserId(user.ID)
	if len(userResetRecord) == 0 && isCanEditData != true { // 若可以修改 則不顯示重設鈕
		isCanReset = true
	}

	userResponseItem := UserResponseItem{
		Id:               user.ID,
		Account:          user.Account,
		Email:            user.Email,
		VerifyingEmail:   verifyingEmail,
		Address:          address,
		UnBorrowMt:       user.UnBorrowMt,
		BorrowMt:         user.BorrowMt,
		BenefitBalance:   user.BenefitBalance,
		IsCanReset:       isCanReset,
		IsCanEditData:    isCanEditData,
		TotalProfit:      user.TotalProfit,
		ActiveStatus:     user.ActiveStatus,
		WithdrawStatus:   user.WithdrawStatus,
		ProfitPercentage: profitPercentage.Mul(decimal.NewFromInt(100)),
		IsHaveTotp:       isHaveTotp,
		IpAddress:        user.LastIpAddress,
	}

	// To do借貸紀錄
	records, err := repo.GetBorrowRecordsByAccount(user.Account)
	if err != nil {
		zap.S().Warnw("", "err", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}
	borrowRecord := []*BorrowResponseItem{}
	for _, data := range records {
		var doneAt *string
		if data.DonedAt != nil {
			t := data.DonedAt.Format(time.RFC3339)
			doneAt = &t
		}
		borrowRecord = append(borrowRecord, &BorrowResponseItem{
			Account:     data.Account,
			CreatedAt:   data.CreatedAt.Format(time.RFC3339),
			BorrowMt:    data.BorrowMt,
			TotalProfit: data.TotalProfit,
			DonedAt:     doneAt,
		})
	}

	// To do收益紀錄
	pRecord, err := repo.ProfitRecordByAccountAndTime(user.Account, time.Time{}, time.Time{})
	if err != nil {
		zap.S().Warnw("", "err", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}
	profitRecord := []*ProfitResponseItem{}
	for _, data := range pRecord {
		profitRecord = append(profitRecord, &ProfitResponseItem{
			CreatedTime:       data.CreatedAt.Format(time.RFC3339),
			BorrowMt:          data.BorrowMt.String(),
			Rate:              data.Rate.String(),
			CoinRate:          data.CoinRate.String(),
			BeforeTotalProfit: data.BeforeTotalProfit.String(),
			AfterTotalProfit:  data.AfterTotalProfit.String(),
			ProfitUSDT:        data.ProfitUSDT.String(),
			ProfitSGB:         data.ProfitSGB.String(),
		})
	}

	// To do提領紀錄
	wRecord, err := repo.WithdrawRecordByAccountAndTime(user.Account, time.Time{}, time.Time{})
	if err != nil {
		zap.S().Warnw("", "err", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}
	withdrawRecord := []*WithdrawResponseItem{}
	for _, data := range wRecord {
		completedTime := ""
		if data.DoneAt != nil {
			completedTime = data.DoneAt.Format(time.RFC3339)
		}
		withdrawRecord = append(withdrawRecord, &WithdrawResponseItem{
			BeforeBalance: data.BeforeBenefitBalance.String(),
			AfterBalance:  data.AfterBenefitBalance.String(),
			Amount:        data.Amount.String(),
			CreatedTime:   data.CreatedAt.Format(time.RFC3339),
			CompletedTime: completedTime,
			Hash:          data.TxHash,
			Status:        data.TxStatus,
		})
	}

	editRecords, err := repo.GetUserEditRecordByAccount(user.Account, "", "")
	if err != nil {
		zap.S().Warnw("", "err", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}

	editRecord := []*UserEditAddressRecordList{}
	for _, val := range editRecords {
		beforeAddress, _ := SubGameHelp.SubGamePubkeyToAddress(config, val.BeforeAddress)
		afterAddress, _ := SubGameHelp.SubGamePubkeyToAddress(config, val.AfterAddress)
		editRecord = append(editRecord, &UserEditAddressRecordList{
			Account:       val.Account,
			CreatedAt:     val.CreatedAt.Format(time.RFC3339),
			BeforeAddress: beforeAddress,
			AfterAddress:  afterAddress,
		})
	}

	output := UserResponse{
		Basic:          userResponseItem,
		BorrowRecord:   borrowRecord,
		ProfitRecord:   profitRecord,
		WithdrawRecord: withdrawRecord,
		UserEditRecord: editRecord,
	}
	handler.Success(c, output)
}
