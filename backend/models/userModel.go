package models

import (

	// "errors"


	"github.com/gamagamol/todos/backend/entity"
	"golang.org/x/crypto/bcrypt"
)

func (m *model) Login(username string, password string) (entity.User,error) {

	if len(m.user) >0 {
		for _,u := range m.user{
			if u.Username==username{
				if err:=bcrypt.CompareHashAndPassword([]byte(u.Password),[]byte(password));err!=nil{
					return u,err
				}else{
					return entity.User{},nil
				}
			}
		}
		return m.register(username,password),nil
	}else{
		return m.register(username,password),nil
	}


	// return entity.User{},nil 

}

func (m *model) register(username string, password string) entity.User{


    id := len(m.user) + 1
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	newUser := entity.User{
		UserId:   id,
		Username: username,
		Password: string(hash),
	}

	m.user = append(m.user, newUser)
	return newUser
}


func (m *model)CountUser() int{

	return len(m.user)
}

