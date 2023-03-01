package service

import (
	"career/model/web"
	"context"
)

type JobCategoryService interface {
	GetAllJobCategory(ctx context.Context) ([]web.JobCategory, error)
}
