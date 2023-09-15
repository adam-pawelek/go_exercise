package initializers

import (
	"github.com/adam-pawelek/go_exercise/tree/main/gin_and_gorm/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})

}
