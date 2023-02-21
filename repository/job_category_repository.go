package repository

import (
	"career/model/domain"
	"database/sql"

	"golang.org/x/net/context"
)

type JobCategoryRepository interface {
	GetAllJobCategory(ctx context.Context, db *sql.DB) ([]domain.JobCategory, error)
}
