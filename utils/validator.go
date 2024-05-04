package utils

import (
	"regexp"

	"unicode"

	"github.com/ftsog/ecom/customerrors"
	"github.com/nyaruka/phonenumbers"
)

const (
	// accepted length
	userNameMinLength  = 6
	userNameMaxLength  = 20
	firstNameMaxLength = 100
	lastNameMaxLength  = 100
)

func ValidateUsername(username string) (*string, error) {
	if len(username) < userNameMinLength || len(username) > userNameMaxLength {
		return nil, customerrors.UserLengthError
	}

	return &username, nil
}

func ValidateEmail(email string) (*string, error) {

	emailRegexp := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegexp.MatchString(email) {
		return nil, customerrors.EmailInvalidError
	}

	return &email, nil

}

func ValidateFirstName(firstName string) (*string, error) {
	if len(firstName) > firstNameMaxLength {
		return nil, customerrors.NameLengthError
	}
	return &firstName, nil
}

func ValidateLastName(lastName string) (*string, error) {
	if len(lastName) > lastNameMaxLength {
		return nil, customerrors.NameLengthError
	}

	return &lastName, nil
}

func ValidatePassword(password string) (*string, error) {
	var (
		hasMin     = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	if len(password) == 0 {
		return nil, customerrors.ZeroPassError
	}

	if len(password) >= 8 {
		hasMin = true
	}

	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpper = true
		} else if unicode.IsLower(char) {
			hasLower = true
		} else if unicode.IsNumber(char) {
			hasNumber = true
		} else if unicode.IsPunct(char) {
			hasSpecial = true
		}
	}

	if !hasMin || !hasUpper || !hasLower || !hasNumber || !hasSpecial {
		return nil, customerrors.PasswordValidationError
	}

	return &password, nil
}

func IsEmpty(username, email, firstName, lastName, password, phoneNumber string) error {
	if username == "" || email == "" || firstName == "" || lastName == "" || password == "" || phoneNumber == "" {
		return customerrors.FieldRequiredError
	}

	return nil
}

func ProductEmpty(name, category string, price float64, description string) error {
	if name == "" || category == "" || price == 0.0 || description == "" {
		return customerrors.FieldRequiredError
	}

	return nil
}

func ValidatePhone(phone string) (*string, error) {
	phoneNumber, err := phonenumbers.Parse(phone, "NG")
	if err != nil {
		return nil, customerrors.PhoneValidationError
	}

	isValid := phonenumbers.IsValidNumber(phoneNumber)

	if !isValid {
		return nil, customerrors.PhoneValidationError
	}

	number := phonenumbers.Format(phoneNumber, phonenumbers.INTERNATIONAL)
	return &number, nil
}
