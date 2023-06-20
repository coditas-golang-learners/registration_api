package post_signup

import (
	Register "boilerplate/model/register"
	validated "boilerplate/util/Validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var loginData Register.Login

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := validated.Check_username_login(loginData)
	if result {
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
	}
}
