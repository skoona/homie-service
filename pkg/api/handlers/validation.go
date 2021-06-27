package handlers

import (
	"fmt"
	validator "github.com/go-playground/validator/v10"
	"regexp"
)

// ValidationErrorMessage is a collection of validation error messages
type ValidationErrorMessage struct {
	Messages []string `json:"messages"`
}

// Validation contains
type Validation struct {
	validate *validator.Validate
}

//NewValidation creates a new Validation type
func NewValidation() *Validation {
	validator := validator.New()
	validator.RegisterValidation("fwfn", validateFwFn)
	validator.RegisterValidation("eid", ValidateEID)

	return &Validation{validate: validator}
}

// ValidateStruct verify contents of a structure object
//
// if ve, ok := err.(validator.ValidationErrors); ok {
//			fmt.Println(ve.Namespace())
//			fmt.Println(ve.Field())
//			fmt.Println(ve.StructNamespace())
//			fmt.Println(ve.StructField())
//			fmt.Println(ve.Tag())
//			fmt.Println(ve.ActualTag())
//			fmt.Println(ve.Kind())
//			fmt.Println(ve.Type())
//			fmt.Println(ve.Value())
//			fmt.Println(ve.Param())
//			fmt.Println()
//	}
func (v *Validation) ValidateStruct(i interface{}) ValidationErrorMessage {
	errs := v.validate.Struct(i)

	if errs == nil {
		return ValidationErrorMessage{}
	}

	if _, ok := errs.(*validator.InvalidValidationError); ok {
		fmt.Println(errs)
		return ValidationErrorMessage{Messages: []string{errs.Error()}}
	}

	msgs := []string{}
	for _, err := range errs.(validator.ValidationErrors) {
		// cast the FieldError into our ValidationError and append to the slice
		msgs = append(msgs, fmt.Sprintf(
			"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
			err.Namespace(),
			err.Field(),
			err.Tag(),
		))
	}
	return ValidationErrorMessage{Messages: msgs}
}

// ValidateParam validates individual params
// example: ValidateParam(email, "required,email")
// returns: []string
func (v *Validation) ValidateParam(param interface{}, tag string) ValidationErrorMessage {
	msgs := []string{}

	errs := v.validate.Var(param, tag)
	if errs != nil {
		msgs = append(msgs, errs.Error())
		// output: Key: "" Error:Field validation for "" failed on the "email" tag
	}

	return ValidationErrorMessage{Messages: msgs}
}

// validateFwFn verifies Firmware Filenames
// key: fwfn
func validateFwFn(fl validator.FieldLevel) bool {
	// Firmware File Name must be in the format 'alphaNum.bin'
	re := regexp.MustCompile(`[[:print:]]+\.bin`)
	items := re.FindAllString(fl.Field().String(), -1)

	return (len(items) == 1)
}

// ValidateEID verify EID uuid-like field
// key: eid
func ValidateEID(fl validator.FieldLevel) bool {
	// EID is a UUID with/out dashes, no more than 36 chars
	re := regexp.MustCompile(`[\-0-9A-Fa-f]{32,36}?`)
	items := re.FindAllString(fl.Field().String(), -1)

	return (len(items) == 1)
}
