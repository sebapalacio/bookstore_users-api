package app

import (
	"github.com/sebapalacio/bookstore_users-api/controllers/ping"
	"github.com/sebapalacio/bookstore_users-api/controllers/user"
)

func mapUrls() {
	//Toda url de tipo get al server /ping, va a ser handleada por el controller -> ping_controller -> Ping() -> Function to need be executed.
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", user.GetUser)
	router.GET("/users/search", user.SearchUser)
	router.POST("/users", user.CreateUser)
}
