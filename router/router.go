package router

import (
	"career/controller"
	"career/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(job controller.JobController) *gin.Engine {
	// gin.SetMode(config.GinMode())
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.ErrorAppHandler())

	api := router.Group("/api")
	{
		api.POST("/job", job.CreateJob)
		api.GET("/jobs", job.GetAllJob)
		api.GET("/job/:job-id", job.GetJobDetail)
		api.DELETE("/job/:job-id", job.DeleteJob)
	}
	return router
}
