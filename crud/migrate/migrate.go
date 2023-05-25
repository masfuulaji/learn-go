package main

import (
	"crud/initializers"
	"crud/models"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal("Migrate fail")
	}
}
