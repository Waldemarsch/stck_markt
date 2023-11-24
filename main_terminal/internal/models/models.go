package models

type Currency struct {
}

type Stock struct {
	TradeNo   string `json:"tradeno"`
	TradeTime string `json:"tradetime"`
	Price     string `json:"price"`
}

type StockCompany struct {
	Company string  `json:"company"`
	Stocks  []Stock `json:"stocks"`
}
