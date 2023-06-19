package post

import (
	Register "boilerplate/model/register"
	validated "boilerplate/util/Validation"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostUserInfo(c *gin.Context) {

	//Language in header
	_, _ = c.Get("language")

	// Create a new Register object to hold the request data
	var RequstPayload Register.Register
	// Bind the request JSON data to the newReg object
	if err := c.ShouldBindJSON(&RequstPayload); err != nil {
		// Handle the binding error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate_error := validated.SendRequestValidationError(RequstPayload)
	if validate_error != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": validate_error})
		return
	}
	user_check := validated.Check_username(RequstPayload)
	if user_check {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	email_check := validated.Check_email(RequstPayload)
	if email_check {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already exists"})
		return
	}

	mobile_check := validated.Check_mobile(RequstPayload)
	if mobile_check {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mobile number already exists"})
		return
	}

	fmt.Println("request pay load ", RequstPayload)
	_, err := Register.InsertUser(RequstPayload)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save data"})
		return
	}

	//	id, _ := rowsInserted.LastInsertId()
	RequstPayload.ID = int(123)

	c.JSON(http.StatusCreated, RequstPayload)
}
