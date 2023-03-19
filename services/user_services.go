package services

import (
	"go-rest-api/db"
	"go-rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type Users struct {
// 	UserId      string `json:"userId"`
// 	FirstName   string `json:"firstName"`
// 	LastName    string `json:"lastName"`
// 	PhoneNumber string `json:"phoneNumber"`
// }

// var users = []Users{
// 	{UserId: "1", FirstName: "Harry Potter", LastName: "J. K. Rowling", PhoneNumber: "0000001"},
// }

// func ListALLUser(c *gin.Context) {
// 	c.JSON(http.StatusOK, users)
// }

func ListALLUser(c *gin.Context) {
	// Query database for all users
	var users []models.User
	db.Connect()
	db.DB.Table("users").Select("user_id, first_name, last_name, mobile_phone").Scan(&users)

	// Return JSON response
	c.JSON(http.StatusOK, users)
}
