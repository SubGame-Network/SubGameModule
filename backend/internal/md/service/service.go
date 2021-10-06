package service

import (
	"sync"

	"github.com/SubGame-Network/SubGameModuleService/config"
	"github.com/SubGame-Network/SubGameModuleService/domain"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v3"
	"github.com/itering/substrate-api-rpc"
	"github.com/itering/substrate-api-rpc/metadata"
	"github.com/itering/substrate-api-rpc/pkg/recws"
	"github.com/itering/substrate-api-rpc/rpc"
	"github.com/itering/substrate-api-rpc/websocket"
)

type MDService struct {
	config        *config.Config
	MDRepo        domain.MDRepository
	subgameWsConn *recws.RecConn
	subgameApi    *gsrpc.SubstrateAPI
	notifySvc     domain.NotifyService
	redis         domain.GoRedis
}

var subgameWsConnOnce sync.Once
var subgameWsConn *recws.RecConn

var subgameOnce sync.Once
var subgameApi *gsrpc.SubstrateAPI

func NewMDService(config *config.Config, MDRepo domain.MDRepository, notifySvc domain.NotifyService, redisServ domain.GoRedis) (domain.MDService, error) {
	websocket.SetEndpoint(config.SubGame.RPC)
	poolConn, err := websocket.Init()
	if err != nil {
		panic(err)
	}
	subgameWsConn = poolConn.Conn

	// chain types
	var hash []string
	metaRaw, err := rpc.GetMetadataByHash(nil, hash...)
	if err != nil {
		panic(err)
	}
	metaRuntimeRaw := &metadata.RuntimeRaw{
		Spec: 1,
		Raw:  metaRaw,
	}
	metadata.Latest(metaRuntimeRaw)
	substrate.RegCustomTypes(config.SubGame.ChainRuntimeTypes)

	if subgameApi == nil {
		var err error
		subgameApi, err = gsrpc.NewSubstrateAPI(config.SubGame.RPC)
		if err != nil {
			panic(err)
		}
	}

	return &MDService{
		config:        config,
		MDRepo:        MDRepo,
		subgameWsConn: subgameWsConn,
		subgameApi:    subgameApi,
		notifySvc:     notifySvc,
		redis:         redisServ,
	}, nil
}
