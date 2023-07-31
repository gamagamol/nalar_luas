package controllers

import (
	"github.com/gamagamol/todos/backend/entity"
	"github.com/gamagamol/todos/backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type IUserController interface {
	Login(c *fiber.Ctx) error
	validate(user entity.User) []entity.ErrorResponse
	Logout(c *fiber.Ctx)error
	GetUseridFromSessionMapping(sessionId string) int
}

type ITodoController interface {
	GetTodo(ctx *fiber.Ctx) error
	InsertTodo(ctx *fiber.Ctx) error
	UpdateTodo(ctx *fiber.Ctx) error
	DeleteTodo(ctx *fiber.Ctx) error
}

type Icontroller interface {
	IUserController 
	ITodoController
}

type controller struct {
	model models.Imodel
	
	totalLogin 	int
	todoCount 	int
	userCount 	int
	sessionMap 	[]entity.SessionMap
	session 	*session.Store

}

func NewController(model models.Imodel,store *session.Store) *controller{
	return &controller{
		model: model,
		totalLogin: 0,
		todoCount: 0,
		userCount: 0,
		session: store,
		sessionMap: []entity.SessionMap{},
	}
}