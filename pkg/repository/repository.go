package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/p12s/2gis-catalog-api"
	todo "github.com/p12s/ispring-todo-list-api"
)

// Rubric
type Rubric interface {
	Create(rubric common.Rubric) (int, error)
	GetAll(rubricId int) ([]common.Rubric, error)
	GetById(rubricId int) (common.Rubric, error)
	Delete(rubricId int) error
}

type Street interface {
	Create(street common.Street) (int, error)
	GetById(streetId int) (common.Street, error)
	Delete(streetId int) error
}

type Phone interface {
	Create(phone common.Phone) (int, error)
	GetById(phoneId int) (common.Phone, error)
	Delete(phoneId int) error
}


type Building interface {
	Create(phone common.Phone) (int, error) // добавление в справочник здания вместе с организациями, которые в ней находятся
}

// Company - компания
type Company interface {
	Create(listId int, item todo.TodoItem) (int, error)
	GetById(userId, itemId int, input todo.UpdateItemInput) error // выдачу информации об организациях по их идентификаторам
	GetAllInBuilding(userId, listId int) ([]todo.TodoItem, error) // выдачу всех организаций находящихся в конкретном здании
	GetAllByRubricId(userId, itemId int) (todo.TodoItem, error) // выдачу списка всех организаций, которые относятся к указанной рубрике
	Delete(userId, itemId int) error
}

// UserAction - интерфейс для удаления пользователей
type UserAction interface {
	Delete(userId int) error
}

// Repository - репозитрий
type Repository struct {
	Authorization
	TodoList
	TodoItem
	UserAction
}

// NewRepository - конструктор
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
		UserAction:    NewUserPostgres(db),
	}
}
