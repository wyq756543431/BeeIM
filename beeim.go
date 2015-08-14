package main

import (
	"github.com/go-xweb/log"
	"github.com/wyq756543431/BeeIM/s2c"
)

func main(){
	s2c_server := s2c.CreateServer()
	log.Println("Running on")
	s2c_server.Start(":1114")
}
