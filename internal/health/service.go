package health

import (
	"github.com/ezegrosfeld/basic-api/internal/domain"
	"github.com/ezegrosfeld/basic-api/pkg/logger"
	"github.com/rs/zerolog"
)

type Service interface {
	Get() (domain.Health, error)
}

type service struct {
	log  zerolog.Logger
	repo Repository
}

func NewService(repo Repository) *service {
	logger := logger.New("health", "service")
	return &service{
		log:  logger,
		repo: repo,
	}
}

func (s *service) Get() (domain.Health, error) {
	s.log.Debug().Msg("getting health from service")
	health, err := s.repo.Get()
	if err != nil {
		s.log.Error().Err(err).Msg("error retrieving health")
		return domain.Health{}, err
	}
	return health, nil
}
