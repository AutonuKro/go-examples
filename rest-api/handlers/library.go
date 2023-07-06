package handlers

import (
	"context"

	"github.com/AutonuKro/go-examples/rest-api/database"
	"github.com/AutonuKro/go-examples/rest-api/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type libraryDTO struct {
	Name    string `json:"name" bson:"name"`
	Address string `json:"address" bson:"address"`
}

func CreatLibrary(c *fiber.Ctx) error {
	aLibraryDTO := new(libraryDTO)
	if err := c.BodyParser(aLibraryDTO); err != nil {
		return err
	}
	collection := database.GetCollection("libraries")
	result, err := collection.InsertOne(context.TODO(), aLibraryDTO)
	if err != nil {
		return err
	}

	return c.JSON(map[string]interface{}{"id": result.InsertedID})
}

func GetLibraries(c *fiber.Ctx) error {
	collection := database.GetCollection("libraries")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return err
	}
	var libraries []models.Library
	err = cursor.All(context.TODO(), &libraries)
	if err != nil {
		return err
	}

	return c.JSON(libraries)
}
