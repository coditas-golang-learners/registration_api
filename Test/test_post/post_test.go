package validation_test

import (
	Register "boilerplate/model/register"
	valid "boilerplate/util/Validation"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type test struct {
	name    string
	request Register.Register
	want    bool
}

func TestRegistration_form_username(t *testing.T) {
	testcase := test{
		name: "ValidRegistration_user",
		request: Register.Register{
			Username:  "Pratigya",
			Firstname: "pratigya",
			Lastname:  "123456",
			Email:     "vivek@gmail.com",
			Pan:       "",
			Adhar:     "123456789012",
			Mobile:    "6158911234",
			Password:  "Prat@235687",
		},
		want: true,
	}

	t.Run(testcase.name, func(t *testing.T) {
		v := validator.New()
		v.RegisterValidation("UserValidation", valid.ValidateUser)
		err := v.Struct(testcase.request)

		got := (err == nil)
		assert.Equal(t, testcase.want, got)
	})

}

func TestRegistration_form_invalid_username(t *testing.T) {
	testcase := test{
		name: "In_validRegistration_user",
		request: Register.Register{
			Username:  "_Pratigya",
			Firstname: "pratigya",
			Lastname:  "123456",
			Email:     "vivek@gmail.com",
			Pan:       "",
			Adhar:     "123456789012",
			Mobile:    "6158911234",
			Password:  "Prat@235687",
		},
		want: false,
	}

	t.Run(testcase.name, func(t *testing.T) {
		v := validator.New()
		v.RegisterValidation("UserValidation", valid.ValidateUser)
		err := v.Struct(testcase.request)

		got := (err == nil)
		assert.Equal(t, testcase.want, got, "Validation result does not match for test: %s", testcase.name)
	})

}

func TestRegistration_form_pan_adhar(t *testing.T) {
	testcase := test{
		name: "ValidRegistration_pan_adhar",
		request: Register.Register{
			Username:  "Pratigya",
			Firstname: "pratigya",
			Lastname:  "123456",
			Email:     "vivek@gmail.com",
			Pan:       "",
			Adhar:     "123456789012",
			Mobile:    "6158911234",
			Password:  "Prat@235687",
		},
		want: true,
	}

	t.Run(testcase.name, func(t *testing.T) {
		v := validator.New()
		v.RegisterValidation("UserValidation", valid.ValidateUser)
		err := v.Struct(testcase.request)

		got := (err == nil)
		assert.Equal(t, testcase.want, got)
	})

}

func TestRegistration_form_invalid_pan_adhar(t *testing.T) {
	testcase := test{
		name: "In_validRegistration_pan_adhar",
		request: Register.Register{
			Username:  "_Pratigya",
			Firstname: "pratigya",
			Lastname:  "123456",
			Email:     "vivek@gmail.com",
			Pan:       "",
			Adhar:     "qwer68-098-0",
			Mobile:    "6158911234",
			Password:  "Prat@235687",
		},
		want: false,
	}

	t.Run(testcase.name, func(t *testing.T) {
		v := validator.New()
		v.RegisterValidation("UserValidation", valid.ValidateUser)
		err := v.Struct(testcase.request)

		got := (err == nil)
		assert.Equal(t, testcase.want, got, "Validation result does not match for pan or adhar test: %s", testcase.name)
	})

}

func TestRegistration_form_password(t *testing.T) {
	testcase := test{
		name: "ValidRegistration_password",
		request: Register.Register{
			Username:  "Pratigya",
			Firstname: "pratigya",
			Lastname:  "123456",
			Email:     "vivek@gmail.com",
			Pan:       "",
			Adhar:     "123456789012",
			Mobile:    "6158911234",
			Password:  "Prat@235687",
		},
		want: true,
	}

	t.Run(testcase.name, func(t *testing.T) {
		v := validator.New()
		v.RegisterValidation("UserValidation", valid.ValidateUser)
		err := v.Struct(testcase.request)

		got := (err == nil)
		assert.Equal(t, testcase.want, got, "Validation result does not match for password test: %s", testcase.name)
	})

}

func TestRegistration_form_invalid_password(t *testing.T) {
	testcase := test{
		name: "In_validRegistration_password",
		request: Register.Register{
			Username:  "Pratigya",
			Firstname: "pratigya",
			Lastname:  "123456",
			Email:     "vivek@gmail.com",
			Pan:       "",
			Adhar:     "123456789012",
			Mobile:    "6158911234",
			Password:  "Prat235687",
		},
		want: false,
	}

	t.Run(testcase.name, func(t *testing.T) {
		v := validator.New()
		v.RegisterValidation("UserValidation", valid.ValidateUser)
		err := v.Struct(testcase.request)

		got := (err == nil)
		assert.Equal(t, testcase.want, got, "Validation result does not match for password test: %s", testcase.name)
	})

}
