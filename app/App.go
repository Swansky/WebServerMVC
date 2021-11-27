package app

import (
	server2 "awesomeProject1/server"
)

func Start() {
	server := server2.NewServer(1293)
	MapURL(server)
	server.Start()
}
