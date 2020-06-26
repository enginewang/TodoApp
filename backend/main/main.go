package main

import (
	"StackBackend"
	"StackBackend/db"
	"log"
)

func main() {
	err := db.InitGlobalDB("127.0.0.1", "stack")
	if err != nil{
		log.Panic(err)
	}
	s := StackBackend.NewServer(":1323")
	err = s.Init()
	if err!=nil{
		log.Panic(err)
	}
	s.StartServer()
}