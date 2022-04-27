package app

import (
	"github.com/sri-05/bookstore_users-api/controllers/ping"
	"github.com/sri-05/bookstore_users-api/controllers/users"
)

func MapUrl() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
