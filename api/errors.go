package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	if apiError, ok := err.(Error); ok {
		if errors.Is(apiError, ErrUnauthorized()) {
			fmt.Println("unauthorized")
			// return c.Redirect("/main", fiber.StatusSeeOther)
		}
		return c.Status(apiError.Code).JSON(apiError)
	}
	apiError := NewError(http.StatusInternalServerError, err.Error())
	return c.Status(apiError.Code).JSON(apiError)
}

type Error struct {
	Code int    `json:"code"`
	Err  string `json:"error"`
}

func (e Error) Error() string {
	return e.Err
}

func NewError(code int, err string) Error {
	return Error{
		Code: code,
		Err:  err,
	}
}

func ErrNotFound() Error {
	return Error{
		Code: http.StatusNotFound,
		Err:  "not found",
	}
}

func ErrPrivateInternal() Error {
	return Error{
		Code: http.StatusInternalServerError,
		Err:  "something went wrong",
	}
}

func ErrAlreadyExists() Error {
	return Error{
		Code: http.StatusBadRequest,
		Err:  "already exists",
	}
}
func ErrNotExist() Error {
	return Error{
		Code: http.StatusBadRequest,
		Err:  "command doesn't exists",
	}
}

func ErrInvalidData() Error {
	return Error{
		Code: http.StatusBadRequest,
		Err:  "invalid data",
	}
}

func ErrInvalidCredentials() Error {
	return Error{
		Code: http.StatusBadRequest,
		Err:  "invalid credentials",
	}
}

func ErrUnauthorized() Error {
	return Error{
		Code: http.StatusForbidden,
		Err:  "unauthorized",
	}
}
