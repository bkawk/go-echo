package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

// HealthGet handles health check requests
func HealthGet(c echo.Context) error {
	db := c.Get("db").(*mongo.Client)
	newUser := User{Username: "bkawk", Password: "secret"}
	result, err := db.Database("test").Collection("test").InsertOne(context.TODO(), newUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	return c.JSON(http.StatusOK, map[string]string{
		"status": "OK",
	})
}
