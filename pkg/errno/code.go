package errno

var (
	// OK ... Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	// user errors

	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}
	ErrToken      = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	// user errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user was not found."}
	ErrPasswordIncorrect = &Errno{Code: 20103, Message: "The password was incorrect."}
	ErrUserExist         = &Errno{Code: 20104, Message: "The user was already exists."}
	ErrTokenInvalid      = &Errno{Code: 20105, Message: "The token was invalid."}
	ErrTokenExpired      = &Errno{Code: 20106, Message: "The token was expired."}
)
