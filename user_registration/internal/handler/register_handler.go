package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/vadgun/goApp/user_registration/internal/entity"
	"github.com/vadgun/goApp/user_registration/internal/usecase"
)

type UserHandler struct {
	useCase *usecase.RegisterUseCase
}

func NewUserHandler(useCase *usecase.RegisterUseCase) *UserHandler {
	return &UserHandler{useCase: useCase}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	var missingFields []string

	if user.User == "" {
		missingFields = append(missingFields, "usuario")
	}
	if user.Email == "" {
		missingFields = append(missingFields, "correo")
	}
	if user.Password == "" {
		missingFields = append(missingFields, "contraseña")
	}
	if user.Phone == "" {
		missingFields = append(missingFields, "telefono")
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

	err := h.useCase.Register(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	log.Println("Usuario registrado correctamente", fmt.Sprintf("%+v", user))

	w.Write([]byte("Usuario registrado correctamente"))
}
