package main

import (
	"career/config"
	"career/controller"
	"career/repository"
	"career/router"
	"career/service"
	"log"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := config.Connection()
	validate := validator.New()

	jobRepository := repository.NewJobRespository()
	adminRepository := repository.NewAdminRepository()
	jobCategoryRepository := repository.NewJobCategoryRespository()

	jobService := service.NewJobService(jobRepository, db, validate)
	authService := service.NewAuthService(db, validate, adminRepository)
	jobCategoryService := service.NewJobCategoryService(db, jobCategoryRepository)

	jobController := controller.NewJobController(jobService)
	authController := controller.NewAuthController(authService)
	jobCategoryController := controller.NewJobCategoryController(jobCategoryService)

	router := router.NewRouter(jobController, authController, jobCategoryController)
	router.Run(":" + config.AppPort())
}
