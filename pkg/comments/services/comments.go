package services

import (
	"dot-app/pkg/comments/models"
	"dot-app/pkg/comments/requests"
	"errors"
)

func Create(req *requests.Create, out *models.Comment) error {
	if req.ArticleId == nil && req.ParentId == nil {
		return errors.New("article_id or parent_id is required")
	}

	if req.ArticleId != nil && req.ParentId != nil {
		return errors.New("must be either article_id or parent_id")
	}

	out.ArticleId = req.ArticleId
	out.ParentId = req.ParentId
	out.Content = req.Content

	return nil
}
