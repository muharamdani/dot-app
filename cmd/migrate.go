package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	conn "dot-app/db"
	"dot-app/db/migrations"
	"dot-app/utils"

	"github.com/spf13/cobra"
)

var migrationTemplate = `package migrations

import (
	"fmt"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// Migrate Change this to your own migration function
func Migrate(db *gorm.DB) {
	m = gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "%s",
			Migrate: func(tx *gorm.DB) error {
				type TableName struct {}
				return tx.AutoMigrate(&TableName{})
			},
		},
	})
	if err := m.Migrate(); err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Migrated %s\n")
}
`

// migrateCmd represents the migrating schema
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "This command will migrate all migrations func inside the migrations folder.",
	Run: func(cmd *cobra.Command, args []string) {
		fileName, _ := cmd.Flags().GetString("create")
		if fileName != "" {
			if err := createMigrationFile(fileName); err != nil {
				log.Fatalf("Error creating migration file: %v", err)
				return
			}
			fmt.Printf("Migration file %s created!\n", fileName)
			os.Exit(0)
		}
		conn.Connect()

		migrations.Execute(conn.DB)
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.PersistentFlags().String("create", "", "Create migration file")
}

func createMigrationFile(filename string) error {
	t := time.Now()
	currTime := fmt.Sprintf("%d%02d%02d%02d%02d%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	migrationTemplate = fmt.Sprintf(migrationTemplate, currTime, filename)

	rootPath := utils.GetRootPath()

	fileName := fmt.Sprintf("%s/db/migrations/%s_%s.go", rootPath, currTime, filename)
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err2 := f.WriteString(migrationTemplate)

	if err2 != nil {
		return err
	}

	return nil
}
