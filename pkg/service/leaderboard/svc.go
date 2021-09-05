package leaderboard

type Service interface {

}

type svc struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &svc{
		repo: r,
	}
}


