package models

type Currency struct {
}

type Stock struct {
	TradeNo   int64   `json:"tradeno"`
	TradeTime string  `json:"tradetime"`
	Price     float64 `json:"price"`
}
