package controllers

import (
	"github.com/fasthttp/session/v2"
	"github.com/gamagamol/todos/backend/models"
)

type controller struct {
	model models.Imodel
	totalLogin int
	sessions *session.Store
	todoCount int
	userCount int
	sessionMap map[string]int

}

func NewController(model models.Imodel,session *session.Store) *controller{

	return &controller{
		model: model,
		totalLogin: 0,
		sessions: session,
		todoCount: 0,
		userCount: 0,
		sessionMap: make(map[string]int),
	}


}