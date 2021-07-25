package submission

type Service interface {

}

type svc struct {
	repo Repository
}


func NewService(repo Repository) Service {
	return &svc{
		repo: repo,
	}
}