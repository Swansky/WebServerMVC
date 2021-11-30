package app

import (
	"awesomeProject1/config"
	"awesomeProject1/mapping"
	"awesomeProject1/repositories"
	server2 "awesomeProject1/server"
)

func Start() {
	config.LoadSettings()
	settings := config.GetSettings()
	repositories.NewRepositoryManager()
	/*  repository := repositories.GetUserRepository()
	user, err := repository.Create(models.NewUser("swansky", "test"))
	if err != nil {
		panic(err)
		return
	}
	println(user.String())*/

	server := server2.NewServer(settings.Server.Port)
	mapping.MapURL(server)
	server.Start()
}
