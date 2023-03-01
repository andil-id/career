package router

import (
	"career/controller"
	"career/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(job controller.JobController, auth controller.AuthController, jobCategory controller.JobCategoryController) *gin.Engine {
	// gin.SetMode(config.GinMode())
	router := gin.Default()
	// router.Use(middleware.Logging())
	router.Use(middleware.ErrorAppHandler())

	api := router.Group("/api")
	{
		api.POST("/job", job.CreateJob)
		api.GET("/jobs", job.GetAllJob)
		api.GET("/job/:job-id", job.GetJobDetail)
		api.DELETE("/job/:job-id", job.DeleteJob)
		api.PATCH("/job", job.UpdateJob)
		api.PATCH("/job-categories", jobCategory.GetAllJobCategory)
		api.POST("/auth/login", auth.Login)
	}
	return router
}
