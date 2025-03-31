package initializers

import "github.com/MatthewSatt/starter/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
