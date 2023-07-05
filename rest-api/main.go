package main

import (
	"context"

	"github.com/AutonuKro/go-examples/rest-api/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	err := initApp()
	if err != nil {
		panic(err)
	}
	defer database.CloseMongoDB()
	app := fiber.New()
	app.Post("/customer", func(c *fiber.Ctx) error {
		customer := bson.M{"name": "Hemoprobha", "lastName": "Wary", "money": 500}
		collection := database.GetCollection("customer")
		nDoc, err := collection.InsertOne(context.TODO(), customer)
		if err != nil {
			return c.Status(fiber.ErrInternalServerError.Code).SendString("Internal server error")
		}

		return c.JSON(nDoc)
	})

	app.Listen(":3000")
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
