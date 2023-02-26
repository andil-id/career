package repository

import (
	"career/model/domain"
	"context"
	"database/sql"
)

type JobRepository interface {
	CreateJob(ctx context.Context, tx *sql.Tx, job domain.Job) (string, error)
	GetJobTotal(ctx context.Context, db *sql.DB, companyName string, categoryId string, limit int, offset int) (int, error)
	GetAllJob(ctx context.Context, db *sql.DB, companyName string, categoryId string, limit int, offset int) ([]domain.Job, error)
	GetJobById(ctx context.Context, db *sql.DB, jobId string) (domain.Job, error)
	UpdateJob(ctx context.Context, tx *sql.Tx, job domain.Job) error
	DeleteJob(ctx context.Context, tx *sql.Tx, jobId string) error
}
