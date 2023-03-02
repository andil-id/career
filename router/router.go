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
	router.Use(middleware.CORS())
	router.Use(middleware.ErrorAppHandler())

	api := router.Group("/api")
	{
		api.POST("/job", middleware.JwtAuth(), job.CreateJob)
		api.GET("/jobs", job.GetAllJob)
		api.GET("/job/:job-id", job.GetJobDetail)
		api.DELETE("/job/:job-id", middleware.JwtAuth(), job.DeleteJob)
		api.PATCH("/job", middleware.JwtAuth(), job.UpdateJob)
		api.GET("/job-categories", jobCategory.GetAllJobCategory)
		api.POST("/auth/login", auth.Login)
	}
	return router
}
