package controllers

import (
	
	"time"

	"github.com/gamagamol/todos/backend/entity"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)
func (co *controller) Login(c *fiber.Ctx) error {
	
	var user entity.User

	sess,err:=co.session.Get(c)

	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":fiber.StatusBadRequest,
			"Message":"Wrong Password",
			"errors":err,
			
		})
	}

	// check apakah format json bener
	if err:=c.BodyParser(&user);err!=nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":err.Error(),
		})
	}
	// check validasinya bener
	if errors:=co.validate(user);errors!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(entity.ResponseError{
			Status: fiber.StatusBadRequest,
			Message: "error",
			Errors: errors,
		})
	}

	// check login 
	user,er:=co.model.Login(user.Username,user.Password)

	if er!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":fiber.StatusBadRequest,
			"Message":"Wrong Password",
			
		})
	}

	sessionId:=uuid.New().String()
	sess.Set("sessionId",sessionId)

	co.sessionMap=append(co.sessionMap, entity.SessionMap{
		SessionId: sessionId,
		UserId: user.UserId,
	})
	co.totalLogin=co.totalLogin+1
	co.userCount=co.model.CountUser()
	co.sessionMap=append(co.sessionMap, entity.SessionMap{
		SessionId: sessionId,
		UserId: user.UserId,
	})


	sess.Set("sessionId",sessionId)
	if err:=sess.Save();err!=nil{
		 return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":fiber.StatusBadRequest,
			"Message":"Wrong Password",
			"errors":err,
			
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:"sessionId",
		Value: sessionId,
		Expires: time.Now().Add(24*time.Hour),
		HTTPOnly: true,
	})

		

	// return kalo udh bener semua
	return c.Status(201).JSON(entity.ResponseSuccess{
		Status: 201,
		Message: "Success",
		// SessionId:sessionId,
	})
	
}

func (co *controller) validate(user entity.User) []entity.ErrorResponse{
	validate:=validator.New()
	var errors []entity.ErrorResponse
	if err:=validate.Struct(user);err!=nil{
		for _,err:= range err.(validator.ValidationErrors){
			var element entity.ErrorResponse
			element.Field=err.StructNamespace()
			element.Tag=err.ActualTag()
			element.Value=err.Param()
			errors=append(errors,element)
		}
	}

	return errors

}

func (co *controller)Logout(c *fiber.Ctx) error {

	    sess, err := co.session.Get(c)
  
		if err != nil {
        return err
    }

    // Destroy the session to delete it
    err = sess.Destroy()
    if err != nil {
        return err
    }
	sess.Save()

	 

	c.Cookie(&fiber.Cookie{
		 Name:    "sessionId",
        Value:   "",
        Expires: time.Now().Add(-time.Hour), // Set expiration in the past
	})



	return c.Status(200).JSON(fiber.Map{
		"status":200,
		"Message":"success",
	})
}


func (co *controller) GetUseridFromSessionMapping(sessionId string) int{

	for _,s :=range co.sessionMap {
		if s.SessionId == sessionId{
			return s.UserId
		}
	}
	return 0
}













