package http

import (
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Paginate struct {
	Data        interface{} `json:"data"`
	Total       int64       `json:"total"`
	PerPage     int64       `json:"per_page"`
	CurrentPage int64       `json:"current_page"`
	TotalPages  int64       `json:"total_pages"`
}

func ResponseOK(c *gin.Context, msg string, data interface{}) {
	result := Response{
		Status:  true,
		Message: msg,
		Data:    data,
	}
	c.JSON(http.StatusOK, result)
}

func ResponseCreated(c *gin.Context, msg string, data interface{}) {
	result := Response{
		Status:  true,
		Message: msg,
		Data:    data,
	}
	c.JSON(http.StatusCreated, result)
}

func ResponseInternalServerError(c *gin.Context, msg string, err error) {
	result := Response{
		Status:  false,
		Message: msg,
	}

	if err != nil {
		result.Data = err.Error()
	} else {
		result.Data = err
	}

	c.AbortWithStatusJSON(http.StatusInternalServerError, result)
}

func ResponseBadRequest(c *gin.Context, msg string, err error) {
	result := Response{
		Status:  false,
		Message: msg,
	}

	if err != nil {
		result.Data = err.Error()
	} else {
		result.Data = err
	}

	c.AbortWithStatusJSON(http.StatusBadRequest, result)
}

func ResponseNotFound(c *gin.Context, msg string) {
	result := Response{
		Status:  false,
		Message: msg,
		Data:    nil,
	}
	c.AbortWithStatusJSON(http.StatusNotFound, result)
}

func ResponseUnprocessableEntity(c *gin.Context, msg string, err error) {
	validationResult := GetValidationResult(err)
	result := Response{
		Status:  false,
		Message: msg,
		Data:    validationResult,
	}
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, result)
}

func ResponsePaginate(c *gin.Context, msg string, data interface{}, total int64, page int64, limit int64) {
	paginate := Paginate{
		Data:        data,
		Total:       total,
		CurrentPage: page,
		PerPage:     limit,
		TotalPages:  int64(math.Ceil(float64(total) / float64(limit))),
	}

	result := Response{
		Status:  true,
		Message: msg,
		Data:    paginate,
	}

	c.JSON(http.StatusOK, result)
}
