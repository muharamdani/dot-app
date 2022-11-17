package requests

import (
	"github.com/gofrs/uuid"
)

type Create struct {
	ArticleId *uuid.UUID `json:"article_id"`
	ParentId  *uuid.UUID `json:"parent_id"`
	Content   string     `json:"content" binding:"required"`
}

type Patch struct {
	Content string `json:"content" binding:"required"`
}

func (Create) TableName() string {
	return "comments"
}

func (Patch) TableName() string {
	return "comments"
}
