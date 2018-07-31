package migration

import (
	"github.com/FelixMH/ProComments/configuration"
	"github.com/FelixMH/ProComments/models"
)

// Función para crear las tablas en la BD
func Migrate() {
	db := configuration.GetConnection()
	defer db.Close()

	db.CreateTable(&models.User{})
	db.CreateTable(&models.Comment{})
	db.CreateTable(&models.Vote{})

	db.Model(&models.Vote{}).AddUniqueIndex("comment_id_user_id_unique", "comment_id", "user_id")
}
