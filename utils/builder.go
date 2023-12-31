package utils

import (
	"errors"
	"github.com/avtara/carehub/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type Data struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
}

func Response(c echo.Context, statusCode int, data interface{}) error {
	return c.JSON(statusCode, data)
}

func MessageResponse(c echo.Context, statusCode int, data interface{}) error {
	message := map[int]string{
		http.StatusOK:           "Success",
		http.StatusBadRequest:   "Bad request",
		http.StatusUnauthorized: "Unauthorized",
	}

	return Response(c, statusCode, Data{
		Code:    statusCode,
		Message: message[statusCode],
		Data:    data,
	})
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	var (
		message = "Internal Server Error"
		code    = http.StatusInternalServerError
	)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	if strings.Contains(err.Error(), models.ErrorUserNotFound.Error()) {
		code = http.StatusNotFound
		message = "User not found"
		err = models.ErrorUserNotFound
	} else if strings.Contains(err.Error(), models.ErrorUserDuplicate.Error()) {
		code = http.StatusConflict
		message = "Email has been taken"
		err = models.ErrorUserDuplicate
	} else if strings.Contains(err.Error(), models.ErrorUserWrongPassword.Error()) {
		code = http.StatusUnprocessableEntity
		message = "Wrong password or email"
		err = models.ErrorUserWrongPassword
	} else if strings.Contains(err.Error(), models.ErrorCategoryNotFound.Error()) {
		code = http.StatusNotFound
		message = "Category not found"
		err = models.ErrorCategoryNotFound
	} else if strings.Contains(err.Error(), models.ErrorComplainNotFound.Error()) {
		code = http.StatusNotFound
		message = "Complain not found"
		err = models.ErrorComplainNotFound
	} else if strings.Contains(err.Error(), models.ErrorResolutionNotFound.Error()) {
		code = http.StatusUnprocessableEntity
		message = "Resolution not found"
		err = models.ErrorResolutionNotFound
	}

	err = errors.New("internal server error")

	c.JSON(code, Data{
		Code:    code,
		Message: message,
		Error:   err.Error(),
	})
}
