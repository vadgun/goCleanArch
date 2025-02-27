package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vadgun/goApp/user_registration/internal/entity"
	"github.com/vadgun/goApp/user_registration/internal/handler"
	"github.com/vadgun/goApp/user_registration/internal/repository"
	"github.com/vadgun/goApp/user_registration/internal/usecase"

	"github.com/stretchr/testify/assert"
)

func TestLoginHandler(t *testing.T) {
	repo := repository.NewUserRepository()
	useCase := usecase.NewLoginUseCase(repo)
	handler := handler.NewLoginHandler(useCase)

	testUser := entity.User{
		User:     "Jose",
		Phone:    "0000101010",
		Email:    "test@example.com",
		Password: "Password123$",
	}

	err := repo.Save(testUser)
	assert.NoError(t, err, "Configuración de prueba fallida: no se pudo guardar el usuario")

	tests := []struct {
		name             string
		loginCredentials map[string]string
		expectedStatus   int
		validateResponse func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "Login exitoso por correo",
			loginCredentials: map[string]string{
				"correoOUsuario": "test@example.com",
				"contraseña":     "Password123$",
			},
			expectedStatus: http.StatusOK,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				var response map[string]interface{}
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				assert.NoError(t, err, "La respuesta debería ser JSON válido")
				assert.Contains(t, response, "token", "La respuesta debería contener un token")
				assert.NotEmpty(t, response["token"], "El token no debería estar vacío")
			},
		},
		{
			name: "Login exitoso por usuario",
			loginCredentials: map[string]string{
				"correoOUsuario": "Jose",
				"contraseña":     "Password123$",
			},
			expectedStatus: http.StatusOK,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				var response map[string]interface{}
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				assert.NoError(t, err, "La respuesta debería ser JSON válido")
				assert.Contains(t, response, "token", "La respuesta debería contener un token")
			},
		},
		{
			name: "Contraseña incorrecta",
			loginCredentials: map[string]string{
				"correoOUsuario": "test@example.com",
				"contraseña":     "ContraseñaIncorrecta$",
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "Usuario no encontrado",
			loginCredentials: map[string]string{
				"correoOUsuario": "test1",
				"contraseña":     "ContraseñaIncorrecta$",
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "Solicitud JSON inválida",
			loginCredentials: map[string]string{
				"correoOUsuario": "test@example.com",
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			requestBody, err := json.Marshal(tc.loginCredentials)
			assert.NoError(t, err, "Error al convertir credenciales a JSON")

			req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(requestBody))
			assert.NoError(t, err, "Error al crear la solicitud HTTP")
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			handler.Login(rr, req)

			assert.Equal(t, tc.expectedStatus, rr.Code, "El código de estado HTTP no coincide con lo esperado")

			if tc.validateResponse != nil {
				tc.validateResponse(t, rr)
			}
		})
	}
}
