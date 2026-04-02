package main

import (
	"github.com/igor-izvekov/crm/pkg/database"
	"github.com/igor-izvekov/crm/pkg/migrations"
)

func main() {
	if err := database.Connect("crm.db"); err != nil {
		panic(err)
	}

	db := database.GetDB()
	if err := migrations.AutoMigrate(db); err != nil {
		panic(err)
	}
}