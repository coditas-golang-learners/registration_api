package main

import (
	"testing"

	Register "boilerplate/model/register"
	valid "boilerplate/util/Validation"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestValidateUser(t *testing.T) {
	tests := []struct {
		name    string
		request Register.Register
		want    bool
	}{
		{
			name: "Valid username",
			request: Register.Register{
				Username: "john.doe",
			},
			want: true,
		},
		{
			name: "Invalid username",
			request: Register.Register{
				Username: "john$doe",
			},
			want: false,
		},
		{
			name: "Valid password",
			request: Register.Register{
				Password: "Abc123!",
			},
			want: true,
		},
		{
			name: "Invalid password",
			request: Register.Register{
				Password: "password123",
			},
			want: false,
		},
		{
			name: "Valid PAN and Aadhaar",
			request: Register.Register{
				Pan:   "ABCDE1234F",
				Adhar: "123456789012",
			},
			want: true,
		},
		{
			name: "Invalid PAN and Aadhaar",
			request: Register.Register{
				Pan:   "ABCDE1234G",
				Adhar: "12345678901",
			},
			want: false,
		},
		{
			name: "Invalid PAN and empty Aadhaar",
			request: Register.Register{
				Pan:   "ABCDE1234G",
				Adhar: "",
			},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := validator.New()
			v.RegisterValidation("UserValidation", valid.ValidateUser)
			err := v.Struct(test.request)
			got := (err == nil)
			assert.Equal(t, test.want, got)
		})
	}
}
