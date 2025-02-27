package repository_test

import (
	"testing"

	"github.com/vadgun/goApp/user_registration/internal/entity"
	"github.com/vadgun/goApp/user_registration/internal/repository"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository(t *testing.T) {
	repo := repository.NewUserRepository()

	testUsers := map[string]entity.User{
		"byEmail": {
			User:     "Jose",
			Password: "Passwords1$$",
			Email:    "test@example.com",
			Phone:    "0000101010",
		},
		"byUsername": {
			User:     "Roberto",
			Password: "Passwords1$$",
			Email:    "anothertest@example.com",
			Phone:    "1000101010",
		},
		"byPhone": {
			User:     "Camacho",
			Password: "Passwords1$$",
			Email:    "another@example.com",
			Phone:    "0000101010",
		},
	}

	tests := []struct {
		name            string
		setup           func() error
		query           string
		expectedError   bool
		expectedErrMsg  string
		validateResults func(*testing.T, *entity.User, error)
	}{
		{
			name: "Guardar usuario por email",
			setup: func() error {
				return repo.Save(testUsers["byEmail"])
			},
			query:         "test@example.com",
			expectedError: false,
			validateResults: func(t *testing.T, user *entity.User, err error) {
				assert.NoError(t, err)
				assert.Equal(t, testUsers["byEmail"].Email, user.Email)
			},
		},
		{
			name: "Guardar usuario por usuario",
			setup: func() error {
				return repo.Save(testUsers["byUsername"])
			},
			query:         "Roberto",
			expectedError: false,
			validateResults: func(t *testing.T, user *entity.User, err error) {
				assert.NoError(t, err)
				assert.Equal(t, testUsers["byUsername"].User, user.User)
			},
		},
		{
			name: "No permitir correos duplicados",
			setup: func() error {
				return repo.Save(testUsers["byEmail"])
			},
			expectedError:  true,
			expectedErrMsg: "el correo ya se encuentra registrado",
		},
		{
			name: "No permitir telefonos duplicados",
			setup: func() error {
				return repo.Save(testUsers["byPhone"])
			},
			expectedError:  true,
			expectedErrMsg: "el teléfono ya se encuentra registrado",
		},
		{
			name:           "Buscar usuario inexistente",
			query:          "unknown@example.com",
			expectedError:  true,
			expectedErrMsg: "usuario/contraseña incorrectos",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setup != nil {
				err := tc.setup()
				if tc.expectedError {
					assert.Error(t, err)
					if tc.expectedErrMsg != "" {
						assert.Equal(t, tc.expectedErrMsg, err.Error())
					}
					return
				}
				assert.NoError(t, err)
			}

			if tc.query != "" {
				user, err := repo.GetByEmailOrUser(tc.query)
				if tc.validateResults != nil {
					tc.validateResults(t, user, err)
				} else if tc.expectedError {
					assert.Error(t, err)
					if tc.expectedErrMsg != "" {
						assert.Equal(t, tc.expectedErrMsg, err.Error())
					}
				} else {
					assert.NoError(t, err)
				}
			}
		})
	}
}
