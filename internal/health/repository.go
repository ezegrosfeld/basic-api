package health

import (
	"database/sql"

	"github.com/ezegrosfeld/basic-api/internal/domain"
	"github.com/ezegrosfeld/basic-api/pkg/db"
	"github.com/ezegrosfeld/basic-api/pkg/logger"
	"github.com/rs/zerolog"
)

type Repository interface {
	Get() (domain.Health, error)
}

type repo struct {
	log zerolog.Logger
	db  *sql.DB
}

// NewRepository returns a new Repository.
func NewRepository() *repo {
	log := logger.New("health", "repository")
	return &repo{
		log: log,
		db:  db.Database,
	}
}

func (r *repo) Get() (domain.Health, error) {
	r.log.Debug().Msg("getting health")

	health := domain.Health{
		Status:  "OK",
		Message: "Todo de 10!",
	}

	return health, nil
}
