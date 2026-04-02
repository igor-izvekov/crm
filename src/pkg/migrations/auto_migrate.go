package migrations

import (
	"log"

	"gorm.io/gorm"

	"github.com/igor-izvekov/crm/pkg/models"
)

func AutoMigrate(db *gorm.DB) error {
	log.Println("Запускаем миграции")

	err := db.AutoMigrate(
		&models.Client{},
		&models.Manager{},
		&models.RealEstateObject{},
		&models.Deal{},
		&models.Document{},
		&models.Payment{},
		&models.Commission{},
	)

	if err != nil {
		log.Printf("Ошибка миграции: %v", err)
		return err
	}

	log.Println("Завершение миграции")
	return nil
}

func DropAllTables(db *gorm.DB) error {
	err := db.Migrator().DropTable(
		&models.Commission{},
		&models.Payment{},
		&models.Document{},
		&models.Deal{},
		&models.RealEstateObject{},
		&models.Manager{},
		&models.Client{},
	)

	if err != nil {
		log.Printf("Ошибка удаления таблиц: %v", err)
		return err
	}

	return nil
}

func ResetDatabase(db *gorm.DB) error {
	if err := DropAllTables(db); err != nil {
		return err
	}
	return AutoMigrate(db)
}
