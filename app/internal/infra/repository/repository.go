package repository

import "log/slog"

var (
	logger = slog.Default()

	secretRepository   = SecretRepository{}
	pipelineRepository = PipelineRepository{}
)

func GetSecretRepository() SecretRepository {
	return secretRepository
}

func GetPipelineRepository() PipelineRepository {
	return pipelineRepository
}
