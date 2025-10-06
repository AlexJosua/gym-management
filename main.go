package main

import (
	"gym-membership/config"
	"gym-membership/models"
	"gym-membership/routes"
)

func main() {
	config.InitDB()

	// Auto migrate
	config.DB.AutoMigrate(&models.User{}, &models.Membership{})

	r := routes.SetupRouter()
	r.Run(":8080")
}
