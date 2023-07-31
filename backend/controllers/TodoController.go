package controllers

import (
	"strconv"
	"github.com/gofiber/fiber/v2"
)

func (c *controller) GetTodo(ctx *fiber.Ctx) error {

	sess,_:=c.session.Get(ctx)
	sessionId:=sess.Get("sessionId").(string)
	userid:=c.GetUseridFromSessionMapping(sessionId)

	todos,err:=c.model.GetTodo(userid)

	if err!=nil{
		return ctx.Status(200).JSON(fiber.Map{
		"status":  200,
		"Message": "Failed",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  200,
		"Message": "success",
		"todos":todos,
		
	})
}
func (c *controller) InsertTodo(ctx *fiber.Ctx) error {

	
	sess,_:=c.session.Get(ctx)
	sessionId:=sess.Get("sessionId").(string)
	userid:=c.GetUseridFromSessionMapping(sessionId)
	text:=ctx.FormValue("text")



	todos,err:=c.model.InsertTodo(userid,text)

	if err!=nil {
			return ctx.Status(200).JSON(fiber.Map{
			"status":  200,
			"Message": "failed",
			"error":err,
			
			}) 
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  200,
		"Message": "success",
		"todos":todos,
		
	})
}
func (c *controller) UpdateTodo(ctx *fiber.Ctx) error {

	todoid,_:=strconv.Atoi(ctx.FormValue("todoid"))
	
	text:=ctx.FormValue("text")

	
	sess,_:=c.session.Get(ctx)
	sessionId:=sess.Get("sessionId").(string)
	userid:=c.GetUseridFromSessionMapping(sessionId)

	todos,err:=c.model.UpdateTodo(userid,todoid,text)

	if err !=nil{
		return ctx.Status(200).JSON(fiber.Map{
		"status"	:  200,
		"Message"	: "success",
		"errors"	:err,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status"	:  200,
		"Message"	: "success",
		"todos"		:todos,
		
	})
}
func (c *controller) DeleteTodo(ctx *fiber.Ctx) error {

	todoid,_:=strconv.Atoi(ctx.FormValue("todoid"))

	
	sess,_:=c.session.Get(ctx)
	sessionId:=sess.Get("sessionId").(string)
	userid:=c.GetUseridFromSessionMapping(sessionId)
	
	todos,err:=c.model.DeleteTodo(userid,todoid)

	if err!=nil{
		return ctx.Status(500).JSON(fiber.Map{
		"status"	:  500,
		"Message"	: "failed",
		"error"		:err,
		
		})	
	}

	return ctx.Status(201).JSON(fiber.Map{
		"status"	:  201,
		"Message"	: "success",
		"todos"		: todos,
		
	})
}




