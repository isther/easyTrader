package account

type Kline struct {
	OpenTime  int64
	CloseTime int64

	Open  string
	Close string

	High string
	Low  string

	Volume           string
	QuoteAssetVolume string
}
