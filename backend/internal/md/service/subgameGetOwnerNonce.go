package service

import (
	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/centrifuge/go-substrate-rpc-client/v3/types"
	"github.com/pkg/errors"
)

func (svc *MDService) SubGameGetOwnerNonce() (uint32, error) {
	api := svc.subgameApi

	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		return 0, errors.WithStack(err)
	}

	key, err := types.CreateStorageKey(meta, "System", "Account", SubGamePubkeyDecodeToPubkeyByte(svc.config.SubGame.MDOwnerPubkey), nil)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	var accountInfo types.AccountInfo
	ok, err := api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	if !ok {
		return 0, errors.WithStack(errors.New(domain.ErrorEventSubGameWalletNotActive.Message))
	}

	nonce := uint32(accountInfo.Nonce)

	return nonce, nil
}
