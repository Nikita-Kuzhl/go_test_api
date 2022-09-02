package service

import "github.com/Nikita-Kuzhl/go-rest-api/package/repository"

type Authorization interface {

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

func NewService(repos repository.Repository) *Service {
	return &Service{}
}