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

	jobService := service.NewJobService(jobRepository, db, validate)

	jobController := controller.NewJobController(jobService)

	router := router.NewRouter(jobController)
	router.Run(":" + config.AppPort())
}
