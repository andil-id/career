package middleware

import (
	"career/helper"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Logging() gin.HandlerFunc {
	file, err := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.SetOutput(file)
				log.Printf("Error: %v", err)
				log.SetOutput(os.Stdout)
				helper.ResponseError(c, http.StatusInternalServerError, "Internal Server Error")
			}
		}()
		c.Next()
	}

}
