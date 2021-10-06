package front

import (
	"crypto/rand"
	"math/big"

	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/SubGame-Network/SubGameModuleService/internal/gin/handler"
	"github.com/SubGame-Network/SubGameModuleService/internal/provider"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UserAuth(c *gin.Context) {
	userAddress := c.GetString("userAddress")

	stakeRepo, err := provider.NewMDRepo()
	if err != nil {
		zap.S().Warnw("", "err", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}

	user, err := stakeRepo.GetUserByAddress(userAddress)
	if err != nil {
		zap.S().Warnw("", "err", err)
		handler.Failed(c, domain.ErrorServer, "")
		return
	}
	if user == nil {
		handler.Failed(c, domain.ErrorUserNotFound, domain.ErrorUserNotFound.Message)
		return
	}

	nonce, _ := rand.Int(rand.Reader, big.NewInt(10000))

	stakeRepo.UpdateUserNonce(user.ID, nonce.String())

	handler.Success(c, nonce)
}
