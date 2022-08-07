package account

import (
	"testing"
	"time"

	"golang.org/x/net/context"
)

func TestBinanceBoll(t *testing.T) {
	var (
		err error

		ctx = context.Background()
	)

	klines, err := getBinance().SetSymbol("ETHUSDT").GetHttpKlinesWith1m(120, ctx)
	if err != nil {
		t.Fatal(err)
	}

	up, mb, dn := NewBoll(20, 2).Boll(klines)
	t.Logf("\nUP: %f \nMB: %f \nDN: %f \n", up, mb, dn)
}

func getBinance() *Binance {
	var (
		err error

		ctx, cancelFunc = context.WithTimeout(context.Background(), 3*time.Second)
	)
	defer cancelFunc()

	err = NewBinance("", "").Client.NewPingService().Do(ctx)
	if err != nil {
		return NewBinance("", "").WithProxyClient("http://localhost:20171")
	}

	return NewBinance("", "")
}
