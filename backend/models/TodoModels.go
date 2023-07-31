package models

import (
	"errors"

	"github.com/gamagamol/todos/backend/entity"
)

func (m *model) GetTodo(userid int) ([]entity.Todo,error){

	if len(m.todo) == 0 {
		return []entity.Todo{},errors.New("Todos are empty")
	}

	return m.todo,nil
}
func (m *model) InsertTodo(userid int ,text string) ([]entity.Todo,error){
	

	m.todo=append(m.todo, entity.Todo{
		UserId: userid,
		TodoId: len(m.todo)+1,
		Text: text,
	})
	return m.todo,nil
}
func (m *model) DeleteTodo(userid int ,todoid int) ([]entity.Todo,error){
	
	if len(m.todo) ==0{
		return []entity.Todo{},nil
	}

	var newTodo []entity.Todo

	for _,t:=range m.todo{
		if t.UserId==userid {
			if t.TodoId!=todoid{
				newTodo = append(newTodo, t)
			}
		}else{
			newTodo=append(newTodo, t)
		}
	}

		m.todo=newTodo
	

	return m.todo,nil
}
func (m *model) UpdateTodo(userid int ,todoid int,text string) ([]entity.Todo,error){
	
	if len(m.todo) ==0{
		return []entity.Todo{},nil
	}
	
	iOldTodo:=m.FindTodo(userid,todoid)
	m.todo[iOldTodo]=entity.Todo{
		UserId: userid,
		TodoId: todoid,
		Text: text,
	}
	return m.todo,nil
}

func (m *model)FindTodo(userid int,todoid int) int{

	for i,t:=range m.todo{
		if t.UserId==userid {
			if t.TodoId==todoid{
				return i
			}
		}
	}
	return 0
}


