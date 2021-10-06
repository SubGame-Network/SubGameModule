package service

import (
	"encoding/json"

	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/itering/subscan/util"
)

func (svc *MDService) SubGameGetExtrinsicTimestamp(extrinsic *domain.SubGameChainExtrinsic) (blockTimestamp int) {
	if extrinsic.CallModule != "timestamp" {
		return
	}

	var paramsInstant []domain.SubGameExtrinsicParam
	paramsByte, _ := json.Marshal(extrinsic.Params)
	json.Unmarshal(paramsByte, &paramsInstant)

	for _, p := range paramsInstant {
		if p.Name == "now" {
			extrinsic.BlockTimestamp = util.IntFromInterface(p.Value)
			return extrinsic.BlockTimestamp
		}
	}

	return
}
