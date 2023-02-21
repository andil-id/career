package repository

import (
	"career/model/domain"
	"database/sql"

	"golang.org/x/net/context"
)

type JobCategoryRepositoryImpl struct{}

func NewJobCategoryRespository() JobCategoryRepository {
	return &JobCategoryRepositoryImpl{}
}

func (r *JobCategoryRepositoryImpl) GetAllJobCategory(ctx context.Context, db *sql.DB) ([]domain.JobCategory, error) {
	SQL := "SELECT * FROM job_category"
	rows, err := db.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	jobCategories := []domain.JobCategory{}
	for rows.Next() {
		jobCategory := domain.JobCategory{}
		err := rows.Scan(&jobCategory.Id, &jobCategory.Name, &jobCategory.Image)
		if err != nil {
			panic(err)
		}
		jobCategories = append(jobCategories, jobCategory)
	}
	return jobCategories, nil
}
