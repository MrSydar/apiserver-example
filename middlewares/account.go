package middlewares

import (
	"context"
	"errors"
	"fmt"
	"mrsydar/apiserver/configs/constants/contextnames"
	"mrsydar/apiserver/configs/database"
	"mrsydar/apiserver/models"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AssociateAccountWithRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		email := claims["email"].(string)
		if email == "" {
			return errors.New("email value is empty in the received JWT token")
		}

		account := models.Account{}
		if err := database.Collections.Accounts.FindOne(context.Background(), bson.M{"email": email}).Decode(&account); err != nil {
			if err == mongo.ErrNoDocuments {
				account.ID = primitive.NewObjectID()
				account.Email = email
				result, err := database.Collections.Accounts.InsertOne(context.Background(), account)
				if err != nil {
					return fmt.Errorf("failed to insert account resource: %v", err)
				}
				account.ID = result.InsertedID.(primitive.ObjectID)
			} else {
				return fmt.Errorf("failed to get account: %v", err)
			}
		}

		c.Set(contextnames.AccountID, account.ID)

		return next(c)
	}
}
