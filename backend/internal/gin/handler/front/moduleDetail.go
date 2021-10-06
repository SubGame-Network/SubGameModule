package front

import (
	"strconv"

	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/SubGame-Network/SubGameModuleService/internal/gin/handler"
	"github.com/SubGame-Network/SubGameModuleService/internal/provider"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type ModuleDetailResponse struct {
	Module  ModuleItem    `json:"module"`
	Program []ProgramItem `json:"program"`
}
type ModuleItem struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Depiction   string `json:"depiction"`
	ReadmeMdUrl string `json:"readmeMdUrl"`
}
type ProgramItem struct {
	ProgramID   uint64          `json:"programID"`
	PeriodOfUse int             `json:"periodOfUse"`
	Amount      decimal.Decimal `json:"amount"`
	IsCanStake  bool            `json:"isCanStake"`
}

func ModuleDetail(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	repo, err := provider.NewMDRepo()
	if err != nil {
		zap.S().Warn("", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}

	m, err := repo.GetModuleById(id)
	if err != nil {
		zap.S().Warn("", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}
	programs, err := repo.GetAllProgram()
	if err != nil {
		zap.S().Warn("", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}

	program := []ProgramItem{}
	for _, v := range programs {
		program = append(program, ProgramItem{
			ProgramID:   v.ID,
			PeriodOfUse: v.PeriodOfUse,
			Amount:      v.Amount,
			IsCanStake:  true,
		})
	}

	output := ModuleDetailResponse{
		Module: ModuleItem{
			Id:          m.ID,
			Name:        m.Name,
			Depiction:   m.Depiction,
			ReadmeMdUrl: m.ReadmeMdUrl,
		},
		Program: program,
	}
	handler.Success(c, output)
	return
}
