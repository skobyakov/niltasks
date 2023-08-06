package service

type Repository interface {
	GetList()
}

type Service struct {
	repo Repository
}

func New(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) GetList() {
	s.repo.GetList()
}
