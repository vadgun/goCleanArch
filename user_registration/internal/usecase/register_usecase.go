package usecase

import (
	"errors"
	"regexp"
	"unicode"

	"github.com/vadgun/goApp/user_registration/internal/entity"
	"github.com/vadgun/goApp/user_registration/internal/repository"
)

type RegisterUseCase struct {
	repo repository.UserRepository
}

func NewRegisterUseCase(repo repository.UserRepository) *RegisterUseCase {
	return &RegisterUseCase{repo: repo}
}

func validatePassword(password string) bool {
	if len(password) < 6 || len(password) > 12 {
		return false
	}

	allowedChars := regexp.MustCompile(`^[A-Za-z0-9@$&]+$`)
	if !allowedChars.MatchString(password) {
		return false
	}

	var hasLower, hasUpper, hasDigit, hasSpecial bool

	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsDigit(char):
			hasDigit = true
		case char == '@' || char == '$' || char == '&':
			hasSpecial = true
		}
	}

	return hasLower && hasUpper && hasDigit && hasSpecial
}

func validateEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(regex, email)
	return matched
}

func validatePhone(phone string) bool {
	regex := `^\d{10}$`
	matched, _ := regexp.MatchString(regex, phone)
	return matched
}

func (uc *RegisterUseCase) Register(user entity.User) error {

	if !validateEmail(user.Email) {
		return errors.New("correo inválido")
	}

	if !validatePhone(user.Phone) {
		return errors.New("teléfono inválido")
	}

	if !validatePassword(user.Password) {
		return errors.New("contraseña no cumple los requisitos")
	}

	if uc.repo.ExistsByEmailOrPhone(user.Email, user.Phone) {
		return errors.New("el correo o teléfono ya está registrado")
	}

	return uc.repo.Save(user)
}
