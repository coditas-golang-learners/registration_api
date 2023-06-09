package post

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"

	MysqlWrapper "boilerplate/library/mysql"
	Register "boilerplate/model/register"
)

func isEmailValid(email string) bool {
	// Regular expression for email validation
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)
	return match
}

// Regular expression for strong password validation
func isStrongPassword(s string) bool {

	// Check for at least one uppercase letter
	if strings.IndexFunc(s, unicode.IsUpper) == -1 {
		return false
	}

	// Check for at least one lowercase letter
	if strings.IndexFunc(s, unicode.IsLower) == -1 {
		return false
	}

	// Check for at least one digit
	if strings.IndexFunc(s, unicode.IsDigit) == -1 {
		return false
	}

	specialCharRegex := `[@#$%^&+=]`
	match, _ := regexp.MatchString(specialCharRegex, s)
	if !match {
		return false
	}

	// Check for at least one non-alphanumeric character
	nonAlphaNumericRegex := `[^a-zA-Z0-9]`
	match, _ = regexp.MatchString(nonAlphaNumericRegex, s)
	if !match {
		return false
	}

	return true
}

func PostUserInfo(c *gin.Context) {
	// Check if the username already exists
	var newReg Register.Register
	var existingUser Register.Register
	err := MysqlWrapper.Client.QueryRow("SELECT * FROM registers WHERE username = ?", newReg.Username).Scan(&existingUser.ID)
	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user data"})
		return
	}
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	// Check if the mobile number already exists
	err = MysqlWrapper.Client.QueryRow("SELECT * FROM registers WHERE mobile = ?", newReg.Mobile).Scan(&existingUser.ID)
	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user data"})
		return
	}
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mobile number already exists"})
		return
	}

	if err := c.ShouldBindJSON(&newReg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Validation

	if newReg.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}

	if newReg.Firstname == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Firstname is required"})
		return
	}

	// Check if the first name contains only characters
	firstNameRegex := `^[a-zA-Z]+$`
	match, _ := regexp.MatchString(firstNameRegex, newReg.Firstname)
	if !match {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid first name"})
		return
	}

	// Check if the first name contains only characters
	LastNameRegex := `^[a-zA-Z]+$`
	exp, _ := regexp.MatchString(LastNameRegex, newReg.Lastname)
	if !exp {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Last name"})
		return
	}

	if newReg.Mobile == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mobile number is required"})
		return
	}

	if len(newReg.Mobile) != 12 && len(newReg.Mobile) != 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mobile number"})
		return
	}

	if newReg.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	// Check if the email is valid
	if !isEmailValid(newReg.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}

	if newReg.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is not provided"})
		return
	}

	// Check password length
	if len(newReg.Password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password should be at least 8 characters"})
		return
	}

	// Check if the password is alphanumeric
	if !isStrongPassword(newReg.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password should contain alphanumeric and special characters"})
		return
	}

	// Check if either Adhar or Pan has a value
	if newReg.Adhar == "" && newReg.Pan == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Either Adhar or Pan is required"})
		return
	}

	// Save the data into the database
	result, err := MysqlWrapper.Client.Exec("INSERT INTO registers (username, firstname, lastname, email, pan, adhar, mobile, password) VALUES (?, ?, ?, ?, ?, ?, ?,?)",
		newReg.Username, newReg.Firstname, newReg.Lastname, newReg.Email, newReg.Pan, newReg.Adhar, newReg.Mobile, newReg.Password)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save data"})
		return
	}

	id, _ := result.LastInsertId()
	newReg.ID = int(id)

	c.JSON(http.StatusCreated, newReg)
}

func LoginHandler(c *gin.Context) {
	var loginData Register.Register

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the user exists in the database
	var user Register.Register
	err := MysqlWrapper.Client.QueryRow("SELECT * FROM registers WHERE username = ?", loginData.Username).Scan(&user.ID, &user.Email, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user data"})
		return
	}

	// Compare the provided password with the stored password
	if loginData.Username != user.Username {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect Username"})
		return
	}

	if loginData.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}
