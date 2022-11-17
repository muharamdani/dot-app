package repositories

import (
	"dot-app/pkg/comments/models"
	"dot-app/pkg/comments/requests"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

func Create(db *gorm.DB, out *models.Comment) error {
	return db.Create(&out).Error
}

func Patch(db *gorm.DB, out *requests.Patch, id uuid.UUID) error {
	return db.Model(&out).Where("id = ?", id).Updates(&out).Error
}

func Delete(db *gorm.DB, id uuid.UUID) error {
	return db.Where("id = ?", id).Delete(&models.Comment{}).Error
}
