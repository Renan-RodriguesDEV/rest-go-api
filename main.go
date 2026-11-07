package main

import (
	"log"
	"net/http"
	"rest-go-api/internal/db"
	"rest-go-api/internal/handlers"
	"rest-go-api/internal/repositories"
	"rest-go-api/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	// carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error in load env file!", err)
	}
	// inicia a conexão com db para que o repositories use esse pool de conexão
	db, err := db.Connect()
	if err != nil {
		log.Fatal("error in init connection!", err)
	}
	// garante que a conexão com o banco de dados seja fechada ao final da execução do programa
	defer db.Close()
	// inicializa o repositório, serviço e handler do usuário
	userRepository := repositories.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	// inicializa o roteador e define a rota para o endpoint de usuários
	router := chi.NewRouter()
	// define a rota para o endpoint de usuários
	router.Get("/users", userHandler.Get)
	log.Printf("Server init in :8080")
	// inicia o servidor HTTP na porta 8080 e registra o roteador como manipulador de solicitações
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("error in init server!", err)
	}
}
