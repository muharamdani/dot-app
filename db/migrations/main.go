package migrations

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var m *gormigrate.Gormigrate

func Execute(db *gorm.DB) {
	// uuid-ossp extension is required for uuid_generate_v4() function
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	MigrateArticlesTable(db)
	MigrateCommentsTable(db)
	MigrateAddParentIdInComments(db)
	log.Println("---Migration success---")
}
