package models

import (
	"github.com/gamagamol/todos/backend/entity"
)



type Iuser interface{
	Login(username string,password string) error
}

type Imodel interface{
	Iuser
}

type model struct {
	user *[]entity.User
}

func NewModel(user *[]entity.User) *model{

	return &model{user}
}