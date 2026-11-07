package repositories

import (
	"context"
	"rest-go-api/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

// contrutor que usa o * para indicar que db é um ponteiro para pgxpool.Pool, permitindo que o repositório acesse o pool de conexões do banco de dados
func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Get(ctx context.Context, id int) ([]models.User, error) {
	query := "SELECT id,name, age, weight, height FROM users WHERE id=$1"
	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	var users []models.User
	for rows.Next() {
		var user models.User
		// o Scan é usado para ler os valores das colunas da linha atual do resultado da consulta e armazená-los nas variáveis correspondentes
		// o & é usado para passar o endereço das variáveis, permitindo que o Scan modifique diretamente os valores dessas variáveis
		err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Weight, &user.Height)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
