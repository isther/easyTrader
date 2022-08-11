package account

import (
	"context"
	"strconv"
	"sync"
)

type Worker struct {
	Symbols []string
	Context context.Context
	Account *Binance
}

func (worker *Worker) CheckIsOK(symbol string) (bool, error) {
	_, err := worker.Account.SetSymbol(symbol).GetExchangeInfo()
	if err != nil {
		return false, err
	}

	return true, nil
}

type SearchByAmplitueResult struct {
	Symbol string
	Klines []*Kline
}

func (worker *Worker) SearchSymbolsByAmplitue(inteval int64, amplitueLimit, amplituePercentageLimit float64) []*SearchByAmplitueResult {
	var (
		results []*SearchByAmplitueResult
	)

	for i := range worker.Symbols {
		var (
			symbol = worker.Symbols[i]
			result = new(SearchByAmplitueResult)
		)
		klines, err := worker.Account.SetSymbol(symbol).GetHttpKlinesWith1m(int(inteval), worker.Context)
		if err != nil {
			continue
		}

		for i := range klines {
			var kline = klines[i]

			high, _ := strconv.ParseFloat(kline.High, 64)
			low, _ := strconv.ParseFloat(kline.Low, 64)
			close, _ := strconv.ParseFloat(kline.Close, 64)

			amplitue := (high - low) / close
			amplituePercentage := low / close / amplitue

			if amplitue >= amplitueLimit && amplituePercentage >= amplituePercentageLimit {
				result.Klines = append(result.Klines, kline)
			}
		}

		if result.Klines != nil {
			result.Symbol = symbol
			results = append(results, result)
		}
	}
	return results
}

func (worker *Worker) SearchSymbolsByAmplitueMul(limit int, amplitueLimit, amplituePercentageLimit float64) []*SearchByAmplitueResult {
	var (
		results []*SearchByAmplitueResult
	)

	for i := range worker.Symbols {
		var (
			symbol = worker.Symbols[i]
			result = new(SearchByAmplitueResult)
			wg     sync.WaitGroup
			mu     sync.Mutex
		)
		klines, err := worker.Account.SetSymbol(symbol).GetHttpKlinesWith1m(limit, worker.Context)
		if err != nil {
			continue
		}

		for i := range klines {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()

				var kline = klines[i]
				high, _ := strconv.ParseFloat(kline.High, 64)
				low, _ := strconv.ParseFloat(kline.Low, 64)
				close, _ := strconv.ParseFloat(kline.Close, 64)
				amplitue := (high - low) / close
				amplituePercentage := low / close / amplitue

				if amplitue >= amplitueLimit && amplituePercentage >= amplituePercentageLimit {
					mu.Lock()
					result.Klines = append(result.Klines, kline)
					mu.Unlock()
				}
			}(i)
		}
		wg.Wait()

		if len(result.Klines) > 0 {
			results = append(results, result)
		}
	}
	return results
}
