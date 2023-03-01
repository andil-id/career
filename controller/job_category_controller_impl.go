package controller

import (
	"career/helper"
	"career/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JobCategoryControllerImpl struct {
	JobCategoryService service.JobCategoryService
}

func NewJobCategoryController(jobCategoryService service.JobCategoryService) JobCategoryController {
	return &JobCategoryControllerImpl{
		JobCategoryService: jobCategoryService,
	}
}

func (cl *JobCategoryControllerImpl) GetAllJobCategory(c *gin.Context) {
	res, err := cl.JobCategoryService.GetAllJobCategory(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}
	helper.ResponseSuccess(c, res, helper.Meta{
		StatusCode: http.StatusOK,
	})
}
