package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	// Version 1
	group := r.Group("/v1")

	// New instance of transporter
	transporter := NewTransporter()

	v1 := group.Group("/user")
	{
		v1.GET("/:user_id", transporter.GetUserByID())
		v1.POST("", transporter.Create())
	}
}
