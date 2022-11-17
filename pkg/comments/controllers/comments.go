package controllers

import (
	"dot-app/db"
	"dot-app/pkg/comments/models"
	"dot-app/pkg/comments/repositories"
	"dot-app/pkg/comments/requests"
	"dot-app/pkg/comments/services"
	"dot-app/utils/http"
	"github.com/gofrs/uuid"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var req requests.Create
	var model models.Comment

	if err := http.RequestValidation(c, &req); err != nil {
		return
	}

	if err := services.Create(&req, &model); err != nil {
		http.ResponseBadRequest(c, "Invalid payload", err)
		return
	}

	if err := repositories.Create(db.DB, &model); err != nil {
		http.ResponseInternalServerError(c, "Something went wrong", err)
		return
	}

	http.ResponseCreated(c, "Comment created successfully", model)
}

func Patch(c *gin.Context) {
	var comment requests.Patch

	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		http.ResponseBadRequest(c, "Invalid ID", nil)
		return
	}

	if err := http.RequestValidation(c, &comment); err != nil {
		return
	}

	if err := repositories.Patch(db.DB, &comment, id); err != nil {
		http.ResponseInternalServerError(c, "Something went wrong", err)
		return
	}

	http.ResponseOK(c, "Comment updated successfully", nil)
}

func Delete(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		http.ResponseBadRequest(c, "Invalid payload", err)
		return
	}

	if err := repositories.Delete(db.DB, id); err != nil {
		http.ResponseInternalServerError(c, "Something went wrong", err)
		return
	}

	http.ResponseOK(c, "Delete success", nil)
}
