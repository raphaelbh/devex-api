package usecase

import (
	"log/slog"

	"github.com/raphaelbh/devex-api/internal/infra/repository"
)

var (
	logger = slog.Default()

	// repositories
	secretRepository   = repository.GetSecretRepository()
	pipelineRepository = repository.GetPipelineRepository()
)
