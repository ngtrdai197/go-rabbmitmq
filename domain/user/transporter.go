package user

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/ngtrdai197/go-rabbitmq/pkg/helper"
	"github.com/ngtrdai197/go-rabbitmq/pkg/response"
)

type Transporter interface {
	GetUserByID() gin.HandlerFunc
	Create() gin.HandlerFunc
}

type transporter struct {
}

func NewTransporter() Transporter {
	return &transporter{}
}

func (t *transporter) GetUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("user_id")

		if !helper.ValidateUUID(userID) {
			response.SendErrorResponse(c, "invalid user_id", 400, 400)
			return
		}

		response.SendSuccessResponse(c, userID)
	}
}

func (t *transporter) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var message CreateUserRequest
		if err := c.ShouldBindJSON(&message); err != nil {
			response.SendErrorResponse(c, err.Error(), 400, 400)
			return
		}
		fmt.Printf("message: %v\n", message)
		response.SendSuccessResponse(c, "ok")
	}
}
