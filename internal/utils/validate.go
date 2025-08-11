package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator"
)

func validateNameOrInitials(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	// Regex pattern for initials (e.g., "J. K.")
	initialsRegex := regexp.MustCompile(`^([A-Z]\. )*[A-Z]\.$`)
	// Regex pattern for a regular name (e.g., "Doe")
	nameRegex := regexp.MustCompile(`^[A-Za-z]+$`)
	return initialsRegex.MatchString(value) || nameRegex.MatchString(value)
}
func passwordValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Check for at least one uppercase letter
	if match, _ := regexp.MatchString(`[A-Z]`, password); !match {
		return false
	}

	// Check for at least one lowercase letter
	if match, _ := regexp.MatchString(`[a-z]`, password); !match {
		return false
	}

	// Check for at least one digit
	if match, _ := regexp.MatchString(`[0-9]`, password); !match {
		return false
	}

	// Check for at least one special character
	if match, _ := regexp.MatchString(`[!@#\$%\^&\*]`, password); !match {
		return false
	}

	return true
}

// Function to check for leading and trailing spaces
func validateNoLeadingTrailingSpaces(fl validator.FieldLevel) bool {
	name := fl.Field().String()
	return !strings.HasPrefix(name, " ") && !strings.HasSuffix(name, " ")
}

// Function to check for repeating spaces
func validateNoRepeatingSpaces(fl validator.FieldLevel) bool {
	name := fl.Field().String()
	return !strings.Contains(name, "  ")
}
func Validate(value interface{}) error {
	
	validate := validator.New()
	validate.RegisterValidation("nameOrInitials", validateNameOrInitials)
	validate.RegisterValidation("password", passwordValidation)
	validate.RegisterValidation("no_leading_trailing_spaces", validateNoLeadingTrailingSpaces)
	validate.RegisterValidation("no_repeating_spaces", validateNoRepeatingSpaces)
	fmt.Printf("DEBUG: validating struct of type %T with value: %+v\n", value, value)
	err := validate.Struct(value)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			switch e.Tag() {
			case "required":
				return fmt.Errorf("%s is required", e.Field())
			case "email":
				return fmt.Errorf("%s is not a valid email address", e.Field())
			case "numeric":
				return fmt.Errorf("%s shouls contain only digits", e.Field())
			case "len":
				return fmt.Errorf("%s shouls have a length of %s", e.Field(), e.Param())
			case "min":
				return fmt.Errorf("%s shouls have a minimum length of %s", e.Field(), e.Param())
			case "excludesall":
				return fmt.Errorf("%s shouls not contain space", e.Field())
			case "nameOrInitials":
				return fmt.Errorf("%s should be either initials or a regular name", e.Field())
			case "password":
				return fmt.Errorf("%s should contain at least one uppercase letter, one lowercase letter, one digit, and one special character", e.Field())
			case "no_leading_trailing_spaces":
				return fmt.Errorf("%s should not have leading or trailing spaces", e.Field())
			case "no_repeating_spaces":
				return fmt.Errorf("%s  should not have repeating spaces", e.Field())
			case "max":
				return fmt.Errorf("%s exceeds the maximum length", e.Field())
			case "alpha":
				return fmt.Errorf("%s should contain only alphabetic characters", e.Field())
			case "gt":
				return fmt.Errorf("%s must be greater than zero", e.Field())
			default:
				return fmt.Errorf("validation error for field %s", e.Field())
			}
		}
	}
	return nil
}
