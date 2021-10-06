package middleware

import (
	sr25519 "github.com/ChainSafe/go-schnorrkel"
	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/SubGame-Network/SubGameModuleService/internal/gin/handler"
	SubGameHelp "github.com/SubGame-Network/SubGameModuleService/internal/md/service"
	"github.com/SubGame-Network/SubGameModuleService/internal/provider"
	"github.com/centrifuge/go-substrate-rpc-client/v3/types"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func FrontSignMiddleware(c *gin.Context) {
	sigHex := c.GetHeader("SigHex")
	if sigHex == "" {
		handler.Failed(c, domain.ErrorUnauthorized, "")
		c.Abort()
		return
	}
	sigByte := types.MustHexDecodeString(sigHex)
	sigByteStart := 0
	sigByteEnd := len(sigByte)
	if sigByteEnd > 64 {
		sigByteStart = sigByteEnd - 64
	}

	userAddress := c.GetString("userAddress")

	MDRepo, err := provider.NewMDRepo()
	if err != nil {
		zap.S().Warnw("", "err", err)
		handler.Failed(c, domain.ErrorServer, "")
		c.Abort()
		return
	}

	user, err := MDRepo.GetUserByAddress(userAddress)
	if err != nil {
		zap.S().Warnw("", "err", err)
		handler.Failed(c, domain.ErrorServer, "")
		c.Abort()
		return
	}
	if user == nil || user.Address == "" {
		handler.Failed(c, domain.ErrorUserNotFound, domain.ErrorUserNotFound.Message)
		c.Abort()
		return
	}
	data := []byte(user.Nonce)

	pubkeyByte := SubGameHelp.SubGamePubkeyDecodeToPubkeyByte(userAddress)

	pub := [32]byte{}
	copy(pub[:], pubkeyByte)

	var sigs [64]byte
	copy(sigs[:], sigByte[sigByteStart:sigByteEnd])
	sig := new(sr25519.Signature)
	if err := sig.Decode(sigs); err != nil {
		handler.Failed(c, domain.ErrorUnauthorized, err.Error())
		c.Abort()
		return
	}
	signingContext := sr25519.NewSigningContext([]byte("substrate"), data)

	result := sr25519.NewPublicKey(pub).Verify(sig, signingContext)
	if result == false {
		handler.Failed(c, domain.ErrorUnauthorized, "")
		c.Abort()
		return
	}

	c.Next()
}
