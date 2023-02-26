package service

import (
	"career/exception"
	"career/helper"
	"career/model/domain"
	"career/model/web"
	"career/repository"
	"context"
	"database/sql"
	"encoding/json"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	e "github.com/pkg/errors"
)

type JobServiceImpl struct {
	DB            *sql.DB
	Validate      *validator.Validate
	JobRepository repository.JobRepository
}

func NewJobService(jobRepository repository.JobRepository, db *sql.DB, validate *validator.Validate) JobService {
	return &JobServiceImpl{
		DB:            db,
		Validate:      validate,
		JobRepository: jobRepository,
	}
}

func (s *JobServiceImpl) CreateJob(ctx context.Context, data web.CreateJob) (web.Job, error) {
	now := time.Now()
	res := web.Job{}

	err := s.Validate.Struct(data)
	if err != nil {
		return res, err
	}
	tx, err := s.DB.Begin()
	if err != nil {
		return res, err
	}
	defer helper.CommitOrRollback(tx)

	companyProfilePath, err := helper.FirebaseImageUploader(ctx, data.CompanyLogo, "compro")
	if err != nil {
		return res, err
	}
	bannerPath, err := helper.FirebaseMultipleImageUploader(ctx, data.Banner, "banner")
	if err != nil {
		return res, err
	}

	bannerPathStr, err := json.Marshal(bannerPath)
	if err != nil {
		return res, err
	}

	job := domain.Job{
		CategoryId:  data.CategoryId,
		CompanyLogo: companyProfilePath,
		CompanyName: data.CompanyName,
		Location:    data.Location,
		Title:       data.Title,
		Type:        data.Title,
		Banner:      string(bannerPathStr),
		Description: data.Description,
		Email:       data.Email,
		WebsiteUrl:  data.WebsiteUrl,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	id, err := s.JobRepository.CreateJob(ctx, tx, job)
	if err != nil {
		return res, err
	}

	res = web.Job{
		Id:          id,
		CategoryId:  data.CategoryId,
		CompanyLogo: companyProfilePath,
		CompanyName: data.CompanyName,
		Location:    data.Location,
		Title:       data.Title,
		Type:        data.Type,
		Banner:      bannerPath,
		Description: data.Description,
		Email:       data.Email,
		WebsiteUrl:  data.WebsiteUrl,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	return res, nil
}

func (s *JobServiceImpl) GetAllJob(ctx context.Context, companyName string, categoryId string, limit string, offset string) ([]web.Job, web.Pagination, error) {
	var err error
	var res []web.Job
	var pagination = web.Pagination{
		Limit:     0,
		Offset:    0,
		RowCount:  0,
		PageCount: 0,
	}

	limitStr := "10"
	offsetStr := "1"

	if limit != "" {
		limitStr = limit
	}
	if offset != "" {
		offsetStr = offset
	}

	pagination.Limit, err = strconv.Atoi(limitStr)
	if err != nil {
		return res, pagination, err
	}

	pagination.Offset, err = strconv.Atoi(offsetStr)
	if err != nil {
		return res, pagination, err
	}

	if pagination.Limit > 1000 || pagination.Limit < 1 {
		return res, pagination, e.Wrap(exception.ErrBadRequest, "limit parameter out of range")
	}
	if pagination.Offset <= 0 {
		return res, pagination, e.Wrap(exception.ErrBadRequest, "offset parameter must greater than 0")
	}

	totalRecords, err := s.JobRepository.GetJobTotal(ctx, s.DB, companyName, categoryId, pagination.Limit, pagination.Offset-1)
	if err != nil {
		return res, pagination, err
	}
	pagination.RowCount = totalRecords
	pagination.PageCount = (totalRecords + pagination.Limit - 1) / pagination.Limit

	jobs, err := s.JobRepository.GetAllJob(ctx, s.DB, companyName, categoryId, pagination.Limit, pagination.Offset-1)
	if err != nil {
		return res, pagination, err
	}

	for _, job := range jobs {
		var arrBanner []string
		err = json.Unmarshal([]byte(job.Banner), &arrBanner)
		if err != nil {
			return res, pagination, err
		}
		res = append(res, web.Job{
			Id:          job.Id,
			CategoryId:  job.CategoryId,
			CompanyLogo: job.CompanyLogo,
			CompanyName: job.CompanyName,
			Location:    job.Location,
			Title:       job.Title,
			Type:        job.Type,
			Banner:      arrBanner,
			Description: job.Description,
			Email:       job.Email,
			WebsiteUrl:  job.WebsiteUrl,
			CreatedAt:   job.UpdatedAt,
			UpdatedAt:   job.UpdatedAt,
		})
	}
	return res, pagination, nil
}

func (s *JobServiceImpl) GetJobDetail(ctx context.Context, jobId string) (web.Job, error) {
	var res web.Job
	var arrBanner []string

	job, err := s.JobRepository.GetJobById(ctx, s.DB, jobId)
	if err != nil {
		return res, e.Wrap(exception.ErrNotFound, err.Error())
	}
	err = json.Unmarshal([]byte(job.Banner), &arrBanner)
	if err != nil {
		return res, err
	}
	res = web.Job{
		Id:          job.Id,
		CategoryId:  job.CategoryId,
		CompanyLogo: job.CompanyLogo,
		CompanyName: job.CompanyName,
		Location:    job.Location,
		Title:       job.Title,
		Type:        job.Type,
		Banner:      arrBanner,
		Description: job.Description,
		Email:       job.Email,
		WebsiteUrl:  job.WebsiteUrl,
		CreatedAt:   job.UpdatedAt,
		UpdatedAt:   job.UpdatedAt,
	}
	return res, nil
}

func (s *JobServiceImpl) DeleteJob(ctx context.Context, jobId string) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = s.JobRepository.DeleteJob(ctx, tx, jobId)
	if err != nil {
		return err
	}
	return nil
}
