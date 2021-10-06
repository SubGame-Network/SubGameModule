package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"crypto/md5"

	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type MarketInfo struct {
	CurrencyId        int             `json:"currency_id"`
	Logo              string          `json:"logo"`
	Name              string          `json:"name"`
	Symbol            string          `json:"symbol"`
	PriceCny          decimal.Decimal `json:"price_cny"`
	PriceUsd          decimal.Decimal `json:"price_usd"`
	Volume24hCny      decimal.Decimal `json:"volume_24h_cny"`
	Volume24hUsd      decimal.Decimal `json:"volume_24h_usd"`
	PercentChangeUtc0 decimal.Decimal `json:"percent_change_utc0"`
	PercentChangeUtc8 decimal.Decimal `json:"percent_change_utc8"`
	PriceUpdatedAt    decimal.Decimal `json:"price_updated_at"`
}

var CurrencyIDMap = map[domain.CoinType]string{
	domain.CoinTypeSGB: "362151", // subgame
}

const MyTokenURL = "http://openapi.mytokenapi.com"
const APPID = "f04a147467d1564e"
const SECRET = "5b5c772a9de501b3e045ffa86178f800"

// 獲得幣種價值
func (svc *MDService) CurrencyToUsd(coin domain.CoinType) (decimal.Decimal, error) {
	// 預設先找MyToken，沒錯誤就回傳匯率
	price, err := svc.MyTokenRateAPI(coin)
	if err == nil {
		return price, nil
	}

	// 備用找BiKi匯率
	return svc.BiKiRateAPI(coin)
}

func (svc *MDService) MyTokenRateAPI(coin domain.CoinType) (decimal.Decimal, error) {
	if CurrencyIDMap[coin] == "" {
		return decimal.Zero, errors.WithStack(errors.New("not exist currency id"))
	}

	sucRes := struct {
		Code    int          `json:"code"`
		Message string       `json:"message"`
		Data    []MarketInfo `json:"data"`
	}{}

	stringA := fmt.Sprintf("app_id=%v&currency_ids=%v&timestamp=%v", APPID, CurrencyIDMap[coin], time.Now().Unix())
	stringSignTemp := fmt.Sprintf("%v&app_secret=%v", stringA, SECRET)
	sign := strings.ToUpper(hmacSha256(stringSignTemp, SECRET))

	url := fmt.Sprintf("%v/currency/marketinfo?%v&sign=%v", MyTokenURL, stringA, sign)
	zap.S().Infow("MyToken 匯率 API", "url", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return decimal.Zero, errors.WithStack(err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return decimal.Zero, errors.WithStack(err)
	}
	defer res.Body.Close()

	// deal with response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return decimal.Zero, errors.WithStack(err)
	}

	// request success
	if len(body) != 0 {
		if err := json.Unmarshal(body, &sucRes); err != nil {
			return decimal.Zero, errors.WithStack(err)
		}
		if len(sucRes.Data) < 1 {
			return decimal.Zero, errors.WithStack(fmt.Errorf("mytoken service, response body data len < 1, sucRes = %v", sucRes))
		}
		return sucRes.Data[0].PriceUsd, nil
	}

	return decimal.Zero, errors.WithStack(fmt.Errorf("mytoken service, response body null"))
}

type BikiMarketInfo struct {
	High decimal.Decimal `json:"high"`
	Vol  decimal.Decimal `json:"vol"`
	Last decimal.Decimal `json:"last"`
	Low  decimal.Decimal `json:"low"`
	Buy  decimal.Decimal `json:"buy"`
	Sell decimal.Decimal `json:"sell"`
	Rose decimal.Decimal `json:"rose"`
	Time decimal.Decimal `json:"time"`
}

func (svc *MDService) BiKiRateAPI(coin domain.CoinType) (decimal.Decimal, error) {
	time.Sleep(1 * time.Second)

	var BiKiCurrencyIDMap = map[domain.CoinType]string{
		domain.CoinTypeSGB: "sgbusdt", // subgame
	}

	const BiKiURL = "https://openapi.biki.com"
	const BiKiAPPID = "86d9853ab50d1db5f324eb4081e05b4e"
	const BiKiSECRET = "ea7654f656e358cc44c9c895cdacd4a8"

	if BiKiCurrencyIDMap[coin] == "" {
		return decimal.Zero, errors.WithStack(errors.New("not exist currency id"))
	}

	sucRes := struct {
		Code    string         `json:"code"`
		Message string         `json:"msg"`
		Data    BikiMarketInfo `json:"data"`
	}{}

	apiTime := time.Now().Unix()

	// sign加密字串
	signData := fmt.Sprintf("api_key%ssymbol%vtime%v%v", BiKiAPPID, BiKiCurrencyIDMap[coin], apiTime, BiKiSECRET)
	sign := fmt.Sprintf("%x", md5.Sum([]byte(signData)))

	url := fmt.Sprintf("%v/open/api/get_ticker?api_key=%v&symbol=%v&time=%v&sign=%v", BiKiURL, BiKiAPPID, BiKiCurrencyIDMap[coin], apiTime, sign)
	zap.S().Infow("BiKi 匯率 API", "url", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return decimal.Zero, errors.WithStack(err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return decimal.Zero, errors.WithStack(err)
	}
	defer res.Body.Close()

	// deal with response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return decimal.Zero, errors.WithStack(err)
	}

	// request success
	if len(body) != 0 {
		if err := json.Unmarshal(body, &sucRes); err != nil {
			return decimal.Zero, errors.WithStack(err)
		}
		if sucRes.Message != "suc" {
			return decimal.Zero, errors.WithStack(fmt.Errorf("biki service, Err = %v", string(body)))
		}
		return sucRes.Data.Last, nil
	}

	return decimal.Zero, errors.WithStack(fmt.Errorf("biki service, response body null"))
}

func hmacSha256(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
