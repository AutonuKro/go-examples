package main

import (
	"os"

	"github.com/AutonuKro/go-examples/rest-api/database"
	"github.com/joho/godotenv"
)

func main() {
	err := initApp()
	if err != nil {
		panic(err)
	}
	defer database.CloseMongoDB()
	app := generateApp()
	port := os.Getenv("PORT")
	app.Listen(port)
}

func initApp() error {
	err := loadENV()
	if err != nil {
		return err
	}
	err = database.StartMongoDB()
	if err != nil {
		return err
	}

	return nil
}

func loadENV() error {
	error := godotenv.Load()
	if error != nil {
		return error
	}

	return nil
}
