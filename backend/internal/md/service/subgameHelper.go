package service

import (
	"fmt"
	"math"
	"strings"

	"github.com/SubGame-Network/SubGameModuleService/config"
	"github.com/btcsuite/btcutil/base58"
	"github.com/centrifuge/go-substrate-rpc-client/v3/hash"
	"github.com/centrifuge/go-substrate-rpc-client/v3/types"
	"github.com/shopspring/decimal"
)

func SubGameTypesAccountIDToPubkey(input types.AccountID) (string, error) {
	pubkey := fmt.Sprintf("%x", input[:])
	if len(pubkey) != 64 {
		return "", fmt.Errorf("not pubkey: %s", pubkey)
	}
	return "0x" + pubkey, nil
}

func SubGamePubkeyToAddress(config *config.Config, pubkey string) (string, error) {
	addressType := ""
	switch config.SubGame.SS58Prefix {
	case "42":
		addressType = "2a"
	case "27":
		addressType = "1b"
	default:
		return "", fmt.Errorf("unknow SS58Prefix: %v", config.SubGame.SS58Prefix)
	}

	pubkey = strings.TrimPrefix(pubkey, "0x")
	var addr string = ""
	contextPrefix := []byte("SS58PRE")
	noSum := types.MustHexDecodeString(addressType + pubkey)
	all := append(contextPrefix, noSum...)
	checksum, err := hash.NewBlake2b512(nil)
	if err != nil {
		return addr, err
	}
	checksum.Write(all)
	resSum := checksum.Sum(nil)
	addr = base58.Encode(types.MustHexDecodeString(addressType + pubkey + fmt.Sprintf("%x", resSum[:2])))
	return addr, nil
}

func SubGameAddressToPubkey(input string) (string, error) {
	decodedByte := base58.Decode(input)
	decodedStr := fmt.Sprintf("%x", decodedByte)
	decodedLen := len(decodedStr)
	if decodedLen < 70 {
		return "", fmt.Errorf("not address: %s", input)
	}
	lenStart := 2 + (decodedLen - 70)
	lenEnd := 66 + (decodedLen - 70)
	pubkey := "0x" + decodedStr[lenStart:lenEnd]
	_, err := types.NewAddressFromHexAccountID(pubkey)
	if err != nil {
		return "", fmt.Errorf("not address: %s", input)
	}
	return pubkey, nil
}

func SubGameAddressDecodeToPubkeyByte(addr string) []byte {
	decodedByte := base58.Decode(addr)
	decodedStr := fmt.Sprintf("%x", decodedByte)
	decodedLen := len(decodedStr)
	if decodedLen < 70 {
		return nil
	}
	lenStart := 2 + (decodedLen - 70)
	lenEnd := 66 + (decodedLen - 70)
	pubkeyStr := "0x" + fmt.Sprintf("%x", decodedByte)[lenStart:lenEnd]
	pubkeyByte := types.MustHexDecodeString(pubkeyStr)
	return pubkeyByte
}

func SubGamePubkeyDecodeToPubkeyByte(input string) []byte {
	input = strings.TrimPrefix(input, "0x")
	pubkeyStr := "0x" + input
	pubkeyByte := types.MustHexDecodeString(pubkeyStr)
	return pubkeyByte
}

func SubGameDotToUnit(input decimal.Decimal) decimal.Decimal {
	val := decimal.NewFromFloat(math.Pow10(10))
	return input.Div(val)
}

func SubGameUnitToDot(input decimal.Decimal) decimal.Decimal {
	val := decimal.NewFromFloat(math.Pow10(10))
	return input.Mul(val)
}

func SubGameCheckStringToHex(s string) string {
	if strings.TrimSpace(s) == "" {
		return ""
	}
	if strings.HasPrefix(s, "0x") {
		return s
	}
	return strings.ToLower("0x" + s)
}
