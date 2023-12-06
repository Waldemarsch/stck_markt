package infrastructure

type Cache interface {
	GetStocks()
}

type Infrastructure struct {
	Cache
}
