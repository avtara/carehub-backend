package service

import (
	"context"
	"github.com/avtara/carehub/internal/models"
)

type AuthenticationUseCase interface {
	AuthenticationGoogle(ctx context.Context) (redirectURL string, err error)
	GoogleCallback(ctx context.Context, code string) (response models.AuthenticationResponse, err error)
	Login(ctx context.Context, args models.AuthenticationParams) (response models.AuthenticationResponse, err error)
}

type UserUseCase interface {
	UpdateProfile(ctx context.Context, args models.UpdateProfileUserParams, userID int64) (err error)
}