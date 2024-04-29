package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ramzyrsr/domain/user"
)

func UserSetupRoutes(r *gin.Engine, userService *user.UserService) {
	r.POST("/users", func(c *gin.Context) {
		var user user.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if name or email is empty
		if user.Name == "" || user.Email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Name and email cannot be empty"})
			return
		}

		// Check if email already exists
		// existingUser, err := userService.GetUserByEmail(user.Email)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check email existence"})
		// 	return
		// }
		// if existingUser != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		// 	return
		// }

		// Generate access token (dummy token for demonstration)
		accessToken := "qwertyuiopasdfghjklzxcvbnm"

		// Assuming we set expiration to 8 hours from now
		expiration := time.Now().Add(8 * time.Hour)

		// Create user
		err := userService.CreateUser(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		// Prepare response JSON
		response := gin.H{
			"message": "User registered successfully",
			"data": gin.H{
				"name":        user.Name,
				"email":       user.Email,
				"accessToken": accessToken,
				"expiration":  expiration.Format(time.RFC3339), // Format expiration time as RFC3339
			},
		}

		// Return response JSON
		c.JSON(http.StatusCreated, response)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		user, err := userService.GetUserByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, user)
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		var user user.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user.ID = id // assuming the ID is in the JSON payload
		err := userService.UpdateUser(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}

		c.Status(http.StatusOK)
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		err := userService.DeleteUser(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}

		c.Status(http.StatusOK)
	})
}
