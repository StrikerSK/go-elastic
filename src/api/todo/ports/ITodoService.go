package ports

import todoDomain "github.com/strikersk/go-elastic/src/api/todo/domain"

type ITodoService interface {
	FindTodo(string) (todoDomain.Todo, error)
	FindTodos() ([]todoDomain.Todo, error)
	SearchTodos([]string) ([]todoDomain.Todo, error)
	CreateTodo(todoDomain.Todo) (string, error)
	UpdateTodo(string, todoDomain.Todo) error
	DeleteTodo(string) error
}
