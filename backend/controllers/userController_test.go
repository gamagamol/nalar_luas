package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gamagamol/todos/backend/entity"
	"github.com/gamagamol/todos/backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/stretchr/testify/assert"
)

func Test_controller_Login(t *testing.T) {

	app := fiber.New()

	store:=session.New(session.Config{
		CookieHTTPOnly: true,
		Expiration: time.Hour*5,
	})

	

	model:=models.NewModel()
	controller:=NewController(model,store)
	app.Post("/login",controller.Login)
	
	tests := []struct {
		name    		string
		method 			string
		uri 			string
		expectedCode 	int
		request 		entity.User
		response 		entity.ResponseSuccess
	}{
		{
			name: "success",
			method: http.MethodPost,
			uri:"/login",
			request: entity.User{
				Username: "gamagamol",
				Password: "gamagamol",
			},
			response: entity.ResponseSuccess{
				Status: 201,
				Message: "Success",	
			},
			expectedCode: 201,
		},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reqBody, err := json.Marshal(test.request)
			assert.Nil(t, err, "Failed to marshal request body")

			req := httptest.NewRequest(test.method, test.uri, bytes.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)

			assert.Nil(t, err, "Should not return an error")
			assert.Equal(t, test.expectedCode, resp.StatusCode, "Status code should be as expected")

			defer resp.Body.Close()
			respBody, err := ioutil.ReadAll(resp.Body)
			assert.Nil(t, err, "Failed to read response body")

			assert.Equal(t, test.response, respBody, "Response body should be as expected")
		})
	}
}
