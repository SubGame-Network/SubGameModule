package domain

import (
	"github.com/shopspring/decimal"
)

const (
	ChainWsBlockHash = iota + 1
	ChainWsBlock
	ChainWsEvent
	ChainWsSpec
)

type SubGameChainEvent struct {
	EventIndex    string      `json:"event_index"`
	BlockNum      int         `json:"block_num" `
	ExtrinsicIdx  int         `json:"extrinsic_idx"`
	ModuleId      string      `json:"module_id" `
	EventId       string      `json:"event_id" `
	Params        interface{} `json:"params"`
	ExtrinsicHash string      `json:"extrinsic_hash"`
	EventIdx      int         `json:"event_idx"`
}

type SubGameChainExtrinsic struct {
	ExtrinsicIndex     string                  `json:"extrinsic_index"`
	BlockNum           int                     `json:"block_num" `
	BlockTimestamp     int                     `json:"block_timestamp"`
	ExtrinsicLength    string                  `json:"extrinsic_length"`
	VersionInfo        string                  `json:"version_info"`
	CallCode           string                  `json:"call_code"`
	CallModuleFunction string                  `json:"call_module_function"`
	CallModule         string                  `json:"call_module"`
	Params             []SubGameExtrinsicParam `json:"params"`
	AccountId          string                  `json:"account_id"`
	Signature          string                  `json:"signature"`
	Nonce              int                     `json:"nonce"`
	Era                string                  `json:"era"`
	ExtrinsicHash      string                  `json:"extrinsic_hash"`
	IsSigned           bool                    `json:"is_signed"`
	Success            bool                    `json:"success"`
	Fee                decimal.Decimal         `json:"fee"`
}

type SubGameExtrinsicParam struct {
	Name  string      `json:"name"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

type SubGameAccountDecode struct {
	Nonce    int `json:"nonce"`
	RefCount int `json:"ref_count"`
	Data     struct {
		Free       decimal.Decimal `json:"free"`
		Reserved   decimal.Decimal `json:"reserved"`
		MiscFrozen decimal.Decimal `json:"miscFrozen"`
		FeeFrozen  decimal.Decimal `json:"feeFrozen"`
	} `json:"data"`
}
