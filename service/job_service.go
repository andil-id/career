package service

import (
	"career/model/web"
	"context"
)

type JobService interface {
	DeleteJob(ctx context.Context, jobId string) error
	GetJobDetail(ctx context.Context, jobId string) (web.Job, error)
	CreateJob(ctx context.Context, data web.CreateJob) (web.Job, error)
	UpdateJob(ctx context.Context, data web.UpdateJob, jobId string) (web.Job, error)
	GetAllJob(ctx context.Context, companyName string, categoryId string, limit string, offset string) ([]web.Job, web.Pagination, error)
}
