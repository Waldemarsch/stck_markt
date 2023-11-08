package infrastructure

type WebAPI interface {
}

type Repository interface {
}

type Infrastructure struct {
	WebAPI
	Repository
}
