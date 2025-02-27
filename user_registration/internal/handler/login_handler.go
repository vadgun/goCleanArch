package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/vadgun/goApp/user_registration/internal/usecase"
)

type LoginHandler struct {
	useCase *usecase.LoginUseCase
}

func NewLoginHandler(useCase *usecase.LoginUseCase) *LoginHandler {
	return &LoginHandler{useCase: useCase}
}

func (h *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var request struct {
		EmailOrUsername string `json:"correoOUsuario"`
		Password        string `json:"contraseña"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	var missingFields []string

	if request.EmailOrUsername == "" {
		missingFields = append(missingFields, "correoOUsuario")
	}
	if request.Password == "" {
		missingFields = append(missingFields, "contraseña")
	}

	if len(missingFields) > 0 {
		msg := "Falta el campo "
		if len(missingFields) > 1 {
			msg = "Faltan los campos "
		}
		msg += strings.Join(missingFields, ", ")
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	token, err := h.useCase.Login(request.EmailOrUsername, request.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := struct {
		Token string `json:"token"`
	}{Token: token}

	log.Printf("%+v logeado satisfactoriamente, token enviado\n", request.EmailOrUsername)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
