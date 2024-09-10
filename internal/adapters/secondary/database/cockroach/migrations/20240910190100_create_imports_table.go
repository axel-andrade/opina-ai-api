package cockroach_migrations

import (
	"log"

	cockroach_models "github.com/axel-andrade/opina-ai-api/internal/adapters/secondary/database/cockroach/models"
	"gorm.io/gorm"
)

func MigrateCreateImport(tx *gorm.DB) error {
	if !tx.Migrator().HasTable(&cockroach_models.ImportModel{}) {
		if err := tx.AutoMigrate(&cockroach_models.ImportModel{}); err != nil {
			return err
		}
		log.Printf("Migration executed: 20240910190100_create_imports_table")
	}

	return nil
}

func RollbackCreateImportTable(tx *gorm.DB) error {
	err := tx.Migrator().DropTable(&cockroach_models.ImportModel{})
	if err != nil {
		return err
	}

	return nil
}
