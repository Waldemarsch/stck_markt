package models

type Currency struct {
}

type Stock struct {
	SecID     string  `json:"secID"`
	TradeNo   int64   `json:"tradeno"`
	TradeTime string  `json:"tradetime"`
	Price     float64 `json:"price"`
}

type StockCompany struct {
	Company string  `json:"company"`
	Stocks  []Stock `json:"stocks"`
}