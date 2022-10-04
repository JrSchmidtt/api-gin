package main

import (
	"github.com/JrSchmidtt/api-gin/database"
	"github.com/JrSchmidtt/api-gin/server"
)

func main(){
	database.InitDB()
	server := server.NewServer()
	server.Run()
}