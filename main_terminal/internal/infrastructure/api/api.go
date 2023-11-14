package api

type WebAPI interface {
}

type ExchangeAPI interface {
}

type API struct {
	WebAPI
	ExchangeAPI
}

func NewAPI() *API {
	return &API{
		WebAPI:      nil,
		ExchangeAPI: nil,
	}
}
