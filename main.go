package main

import "github.com/ahsen17/BlogServ/src/server"

func main() {
	serverMgr := server.ServMgr{}

	serverMgr.RunServer()
}
