package main

import (
	"log"

	"github.com/fasthttp/session/v2"
	"github.com/gamagamol/todos/backend/controllers"
	"github.com/gamagamol/todos/backend/entity"
	"github.com/gamagamol/todos/backend/models"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	model:=models.NewModel(&[]entity.User{})
	controller:=controllers.NewController(model,&session.Store{})

	api:=app.Group("/api")
	
	api.Post("/login",controller.Login)

    log.Fatal(app.Listen(":3000"))
}