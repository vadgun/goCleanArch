package usecase

import (
	"errors"

	"github.com/vadgun/goApp/user_registration/internal/repository"
	"github.com/vadgun/goApp/user_registration/pkg"
)

type LoginUseCase struct {
	repo repository.UserRepository
}

func NewLoginUseCase(repo repository.UserRepository) *LoginUseCase {
	return &LoginUseCase{repo: repo}
}

func (uc *LoginUseCase) Login(email, password string) (string, error) {
	user, err := uc.repo.GetByEmailOrUser(email)
	if err != nil {
		return "", err
	}

	if user.Password != password {
		return "", errors.New("usuario/contraseña incorrectos")
	}

	token, err := pkg.GenerateJWT(user.Email)
	if err != nil {
		return "", errors.New("error generando token")
	}

	return token, nil
}
