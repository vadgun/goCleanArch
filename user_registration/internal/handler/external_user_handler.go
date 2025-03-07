package handler

import (
	"encoding/json"
	"net/http"

	"github.com/vadgun/goApp/user_registration/internal/entity"
	"github.com/vadgun/goApp/user_registration/internal/usecase"
)

type ExternalUserHandler struct {
	fetchUserUseCase *usecase.FetchExternalUserUseCase
}

func NewExternalUserHandler(useCase *usecase.FetchExternalUserUseCase) *ExternalUserHandler {
	return &ExternalUserHandler{
		fetchUserUseCase: useCase,
	}
}

func (h *ExternalUserHandler) FetchUserHandler(w http.ResponseWriter, r *http.Request) {

	randomUser, err := h.fetchUserUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response := map[string]any{
		"results": []entity.ExternalUser{*randomUser},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
