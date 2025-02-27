package usecase_test

import (
	"testing"

	"github.com/vadgun/goApp/user_registration/internal/entity"
	"github.com/vadgun/goApp/user_registration/internal/repository/mocks"
	"github.com/vadgun/goApp/user_registration/internal/usecase"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	useCase := usecase.NewRegisterUseCase(mockRepo)

	testUsers := map[string]entity.User{
		"valid": {
			Email:    "test@example.com",
			Phone:    "1234567890",
			Password: "Password123$",
		},
		"duplicatePhone": {
			Email:    "test1@example.com",
			Phone:    "1234567890",
			Password: "Password123$",
		},
		"invalidPassword": {
			Email:    "new@example.com",
			Phone:    "0987654321",
			Password: "123",
		},
		"invalidEmail": {
			Email:    "newexample.com",
			Phone:    "0987654321",
			Password: "Password123$",
		},
		"invalidPhone": {
			Email:    "new@example.com",
			Phone:    "098765",
			Password: "Password123$",
		},
	}

	tests := []struct {
		name           string
		user           entity.User
		setupMock      func()
		expectedError  bool
		expectedErrMsg string
	}{
		{
			name: "Registro exitoso",
			user: testUsers["valid"],
			setupMock: func() {
				mockRepo.EXPECT().
					ExistsByEmailOrPhone(testUsers["valid"].Email, testUsers["valid"].Phone).
					Return(false)
				mockRepo.EXPECT().
					Save(gomock.Any()).
					Return(nil)
			},
			expectedError: false,
		},
		{
			name: "Teléfono ya registrado",
			user: testUsers["duplicatePhone"],
			setupMock: func() {
				mockRepo.EXPECT().
					ExistsByEmailOrPhone(testUsers["duplicatePhone"].Email, testUsers["duplicatePhone"].Phone).
					Return(true)
			},
			expectedError:  true,
			expectedErrMsg: "el correo o teléfono ya está registrado",
		},
		{
			name: "Correo ya registrado",
			user: testUsers["valid"],
			setupMock: func() {
				mockRepo.EXPECT().
					ExistsByEmailOrPhone(testUsers["valid"].Email, testUsers["valid"].Phone).
					Return(true)
			},
			expectedError:  true,
			expectedErrMsg: "el correo o teléfono ya está registrado",
		},
		{
			name:           "Contraseña no válida",
			user:           testUsers["invalidPassword"],
			expectedError:  true,
			expectedErrMsg: "contraseña no cumple los requisitos",
		},
		{
			name:           "Correo no válido",
			user:           testUsers["invalidEmail"],
			expectedError:  true,
			expectedErrMsg: "correo inválido",
		},
		{
			name:           "Teléfono no válido",
			user:           testUsers["invalidPhone"],
			expectedError:  true,
			expectedErrMsg: "teléfono inválido",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			if tc.setupMock != nil {
				tc.setupMock()
			}

			err := useCase.Register(tc.user)

			if tc.expectedError {
				assert.Error(t, err, "Se esperaba un error pero no ocurrió")
				if tc.expectedErrMsg != "" {
					assert.Equal(t, tc.expectedErrMsg, err.Error(),
						"El mensaje de error no coincide con lo esperado")
				}
			} else {
				assert.NoError(t, err, "No se esperaba un error pero ocurrió")
			}
		})
	}
}
