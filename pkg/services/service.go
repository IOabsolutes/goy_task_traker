package services

import "todo_api/pkg/repositories"

type User interface {
}

type TaskItem interface {
}

type TaskList interface {
}

type Authorization interface {
}

type Service struct {
	User
	TaskList
	TaskItem
	Authorization
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{}
}
