package front

import (
	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/SubGame-Network/SubGameModuleService/internal/gin/handler"
	"github.com/SubGame-Network/SubGameModuleService/internal/provider"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ModuleResponse struct {
	Id        uint64 `json:"id"`
	Name      string `json:"name"`
	Depiction string `json:"depiction"`
	Tags      string `json:"tags"`
}

func Module(c *gin.Context) {
	repo, err := provider.NewMDRepo()
	if err != nil {
		zap.S().Warn("", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}

	m, err := repo.GetAllModule()
	if err != nil {
		zap.S().Warn("", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}

	var output []*ModuleResponse
	for _, v := range m {
		output = append(output, &ModuleResponse{
			Id:        v.ID,
			Name:      v.Name,
			Depiction: v.Depiction,
			Tags:      v.Tags,
		})
	}
	handler.Success(c, output)
	return
}
