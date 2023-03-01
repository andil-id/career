package controller

import "github.com/gin-gonic/gin"

type JobCategoryController interface {
	GetAllJobCategory(c *gin.Context)
}
