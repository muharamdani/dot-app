package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RequestValidation(c *gin.Context, obj any) error {
	err := c.ShouldBindJSON(obj)
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			ResponseUnprocessableEntity(c, "Invalid payload!", err)
		default:
			ResponseBadRequest(c, "Request not in json payload!", err)
		}
	}
	return err
}

func GetValidationResult(err error) map[string]string {
	errResult := make(map[string]string)
	for _, v := range err.(validator.ValidationErrors) {
		errResult[v.Field()] = v.Error()
	}
	return errResult
}
