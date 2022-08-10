package account

import (
	"context"

	libBinance "github.com/adshao/go-binance/v2"
	"github.com/sirupsen/logrus"
)

type Binance struct {
	ApiKey    string
	SecretKey string
	Client    *libBinance.Client

	Symbol string
}

func NewBinance(apiKey, secretKey string) *Binance {
	return &Binance{
		ApiKey:    apiKey,
		SecretKey: secretKey,
		Client:    libBinance.NewClient(apiKey, secretKey),
	}
}

// Switch to the proxy client
func (binance *Binance) WithProxyClient(proxyUrl string) *Binance {
	binance.Client = libBinance.NewProxiedClient(binance.ApiKey, binance.SecretKey, proxyUrl)
	return binance
}

func (binance *Binance) SetSymbol(symbol string) *Binance {
	binance.Symbol = symbol
	return binance
}

func (binance *Binance) GetSymbol() string {
	return binance.Symbol
}

func (binance *Binance) GetHttpKlinesWith1m(limit int, ctx context.Context) ([]*Kline, error) {
	var (
		err    error
		klines []*Kline
	)
	res, err := binance.Client.NewKlinesService().Symbol(binance.Symbol).Interval("1m").Limit(limit).Do(ctx)
	if err != nil {
		return nil, err
	}

	for i := range res {
		var kline = res[i]
		klines = append(klines, &Kline{
			OpenTime:         kline.OpenTime,
			CloseTime:        kline.CloseTime,
			Open:             kline.Open,
			Close:            kline.Close,
			High:             kline.High,
			Low:              kline.Low,
			Volume:           kline.Volume,
			QuoteAssetVolume: kline.QuoteAssetVolume,
		})
	}
	return klines, nil
}

func (binance *Binance) GetWsKlinesWith1m() (chan *Kline, chan struct{}, chan struct{}) {
	var (
		err error

		klineC = make(chan *Kline)
		doneC  chan struct{}
		stopC  chan struct{}
	)

	wsKlineHandler := func(event *libBinance.WsKlineEvent) {
		var kline = event.Kline
		klineC <- &Kline{
			OpenTime:         kline.StartTime,
			CloseTime:        kline.EndTime,
			Open:             kline.Open,
			Close:            kline.Close,
			High:             kline.High,
			Low:              kline.Low,
			Volume:           kline.Volume,
			QuoteAssetVolume: kline.QuoteVolume,
		}
	}

	errHandler := func(err error) {
		logrus.Error(err)
	}

	doneC, stopC, err = libBinance.WsKlineServe(binance.Symbol, "1m", wsKlineHandler, errHandler)

	if err != nil {
		logrus.Error(err)
		return nil, doneC, stopC
	}
	return klineC, doneC, stopC
}

func (binance *Binance) GetExchangeInfo() (*libBinance.ExchangeInfo, error) {
	res, err := binance.Client.NewExchangeInfoService().Symbol(binance.Symbol).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return res, nil
}
