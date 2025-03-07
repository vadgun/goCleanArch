package usecase

import (
	"github.com/vadgun/goApp/user_registration/internal/entity"
	"github.com/vadgun/goApp/user_registration/internal/repository"
)

type FetchExternalUserUseCase struct {
	externalrepo repository.ExternalAPIRepository
}

func NewFetchExternalUserUseCase(repo repository.ExternalAPIRepository) *FetchExternalUserUseCase {
	return &FetchExternalUserUseCase{
		externalrepo: repo,
	}
}

func (u *FetchExternalUserUseCase) Execute() (*entity.ExternalUser, error) {
	return u.externalrepo.FetchUser()
}
