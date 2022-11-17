package repositories

import (
	"fmt"

	"dot-app/pkg/articles/models"
	"dot-app/pkg/articles/requests"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

func Create(db *gorm.DB, out *models.Article) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&out).Error; err != nil {
			return err
		}
		return nil
	})
}

func Index(db *gorm.DB, out *[]models.Article, paginate *requests.Paginate) error {
	db1 := db.Order(fmt.Sprintf("%s %s", paginate.OrderBy, paginate.Sort))

	if !paginate.Paginate {
		return db1.Find(&out).Error
	}

	return db1.Limit(paginate.PerPage).Offset(paginate.PerPage * (paginate.Page - 1)).Find(&out).Error
}

func Show(db *gorm.DB, out *requests.ArticleShow, id uuid.UUID) error {
	return db.Model(&out).Preload("Comments").First(out, id).Error
}

func Update(db *gorm.DB, out *models.Article, id uuid.UUID) error {
	return db.Model(&out).Where("id = ?", id).Updates(&out).Error
}

func Patch(db *gorm.DB, out *requests.ArticlePatch, id uuid.UUID) error {
	return db.Model(&out).Where("id = ?", id).Updates(&out).Error
}

func Delete(db *gorm.DB, id uuid.UUID) error {
	return db.Where("id = ?", id).Delete(&models.Article{}).Error
}

func Count(db *gorm.DB, count *int64) error {
	return db.Model(&models.Article{}).Count(count).Error
}
