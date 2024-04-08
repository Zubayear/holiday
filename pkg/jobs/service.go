package jobs

import (
	"context"
	"github.com/Zubayear/holiday/pkg/entities"
)

type Service interface {
	GetJobs(ctx context.Context) ([]*entities.Job, error)
	GetJob(ctx context.Context, id uint64) (*entities.Job, error)
	ScheduleJob(ctx context.Context, job entities.Job) bool
	FetchJobStatus(ctx context.Context, id uint64) (string, error)
}

type service struct {
	Repo Repository
}

func NewService(repo Repository) Service {
	return &service{Repo: repo}
}

func (s *service) GetJobs(ctx context.Context) ([]*entities.Job, error) {
	return s.Repo.FetchJobs(ctx)
}

func (s *service) GetJob(ctx context.Context, id uint64) (*entities.Job, error) {
	return s.Repo.FetchJob(ctx, id)
}

func (s *service) ScheduleJob(ctx context.Context, job entities.Job) bool {
	return s.Repo.ScheduleJob(ctx, job)
}

func (s *service) FetchJobStatus(ctx context.Context, id uint64) (string, error) {
	return s.Repo.FetchJobStatus(ctx, id)
}
