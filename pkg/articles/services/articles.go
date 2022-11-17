package services

import (
	"dot-app/pkg/articles/models"
	"dot-app/pkg/articles/requests"
)

func CreateOrUpdate(req *requests.ArticleRequest, ar *models.Article) {
	ar.Title = req.Title
	ar.Content = req.Content
	ar.Author = req.Author
}
