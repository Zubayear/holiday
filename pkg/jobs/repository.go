package jobs

import (
	"context"
	"github.com/Zubayear/holiday/pkg/entities"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type Repository interface {
	FetchJobs(ctx context.Context) ([]*entities.Job, error)
	FetchJob(ctx context.Context, id uint64) (*entities.Job, error)
	ScheduleJob(ctx context.Context, job entities.Job) bool
	FetchJobStatus(ctx context.Context, id uint64) (string, error)
}

type repository struct {
	Pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) Repository {
	return &repository{Pool: pool}
}

func (r *repository) FetchJobs(ctx context.Context) ([]*entities.Job, error) {
	rows, err := r.Pool.Query(ctx, "SELECT * FROM jobs LIMIT 50")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var jobs []*entities.Job
	for rows.Next() {
		var job entities.Job
		err := rows.Scan(&job.Id, &job.JobName, &job.Description, &job.Environment, &job.Status, &job.StartTime, &job.EndTime, &job.Priority)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		jobs = append(jobs, &job)
	}

	return jobs, nil

}

func (r *repository) FetchJob(ctx context.Context, id uint64) (*entities.Job, error) {
	row := r.Pool.QueryRow(ctx, "SELECT * FROM jobs WHERE id = $1", id)
	var job entities.Job
	err := row.Scan(&job.Id, &job.JobName, &job.Description, &job.Environment, &job.Status, &job.StartTime, &job.EndTime, &job.Priority)
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (r *repository) ScheduleJob(ctx context.Context, job entities.Job) bool {
	tag, err := r.Pool.Exec(ctx, "INSERT INTO jobs (job_name, description, envirnoment, status, start_time, end_time, priority) VALUES ($1, $2, $3, $4, $5, $6, $7)", job.JobName, job.Description, job.Environment, job.Status, job.StartTime, job.EndTime)
	if err != nil {
		return false
	}
	return tag.RowsAffected() >= 1
}

func (r *repository) FetchJobStatus(ctx context.Context, id uint64) (string, error) {
	row := r.Pool.QueryRow(ctx, "SELECT status FROM jobs WHERE id = $1", id)
	var result string
	err := row.Scan(&result)
	if err != nil {
		return "", err
	}
	return result, nil
}
