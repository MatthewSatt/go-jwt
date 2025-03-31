package controllers

import (
	"net/http"
	"fmt"

	"github.com/MatthewSatt/starter/initializers"
	"github.com/MatthewSatt/starter/models"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User

	result := initializers.DB.Select("id, created_at, username, email").Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to fetch users",
		})
		return
	}

	var userList []gin.H
	for _, user := range users {
		userList = append(userList, gin.H{
			"id":        user.ID,
			"email":     user.Email,
			"createdAt": user.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"users": userList,
	})
}

func GetUserById(c *gin.Context) {
	userId := c.Param("userId")
	fmt.Printf("Looking for questions for user ID: %s\n", userId)

	var user models.User
	result := initializers.DB.Select("id, created_at, username, email").First(&user, "id = ?", userId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":        user.ID,
			"email":     user.Email,
			"createdAt": user.CreatedAt,
		},
	})
}
