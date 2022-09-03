package service

import (
	todo "github.com/Nikita-Kuzhl/go-rest-api"
	"github.com/Nikita-Kuzhl/go-rest-api/package/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int,error)
	GenerateToken(login,password string) (string,error)
}

type TodoList interface {

}

type TodoItem interface {
	
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}