package service

type Repo interface {
	AddLinksPair(short, long string) error
	GetLongURL(short string) (string, error)
}

type Service struct {
	repo Repo
}

func New(r Repo) *Service {
	return &Service{
		repo: r,
	}
}
