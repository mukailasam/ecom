package customerrors

import "errors"

var (

	// internal server error
	InternalErrorMessage = "Internal Error"
	InternalErrorDetail  = "Sorry, something went wrong on our server, try again later"

	// session
	SessionExistsError  = errors.New("Session already exists")
	InvalidSessionError = errors.New("Invalid session")

	// username error messages
	UserLengthError = errors.New("Invalid Username")
	UserExistError  = errors.New("Username exists")

	// email error messages
	EmailExistError   = errors.New("Email exists")
	EmailInvalidError = errors.New("Invalid Email, enter a vailid email")

	// firstName error messages
	NameLengthError = errors.New("Invalid Name")

	// password error messages
	PasswordValidationError = errors.New("paword should be minimum 8 in length and Password should contain at least a single uppercase letter, lowercase letter, single digit and a special character")
	ZeroPassError           = errors.New("Password should not be empty")
	// required
	FieldRequiredError = errors.New("All field required")

	//Phone Validation Error
	PhoneValidationError = errors.New("Invalid Phone Number")
)
