package controllers

import (
	"dot-app/db"
	"dot-app/pkg/articles/models"
	"dot-app/pkg/articles/repositories"
	"dot-app/pkg/articles/requests"
	"dot-app/pkg/articles/services"
	"dot-app/utils/http"
	"github.com/gofrs/uuid"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var req requests.ArticleRequest
	var ar models.Article

	if err := http.RequestValidation(c, &req); err != nil {
		return
	}

	services.CreateOrUpdate(&req, &ar)

	if err := repositories.Create(db.DB, &ar); err != nil {
		http.ResponseInternalServerError(c, "Something went wrong", nil)
		return
	}

	http.ResponseCreated(c, "Create article success", ar)
}

func Index(c *gin.Context) {
	var paginate requests.Paginate
	var articles []models.Article
	var getCount int64

	if err := c.Bind(&paginate); err != nil {
		http.ResponseBadRequest(c, "Invalid payload", err)
		return
	}

	if err := repositories.Index(db.DB, &articles, &paginate); err != nil {
		http.ResponseInternalServerError(c, "Something went wrong", nil)
		return
	}

	if err := repositories.Count(db.DB, &getCount); err != nil {
		http.ResponseInternalServerError(c, "Something went wrong", nil)
		return
	}

	if paginate.Paginate == false {
		http.ResponsePaginate(c, "Get articles data success", articles, getCount, 1, getCount)
		return
	}

	http.ResponsePaginate(c, "Get articles data success",
		articles, getCount, int64(paginate.Page), int64(paginate.PerPage))
}

func Show(c *gin.Context) {
	var article requests.ArticleShow

	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		http.ResponseBadRequest(c, "Invalid payload", err)
		return
	}

	if err := repositories.Show(db.DB, &article, id); err != nil {
		http.ResponseInternalServerError(c, "Something went wrong", nil)
		return
	}

	http.ResponseOK(c, "Get article success", article)
}

func Update(c *gin.Context) {
	var req requests.ArticleRequest
	var ar models.Article

	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		http.ResponseBadRequest(c, "Invalid payload", err)
		return
	}

	if err := http.RequestValidation(c, &req); err != nil {
		return
	}

	services.CreateOrUpdate(&req, &ar)

	if err := repositories.Update(db.DB, &ar, id); err != nil {
		http.ResponseInternalServerError(c, "Something went wrong", nil)
		return
	}

	http.ResponseOK(c, "Update success", req)
}

func Patch(c *gin.Context) {
	var req requests.ArticlePatch

	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		http.ResponseBadRequest(c, "Invalid payload", err)
		return
	}

	if err := http.RequestValidation(c, &req); err != nil {
		return
	}

	if err := repositories.Patch(db.DB, &req, id); err != nil {
		http.ResponseInternalServerError(c, "Something went wrong", nil)
		return
	}

	http.ResponseOK(c, "Patch success", req)
}

func Delete(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		http.ResponseBadRequest(c, "Invalid payload", err)
		return
	}

	if err := repositories.Delete(db.DB, id); err != nil {
		http.ResponseInternalServerError(c, "Something went wrong", nil)
		return
	}

	http.ResponseOK(c, "Delete success", nil)
}
