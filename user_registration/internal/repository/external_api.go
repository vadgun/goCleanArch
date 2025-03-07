package repository

import "github.com/vadgun/goApp/user_registration/internal/entity"

type ExternalAPIRepository interface {
	FetchUser() (*entity.ExternalUser, error)
}
