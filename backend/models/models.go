package models

import (
	"github.com/gamagamol/todos/backend/entity"
)



type Iuser interface{
	Login(username string,password string) (entity.User,error)
	register(usernamer string,password string) entity.User
	CountUser() int
}

type Itodo interface {
	GetTodo(userid int) ([]entity.Todo,error)
	InsertTodo(userid int ,text string)([]entity.Todo,error)
	DeleteTodo(userid int ,todoid int)([]entity.Todo,error)
	UpdateTodo(userid int ,todoid int , text string)([]entity.Todo,error)
	FindTodo(userid int,todoid int) int
}

type Imodel interface{
	Iuser
	Itodo
}

type model struct {
	user []entity.User
	todo []entity.Todo
}

func NewModel() *model{

	return &model{[]entity.User{},[]entity.Todo{}}
}