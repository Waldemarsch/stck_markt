package service

type Stocks interface {
}

type Currency interface {
}

type Service struct {
	Stocks
	Currency
}
