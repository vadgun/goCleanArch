package usecase_test

import (
	"testing"

	"github.com/vadgun/goApp/user_registration/internal/entity"
	"github.com/vadgun/goApp/user_registration/internal/repository"
	"github.com/vadgun/goApp/user_registration/internal/usecase"

	"github.com/stretchr/testify/assert"
)

func TestLoginUseCase(t *testing.T) {
	repo := repository.NewUserRepository()
	useCase := usecase.NewLoginUseCase(repo)

	testUser := entity.User{
		User:     "Jose",
		Email:    "test@example.com",
		Password: "Password123$",
	}

	err := repo.Save(testUser)
	assert.NoError(t, err, "")

	tests := []struct {
		name           string
		email          string
		password       string
		expectedError  bool
		expectedErrMsg string
		validateToken  func(*testing.T, string)
	}{
		{
			name:          "Login exitoso por correo",
			email:         "test@example.com",
			password:      "Password123$",
			expectedError: false,
			validateToken: func(t *testing.T, token string) {
				assert.NotEmpty(t, token, "El token no debe estar vacío en un login exitoso")
			},
		},
		{
			name:          "Login exitoso por usuario",
			email:         "Jose",
			password:      "Password123$",
			expectedError: false,
			validateToken: func(t *testing.T, token string) {
				assert.NotEmpty(t, token, "El token no debe estar vacío en un login exitoso")
			},
		},
		{
			name:           "Error en usuario no encontrado",
			email:          "wrong@example.com",
			password:       "Password123$",
			expectedError:  true,
			expectedErrMsg: "usuario/contraseña incorrectos",
		},
		{
			name:           "Error en contraseña incorrecta con usuario encontrado",
			email:          "test@example.com",
			password:       "WrongPass123$",
			expectedError:  true,
			expectedErrMsg: "usuario/contraseña incorrectos",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			token, err := useCase.Login(tc.email, tc.password)

			if tc.expectedError {
				assert.Error(t, err, "Se esperaba un error pero no ocurrió")
				if tc.expectedErrMsg != "" {
					assert.Equal(t, tc.expectedErrMsg, err.Error(),
						"El mensaje de error no coincide con lo esperado")
				}
				return
			}

			assert.NoError(t, err, "No se esperaba un error pero ocurrió")

			if tc.validateToken != nil {
				tc.validateToken(t, token)
			}
		})
	}
}
