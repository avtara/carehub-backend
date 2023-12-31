package models

import "errors"

var (
	ErrorUserNotFound       = errors.New("user not found")
	ErrorUserDuplicate      = errors.New("email has been taken")
	ErrorUserWrongPassword  = errors.New("wrong password or email")
	ErrorCategoryNotFound   = errors.New("category not found")
	ErrorComplainNotFound   = errors.New("complain not found")
	ErrorResolutionNotFound = errors.New("resolution not found")
)

const (
	GroupFilePhotoUser = "photo_user"
	MaxSize            = 1024 * 1024
)

type (
	SendEmailNewUserParams struct {
		Email    string `json:"email,omitempty"`
		Password string `json:"password,omitempty"`
	}
)
