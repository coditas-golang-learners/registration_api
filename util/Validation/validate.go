package Validate_User

import (
	MysqlWrapper "boilerplate/library/mysql"
	Register "boilerplate/model/register"
	"database/sql"
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateUser(fl validator.FieldLevel) bool {
	request, ok := fl.Top().Interface().(Register.Register)
	fmt.Println("f1 : ", fl.FieldName())
	if !ok {
		return false
	}
	if fl.FieldName() == "Username" {
		pattern := "^[a-zA-Z][a-zA-Z0-9._]+$"
		match, _ := regexp.MatchString(pattern, request.Username)
		return match
	}

	if fl.FieldName() == "Password" {
		var (
			hasUppercase   = regexp.MustCompile(`[A-Z]`).MatchString
			hasLowercase   = regexp.MustCompile(`[a-z]`).MatchString
			hasSpecialChar = regexp.MustCompile(`[!@#$%^&*()]`).MatchString
			hasNumeric     = regexp.MustCompile(`[0-9]`).MatchString
		)
		// Check if the password meets the requirements
		return hasUppercase(request.Password) && hasLowercase(request.Password) && hasSpecialChar(request.Password) && hasNumeric(request.Password)
	}

	if request.Pan == "" && request.Adhar == "" {
		return false
	}

	// At least one of the fields is not empty, perform additional validation
	if request.Pan != "" || request.Adhar != "" {
		pan_pattern := "^[A-Z]{5}[0-9]{4}[A-Z]{1}$"
		match_pan, _ := regexp.MatchString(pan_pattern, request.Pan)

		adhar_pattern := "^[0-9]{12}$"
		match_adhar, _ := regexp.MatchString(adhar_pattern, request.Adhar)
		return (match_adhar && len(request.Adhar) == 12) || (match_pan && len(request.Pan) == 10)
	}
	return true
}

func Check_username(user Register.Register) bool {
	// Check if the MysqlWrapper.Client object is nil
	if MysqlWrapper.Client == nil {
		fmt.Println("MySQL client is not initialized")
		return false
	}

	// Check if the username already exists
	var existingUser Register.Register
	err := MysqlWrapper.Client.QueryRow("SELECT username,FROM registers WHERE username = ?", user.Username).Scan(&existingUser.Username, &existingUser.Password)

	// Handle the error
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println(err)
		}
		return false
	}

	// Username already exists
	return true
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

func SendRequestValidationError(RequstPayload Register.Register) string {

	// Create a new validator instance
	validate := validator.New()

	// Register custom validation functions
	validate.RegisterValidation("UserValidation", ValidateUser)
	errMsg := ""
	if err := validate.Struct(RequstPayload); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := err.Field()
			switch fieldName {
			case "Username":
				switch err.Tag() {
				case "required":
					errMsg += "Username is empty. "
				case "min":
					errMsg += "Username should have a minimum length of 5. "
				case "max":
					errMsg += "Username should have a maximum length of 10. "
				case "alphanum":
					errMsg += "Username should only contain alphanumeric characters. "
				case "UserValidation":
					errMsg += "Username should not start with a special character or number. "
				}

			case "Firstname":
				switch err.Tag() {
				case "required":
					errMsg += "Firstname is empty. "
				case "alpha":
					errMsg += "Firstname should only contain alphabets characters. "
				}

			case "Lastname":
				switch err.Tag() {
				case "alpha":
					errMsg += "Lastname should only contain alphabets characters. "
				}

			case "Email":
				switch err.Tag() {
				case "email":
					errMsg += "Please enter correct email like abc@gmail.com. "
				}

			case "Mobile":
				switch err.Tag() {
				case "required":
					errMsg += "Mobile number is empty. "
				case "len=10|len=12":
					errMsg += "Mobile number should have a min length of 10 and max 12. "
				case "numeric":
					errMsg += "Please enter correct mobile number. "
				}

			case "Pan":
				switch err.Tag() {
				case "UserValidation":
					errMsg += "Please enter correct Pan number like ABCDE1234F or Adhar 12 digit numeric. "
				}

			case "Password":
				switch err.Tag() {
				case "required":
					errMsg += "Password is empty. "
				case "min":
					errMsg += "Password should have a minimum length of 5. "
				case "max":
					errMsg += "Password should have a maximum length of 30. "
				case "UserValidation":
					errMsg += "password should contain special character digit uppercase and lowercase. "
				}

			}
		}

	}

	return errMsg
}

func Check_username_login(user Register.Login) bool {
	// Check if the MysqlWrapper.Client object is nil
	if MysqlWrapper.Client == nil {
		fmt.Println("MySQL client is not initialized")
		return false
	}

	// Check if the username already exists
	var existingUser Register.Login
	err := MysqlWrapper.Client.QueryRow("SELECT username, password FROM registers WHERE username=? AND password=?", user.Username, user.Password).Scan(&existingUser.Username, &existingUser.Password)

	// Handle the error
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println(err)
		}
		return false
	}

	// Username already exists
	return true
}
