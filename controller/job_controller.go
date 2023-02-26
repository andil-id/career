package controller

import "github.com/gin-gonic/gin"

type JobController interface {
	CreateJob(c *gin.Context)
	GetAllJob(c *gin.Context)
	DeleteJob(c *gin.Context)
	GetJobDetail(c *gin.Context)
}
