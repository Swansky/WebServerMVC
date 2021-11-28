package app

import (
	"awesomeProject1/controller"
	"awesomeProject1/server"
)

func MapURL(server *server.Server) {

	server.AddAuthRoute("/", controller.Home)
	server.AddRoute("/login", controller.Login)
	server.AddRoute("/logout", controller.Logout)

	server.AddRoute("/accessRefused", controller.AccessRefused)
}
