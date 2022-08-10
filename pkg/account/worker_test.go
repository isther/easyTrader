package account

import (
	"context"
	"testing"
)

func TestAccountSearchByAmplitue(t *testing.T) {
	worker := &Worker{
		Symbols: []string{"BUSDUSDT","ETHUSDT", "BTCUSDT"},
		Context: context.Background(),
		Account: getBinance(),
	}

	res := worker.SearchSymbolsByAmplitue(120, 0.001, 0.001)
	for i := range res {
		t.Log(res[i].Symbol, "Len = ", len(res[i].Klines))
		for j := range res[i].Klines {
			t.Log(res[i].Klines[j])
		}
	}
}

func TestAccountSearchByAmplitueMul(t *testing.T) {
	worker := &Worker{
		Symbols: []string{"BUSDUSDT","ETHUSDT", "BTCUSDT"},
		Context: context.Background(),
		Account: getBinance(),
	}

	res := worker.SearchSymbolsByAmplitueMul(120, 0.001, 0.001)
	for i := range res {
		t.Log(res[i].Symbol, "Len = ", len(res[i].Klines))
		for j := range res[i].Klines {
			t.Log(res[i].Klines[j])
		}
	}
}
