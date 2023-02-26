package controller

import (
	"career/helper"
	"career/model/web"
	"career/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JobControllerImpl struct {
	JobService service.JobService
}

func NewJobController(jobService service.JobService) JobController {
	return &JobControllerImpl{
		JobService: jobService,
	}
}

func (cl *JobControllerImpl) CreateJob(c *gin.Context) {
	req := web.CreateJob{}
	err := c.Bind(&req)
	if err != nil {
		c.Error(err)
		return
	}

	res, err := cl.JobService.CreateJob(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}
	helper.ResponseSuccess(c, res, helper.Meta{})
}

func (cl *JobControllerImpl) GetAllJob(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")
	categoryId := c.Query("categoryId")
	companyName := c.Query("companyName")

	res, pagination, err := cl.JobService.GetAllJob(c.Request.Context(), companyName, categoryId, limit, offset)
	if err != nil {
		c.Error(err)
		return
	}
	helper.ResponseSuccess(c, res, helper.Meta{
		Pagination: pagination,
		StatusCode: http.StatusOK,
	})
}

func (cl *JobControllerImpl) GetJobDetail(c *gin.Context) {
	jobId := c.Param("job-id")
	res, err := cl.JobService.GetJobDetail(c.Request.Context(), jobId)
	if err != nil {
		c.Error(err)
		return
	}
	helper.ResponseSuccess(c, res, helper.Meta{
		StatusCode: http.StatusOK,
	})
}

func (cl *JobControllerImpl) DeleteJob(c *gin.Context) {
	jobId := c.Param("job-id")
	err := cl.JobService.DeleteJob(c.Request.Context(), jobId)
	if err != nil {
		c.Error(err)
		return
	}
	helper.ResponseSuccess(c, nil, helper.Meta{})
}
