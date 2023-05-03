package helper

import (
	"career/model/web"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Meta struct {
	Pagination any `json:"pagination,omitempty"`
	StatusCode int `json:"status_code"`
	Message    any `json:"message"`
}

func ResponseSuccess(c *gin.Context, data any, meta Meta) {
	metaData := Meta{}
	switch c.Request.Method {
	case "GET":
		_, isPagination := meta.Pagination.(web.Pagination)
		if isPagination {
			metaData = Meta{
				Pagination: meta.Pagination,
				StatusCode: meta.StatusCode,
				Message:    "Data was successfully retrieved!",
			}
		} else {
			metaData = Meta{
				Pagination: nil,
				StatusCode: meta.StatusCode,
				Message:    "Data was successfully retrieved!",
			}
		}

	case "DELETE":
		metaData = Meta{
			StatusCode: http.StatusNoContent,
			Message:    "Data was successfully deleted!",
		}
	case "POST", "PUT", "PATCH":
		metaData = Meta{
			StatusCode: meta.StatusCode,
			Message:    "Data was sunccesfully transmited!",
		}
	default:
		err := errors.New("htpp method not recognized")
		c.Error(err)
	}
	c.JSON(meta.StatusCode, gin.H{
		"data": data,
		"meta": metaData,
	})
}

func ResponseError(c *gin.Context, code int, message any) {
	meta := Meta{
		StatusCode: code,
		Message:    message,
	}
	c.JSON(code, gin.H{
		"data": nil,
		"meta": meta,
	})
	c.Abort()
}
