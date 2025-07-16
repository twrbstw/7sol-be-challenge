package errors

import "github.com/gofiber/fiber/v2"

// for handler layer
const (
	ERR_HANDLER_PARSING_REQ         = "err:handler:parsing_req"
	ERR_HANDLER_NAME_OR_EMAIL_EMPTY = "err:handler:name_or_email_is_empty"
)

//for service layer
const (
	ERR_SERVICE_HASHING = "err:service:hashing_failed"
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
	default:
		return fiber.StatusInternalServerError
	}
}
