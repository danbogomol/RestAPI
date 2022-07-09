package main

import (
	"restapi/internal/server"

	log "github.com/sirupsen/logrus"
)

func main() {
	initLogger()
	
	config := server.NewConfig(":9090")
	s := server.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}

func initLogger() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "15:04:05"
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)
}
