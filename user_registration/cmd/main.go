package main

import (
	"log"
	"net/http"

	"github.com/vadgun/goApp/user_registration/internal/handler"
	"github.com/vadgun/goApp/user_registration/internal/repository"
	"github.com/vadgun/goApp/user_registration/internal/usecase"
)

func main() {
	repo := repository.NewUserRepository()

	registerUseCase := usecase.NewRegisterUseCase(repo)
	registerHandler := handler.NewUserHandler(registerUseCase)

	loginUseCase := usecase.NewLoginUseCase(repo)
	loginHandler := handler.NewLoginHandler(loginUseCase)

	http.HandleFunc("/register", registerHandler.Register)
	http.HandleFunc("/login", loginHandler.Login)

	log.Println("Servidor corriendo en :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
