package repository

import (
	"fmt"

	todo "github.com/Nikita-Kuzhl/go-rest-api"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name,login,password) values($1,$2,$3) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Name, user.Login, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
func (r *AuthPostgres) GetUser(login, password string) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf("Select * from %s where login=$1 and password=$2", userTable)
	err := r.db.Get(&user, query, login, password)
	return user, err
}
