package front

import (
	"bytes"
	"io/ioutil"
	"time"

	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/SubGame-Network/SubGameModuleService/internal/gin/handler"
	"github.com/SubGame-Network/SubGameModuleService/internal/provider"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type StakeRecordRequest struct {
	Row              int    `form:"row"`
	Page             int    `form:"page"`
	NftHash          string `form:"nftId"`
	Status           int    `form:"status"`
	PeriodOfUseMonth int    `form:"periodOfUse"`
}

type StakeRecordItem struct {
	ID               uint64
	ModuleName       string
	StakeSGB         decimal.Decimal
	PeriodOfUseMonth int
	StartTime        time.Time
	EndTime          time.Time
	NFTHash          string
	TxHash           string
	TxStatus         uint8
	DoneAt           *time.Time
	CreatedAt        time.Time
}

type StakeRecordResponse struct {
	Count        int64             `json:"count"`
	List         []StakeRecordItem `json:"list"`
	StakedAmount decimal.Decimal   `json:"stakedAmount"`
	Withrawn     decimal.Decimal   `json:"withrawn"`
}

func StakeRecord(c *gin.Context) {
	userAddress := c.GetString("userAddress")

	var req StakeRecordRequest
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

	user, err := repo.GetUserByAddress(userAddress)
	if err != nil {
		zap.S().Warn("", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}
	if user == nil {
		handler.Failed(c, domain.ErrorUserNotFound, "")
		return
	}
	modules, err := repo.GetAllModule()
	if err != nil {
		zap.S().Warn("", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}
	moduleIdToName := map[uint64]string{}

	for _, val := range modules {
		moduleIdToName[val.ID] = val.Name
	}

	// 如果req.status=3  == status0
	var status *int
	if req.Status != 0 {
		status = &req.Status
	}
	if req.Status == 3 {
		req.Status = 0
		status = &req.Status
	}

	record, count, err := repo.StakeAllRecordsByPage(user.ID, req.Row, req.Page, req.NftHash, status, req.PeriodOfUseMonth)
	if err != nil {
		zap.S().Warn("", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}
	var output []StakeRecordItem
	for _, val := range record {
		output = append(output, StakeRecordItem{
			ID:               val.ID,
			ModuleName:       moduleIdToName[val.ModuleId],
			StakeSGB:         val.StakeSGB,
			PeriodOfUseMonth: val.PeriodOfUseMonth,
			StartTime:        val.StartTime,
			EndTime:          val.EndTime,
			NFTHash:          val.NFTHash,
			TxHash:           val.TxHash,
			TxStatus:         val.TxStatus,
			DoneAt:           val.DoneAt,
			CreatedAt:        val.CreatedAt,
		})
	}

	stakedAmount, err := repo.StakeTotalStakedByUserId(user.ID)
	if err != nil {
		zap.S().Warn("", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}

	withrawn, err := repo.StakeTotalWithrawnByUserId(user.ID)
	if err != nil {
		zap.S().Warn("", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}
	handler.Success(c, StakeRecordResponse{
		Count:        count,
		List:         output,
		StakedAmount: stakedAmount,
		Withrawn:     withrawn,
	})
	return
}
