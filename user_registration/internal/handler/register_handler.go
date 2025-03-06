package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/vadgun/goApp/user_registration/internal/entity"
	"github.com/vadgun/goApp/user_registration/internal/usecase"
	"github.com/vadgun/goApp/user_registration/pkg"
)

type UserHandler struct {
	useCase *usecase.RegisterUseCase
}

func NewUserHandler(useCase *usecase.RegisterUseCase) *UserHandler {
	return &UserHandler{useCase: useCase}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var userReq entity.User
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	missingFields := validateRequiredRegisterFields(userReq)
	if len(missingFields) > 0 {
		msg := pkg.CreateMissingFieldsMessage(missingFields)
		respondWithError(w, msg, http.StatusBadRequest)
		return
	}

	if err := h.useCase.Register(userReq); err != nil {
		respondWithError(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	log.Println("Usuario registrado correctamente", fmt.Sprintf("%+v", userReq))

	json.NewEncoder(w).Encode(map[string]string{
		"Mensaje": "Usuario registrado correctamente",
	})
}

func respondWithError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}

func validateRequiredRegisterFields(registerReq entity.User) []string {
	var missingFields []string

	if registerReq.User == "" {
		missingFields = append(missingFields, "usuario")
	}
	if registerReq.Email == "" {
		missingFields = append(missingFields, "correo")
	}
	if registerReq.Password == "" {
		missingFields = append(missingFields, "contraseña")
	}
	if registerReq.Phone == "" {
		missingFields = append(missingFields, "telefono")
	}

	return missingFields
}
