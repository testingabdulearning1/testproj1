package respcode

const (
	BindingError        = "BINDING_ERROR"
	ValidationError     = "VALIDATION_ERROR"
	Success             = "SUCCESS"
	InternalServerError = "INTERNAL_SERVER_ERROR"
	NotFound            = "NOT_FOUND"
	Unauthorized        = "UNAUTHORIZED"
	Forbidden           = "FORBIDDEN"
	Created             = "CREATED"
	NotMentioned        = "NOT_MENTIONED"
)

// Authentication(middleware) related response codes
const (
	TokenExpired = "TOKEN_EXPIRED"
	InvalidToken = "INVALID_TOKEN"
)

//@Team: Please add more response codes here (if needed) and update the documentation accordingly.
