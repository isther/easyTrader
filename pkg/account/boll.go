package account

import (
	"math"
	"strconv"
)

type Boll struct {
	Length     int
	Multiplier int
}

func NewBoll(length int, multiplier int) *Boll {
	return &Boll{
		Length:     length,
		Multiplier: multiplier,
	}
}

func (boll *Boll) Boll(klines []*Kline) (float64, float64, float64) {
	var (
		MA float64
		MD float64
	)

	MA = boll.ma(klines[len(klines)-boll.Length+1:])
	MD = boll.md(klines[len(klines)-boll.Length+1:], MA)
	return MA + float64(boll.Multiplier)*MD, MA, MA - float64(boll.Multiplier)*MD
}

func (boll *Boll) ma(klines []*Kline) float64 {
	var (
		N   = len(klines)
		SUM float64
	)

	for i := range klines {
		var kline = klines[i]

		close, _ := strconv.ParseFloat(kline.Close, 64)
		SUM += close
	}

	return SUM / float64(N)
}

func (boll *Boll) md(klines []*Kline, ma float64) float64 {
	var (
		N   = len(klines)
		SUM float64
	)

	for i := range klines {
		var kline = klines[i]

		close, _ := strconv.ParseFloat(kline.Close, 64)
		SUM += (close - ma) * (close - ma)
	}

	return math.Sqrt(SUM / float64(N))
}
