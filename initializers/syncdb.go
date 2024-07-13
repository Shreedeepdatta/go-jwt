package initializers

import "github.com/Shreedeepdatta/go-jwt/models"

func SyncDb() {
	DB.AutoMigrate(&models.User{})
}
