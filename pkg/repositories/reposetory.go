package repositories

type User interface {
}

type TaskItem interface {
}

type TaskList interface {
}

type Authorization interface {
}

type Repository struct {
	User
	TaskList
	TaskItem
	Authorization
}

func NewRepository() *Repository {
	return &Repository{}
}
