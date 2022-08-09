package controllers

import (
	"context"
	"encoding/json"
	"io"
	"mrsydar/apiserver/configs/constants/contextnames"
	"mrsydar/apiserver/configs/database"
	"mrsydar/apiserver/configs/log"
	"mrsydar/apiserver/models"
	"mrsydar/apiserver/responses"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetNotes(c echo.Context) error {
	notes := []models.Note{}

	accountID := c.Get(contextnames.AccountID).(primitive.ObjectID)
	cur, err := database.Collections.Notes.Find(context.Background(), bson.M{"owner": accountID})
	if err != nil {
		msg := "failed to list notes"
		log.Logger.Errorf("%v: %v", msg, err)
		return responses.Message(c, http.StatusInternalServerError, msg)
	}

	if err := cur.All(context.Background(), &notes); err != nil {
		msg := "failed to list notes"
		log.Logger.Errorf("%v: %v", msg, err)
		return responses.Message(c, http.StatusInternalServerError, msg)
	}

	return c.JSON(http.StatusOK, notes)
}

func PostNote(c echo.Context) error {
	var note models.Note

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		msg := "failed to post a note"
		log.Logger.Errorf("%v: %v", msg, err)
		return responses.Message(c, http.StatusBadRequest, msg)
	}

	if len(body) == 0 {
		msg := "request body is empty"
		return responses.Message(c, http.StatusBadRequest, msg)
	}

	if err := json.Unmarshal(body, &note); err != nil {
		msg := "failed to post a note"
		log.Logger.Errorf("%v: %v", msg, err)
		return responses.Message(c, http.StatusInternalServerError, msg)
	}

	note.ID = primitive.NewObjectID()
	note.Owner = c.Get(contextnames.AccountID).(primitive.ObjectID)

	result, err := database.Collections.Notes.InsertOne(context.Background(), note)
	if err != nil {
		msg := "failed to post a note"
		log.Logger.Errorf("%v: %v", msg, err)
		return responses.Message(c, http.StatusInternalServerError, msg)
	}

	note.ID = result.InsertedID.(primitive.ObjectID)
	note.Owner = primitive.NilObjectID

	return c.JSON(http.StatusCreated, note)
}
