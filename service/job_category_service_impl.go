package service

import (
	"career/model/web"
	"career/repository"
	"context"
	"database/sql"
)

type JobCategoryServiceImpl struct {
	DB                    *sql.DB
	JobCategoryRepository repository.JobCategoryRepository
}

func NewJobCategoryService(db *sql.DB, jobCategoryRepository repository.JobCategoryRepository) JobCategoryService {
	return &JobCategoryServiceImpl{
		DB:                    db,
		JobCategoryRepository: jobCategoryRepository,
	}
}

func (s *JobCategoryServiceImpl) GetAllJobCategory(ctx context.Context) ([]web.JobCategory, error) {
	res := []web.JobCategory{}

	jobCategories, err := s.JobCategoryRepository.GetAllJobCategory(ctx, s.DB)
	if err != nil {
		return res, err
	}
	for _, jobCategory := range jobCategories {
		res = append(res, web.JobCategory{
			Id:        jobCategory.Id,
			Name:      jobCategory.Name,
			Image:     jobCategory.Image,
			CreatedAt: jobCategory.CreatedAt,
			UpdatedAt: jobCategory.UpdatedAt,
		})
	}

	return res, nil
}
