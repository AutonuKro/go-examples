package handlers

import (
	"context"

	"github.com/AutonuKro/go-examples/rest-api/database"
	"github.com/AutonuKro/go-examples/rest-api/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type bookDTO struct {
	LibraryId string `json:"library_id" bson:"library_id"`
	Title     string `json:"title" bson:"title"`
	Author    string `json:"author" bson:"author"`
	ISBN      string `json:"isbn" bson:"isbn"`
}

func CreateBook(c *fiber.Ctx) error {
	aBookDTO := new(bookDTO)
	err := c.BodyParser(aBookDTO)
	if err != nil {
		return err
	}
	collection := database.GetCollection("libraries")
	oid, err := primitive.ObjectIDFromHex(aBookDTO.LibraryId)
	if err != nil {
		return err
	}
	book := models.Book{
		Title:  aBookDTO.Title,
		Author: aBookDTO.Author,
		ISBN:   aBookDTO.ISBN,
	}
	queryResult, err := collection.UpdateOne(context.TODO(), bson.M{"_id": oid}, bson.D{{Key: "$push", Value: bson.M{"books": book}}})
	if err != nil {
		return err
	}
	return c.JSON(queryResult.UpsertedID)
}
