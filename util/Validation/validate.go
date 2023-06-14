package Validate_User

import (
	MysqlWrapper "boilerplate/library/mysql"
	Register "boilerplate/model/register"
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
)

func ValidateUsername(fl validator.FieldLevel) bool {
	request, not_ok := fl.Top().Interface().(Register.Register)
	fmt.Println("still here", not_ok)
	if !not_ok {
		return false
	}
	fmt.Println("request.Username", request.Username)
	if request.Username == "" {
		return false
	}
	fmt.Println("request.Username2", request.Username)
	if strings.HasPrefix(request.Username, "_") {
		return false
	}

	specialCharRegex := `^[a-zA-Z][a-zA-Z0-9._]+$`
	match, _ := regexp.MatchString(specialCharRegex, request.Username)
	fmt.Println("match", match)
	if !match {
		return false
	}
	return true
}

func ValidateFirstname(fl validator.FieldLevel) bool {
	// Check if the username starts with a special character
	request, ok := fl.Top().Interface().(Register.Register)
	if !ok {
		return false
	}
	specialCharRegex := `^[a-zA-Z]+$`
	match, _ := regexp.MatchString(specialCharRegex, request.Firstname)
	if !match {
		//errors.New("firstname only contain alphabets")
		return false
	}

	return true
}

// Regular expression for strong password validation
func ValidatePassword(fl validator.FieldLevel) bool {

	request, ok := fl.Top().Interface().(Register.Register)
	if !ok {
		return false
	}
	// Check for at least one uppercase letter
	if strings.IndexFunc(request.Password, unicode.IsUpper) == -1 {
		//errors.New("Password should contain one uppper case")
		return false
	}

	// Check for at least one lowercase letter
	if strings.IndexFunc(request.Password, unicode.IsLower) == -1 {
		//errors.New("Password should contain one lower case")
		return false
	}

	// Check for at least one digit
	if strings.IndexFunc(request.Password, unicode.IsDigit) == -1 {
		//errors.New("Password should contain one digit")
		return false
	}

	specialCharRegex := `[@#$%^&+=]`
	match, _ := regexp.MatchString(specialCharRegex, request.Password)
	if !match {
		//errors.New("Password should contain atleast one special character case")
		return false
	}

	// Check for at least one non-alphanumeric character
	nonAlphaNumericRegex := `[^a-zA-Z0-9]`
	match, _ = regexp.MatchString(nonAlphaNumericRegex, request.Password)
	if !match {
		errors.New("alpha numeric")
		return false
	}

	return true
}

func Check_username(user Register.Register) bool {
	// Check if the username already exists
	var existingUser Register.Register
	err := MysqlWrapper.Client.QueryRow("SELECT username FROM registers WHERE username = ?", user.Username).Scan(&existingUser.Username)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		return false
	}

	if err == nil {
		// Username already exists
		return true
	}

	// Username not found
	return false
}

func Check_mobile(mobile Register.Register) bool {
	var existingUser Register.Register
	// Check if the mobile number already exists
	m_err := MysqlWrapper.Client.QueryRow("SELECT mobile FROM registers WHERE mobile = ?", mobile.Mobile).Scan(&existingUser.ID)
	if m_err != nil && m_err != sql.ErrNoRows {
		fmt.Println(m_err)
		return false

	}
	if m_err == nil {
		// mobile number already exists
		return true
	}

	// mobile number not found
	return false
}

func Check_email(email Register.Register) bool {
	var existingUser Register.Register
	// Check if the email already exists
	e_err := MysqlWrapper.Client.QueryRow("SELECT mobile FROM registers WHERE mobile = ?", email.Mobile).Scan(&existingUser.ID)
	if e_err != nil && e_err != sql.ErrNoRows {
		fmt.Println(e_err)
		return false
	}
	if e_err == nil {
		// email email already exists
		return true
	}

	// email not found
	return false
}
