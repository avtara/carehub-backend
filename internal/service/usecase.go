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

type CategoryUseCase interface {
	GetAllCategories(ctx context.Context, limit int) (response []models.Category, err error)
	GetCategoryByID(ctx context.Context, ID int64) (response models.Category, err error)
}

type ComplainUseCase interface {
	GetAllComplain(ctx context.Context, limit int) (response []models.Complain, err error)
	GetComplainByID(ctx context.Context, ID int64) (response models.Complain, err error)
	InsertComplain(ctx context.Context, args models.InsertComplainParams, userID int64) (err error)
	InsertResolution(ctx context.Context, args models.InsertResolutionParams, complainID, adminID int64) (err error)
	UpdateStatus(ctx context.Context, status string, complainID int64) (err error)
}
