package routes

import (
	"log"
	"time"

	"github.com/gamagamol/todos/backend/controllers"
	"github.com/gamagamol/todos/backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var (
	store *session.Store
	
)

func Routes(){

	route := fiber.New()

	store=session.New(session.Config{
		CookieHTTPOnly: true,
		Expiration: time.Hour*5,
	})

	

	model:=models.NewModel()
	controller:=controllers.NewController(model,store)

	
	route.Post("/login",controller.Login)
	route.Get("/logout",controller.Logout)

	api:=route.Group("/api")

	api.Use(NewAuthMiddleware(),cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins: "*",
		AllowHeaders: "Acess-Control-Allow-Origin,Content-Type,Origin,Accept",
	}))

	api.Get("/todos",controller.GetTodo)
	api.Post("/todos",controller.InsertTodo)
	api.Delete("/todos",controller.DeleteTodo)
	api.Put("/todos",controller.UpdateTodo)




    log.Fatal(route.Listen(":3000"))

}