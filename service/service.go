package service

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/teohrt/cruddyAPI/dbclient"
	"github.com/teohrt/cruddyAPI/entity"
)

type Service interface {
	CreateProfile(ctx context.Context, profile entity.ProfileData) (entity.CreateProfileResult, error)
	GetProfile(ctx context.Context, profileID string) (entity.Profile, error)
	UpdateProfile(ctx context.Context, profile entity.ProfileData, profileID string) error
	DeleteProfile(ctx context.Context, profileID string) error
}

type ServiceImpl struct {
	Client dbclient.Client
	Logger *zerolog.Logger
}

func New(config *dbclient.Config) Service {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	client := dbclient.New(config, &logger)
	return ServiceImpl{
		Client: client,
		Logger: &logger,
	}
}

type ProfileNotFoundError struct {
	msg string
}

func (e ProfileNotFoundError) Error() string {
	return e.msg
}

type ProfileAlreadyExistsError struct {
	msg string
}

func (e ProfileAlreadyExistsError) Error() string {
	return e.msg
}

type EmailIncsonsistentWithProfileIDError struct {
	msg string
}

func (e EmailIncsonsistentWithProfileIDError) Error() string {
	return e.msg
}
