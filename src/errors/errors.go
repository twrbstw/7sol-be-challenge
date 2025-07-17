package errors

import "github.com/gofiber/fiber/v2"

// for handler layer
const (
	ERR_HANDLER_PARSING_REQ         = "err:handler:parsing_req"
	ERR_HANDLER_NAME_OR_EMAIL_EMPTY = "err:handler:name_or_email_is_empty"
)

//for service layer
const (
	ERR_SERVICE_HASHING                     = "err:service:hashing_failed"
	ERR_SERVICE_INCORRECT_EMAIL_OR_PASSWORD = "err:service:incorrect_email_or_password"
	ERR_SERVICE_GENERATING_JWT_FAILED       = "err:service:generating_jwt_failed"
)

//for repository layer
const (
	ERR_USER_NOT_FOUND        = "err:user:not_found"
	ERR_USER_EMAIL_DUPLICATED = "err:user:email_duplicated"
)

func HandleErrResp(err error) int {
	switch err.Error() {
	case ERR_USER_NOT_FOUND:
		return fiber.StatusNotFound
	case ERR_USER_EMAIL_DUPLICATED:
		return fiber.StatusBadRequest
	case ERR_SERVICE_INCORRECT_EMAIL_OR_PASSWORD:
		return fiber.StatusUnauthorized
	default:
		return fiber.StatusInternalServerError
	}
}
