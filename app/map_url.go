package app

import (
	"awesomeProject1/controller"
	"awesomeProject1/server"
)

func MapURL(server *server.Server) {
	server.AddRoute("/home",server.BasicAuth(controller.Home) )
}
