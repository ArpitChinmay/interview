package setup

import (
	"database/sql"

	"github.com/ArpitChinmay/interview/repositories"
)

type RepositoryGenerator struct {
	InterviewRepository repositories.Repository
}

func SetupRepositories(db *sql.DB) *RepositoryGenerator {
	repository := repositories.InitializeRepository(db)
	return &RepositoryGenerator{InterviewRepository: repository}
}
