package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() (*pgxpool.Pool, error) {
	// le o valor da variável de ambiente DB_URL
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatalf("Is not possible connect to %s", dbUrl)
	}
	log.Printf("Using database url: %s", dbUrl)
	// cria um pool de conexões com o banco de dados usando a URL fornecida
	// context.Background() é usado para criar um contexto de fundo, que é um contexto sem prazo de expiração ou cancelamento
	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}
	// bate um "ping" no banco de dados para verificar se a conexão foi estabelecida com sucesso
	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, err
	}
	return pool, nil
}
