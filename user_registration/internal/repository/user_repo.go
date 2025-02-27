package repository

import (
	"errors"

	"github.com/vadgun/goApp/user_registration/internal/entity"
)

type UserRepository interface {
	GetByEmailOrUser(email string) (*entity.User, error)
	ExistsByEmailOrPhone(email, phone string) bool
	Save(user entity.User) error
}

type userRepositoryImpl struct {
	usersByEmail map[string]entity.User
	usersByPhone map[string]entity.User
	usersByUser  map[string]entity.User
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{
		usersByEmail: make(map[string]entity.User),
		usersByPhone: make(map[string]entity.User),
		usersByUser:  make(map[string]entity.User),
	}
}

func (repo *userRepositoryImpl) Save(user entity.User) error {
	if _, exists := repo.usersByEmail[user.Email]; exists {
		return errors.New("el correo ya se encuentra registrado")
	}
	if _, exists := repo.usersByPhone[user.Phone]; exists {
		return errors.New("el teléfono ya se encuentra registrado")
	}

	repo.usersByEmail[user.Email] = user
	repo.usersByPhone[user.Phone] = user
	repo.usersByUser[user.User] = user
	return nil
}

func (repo *userRepositoryImpl) ExistsByEmailOrPhone(email, phone string) bool {
	_, emailExists := repo.usersByEmail[email]
	_, phoneExists := repo.usersByPhone[phone]
	return emailExists || phoneExists
}

func (repo *userRepositoryImpl) GetByEmailOrUser(emailOrUsername string) (*entity.User, error) {

	userByEmail, existsByEmail := repo.usersByEmail[emailOrUsername]
	userByUsername, existsByUsername := repo.usersByUser[emailOrUsername]

	if existsByEmail {
		return &userByEmail, nil
	}

	if existsByUsername {
		return &userByUsername, nil
	}

	return &entity.User{}, errors.New("usuario/contraseña incorrectos")
}
