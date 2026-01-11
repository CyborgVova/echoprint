package main

import (
	"flag"
	"log"

	"github.com/cyborgvova/echoprint/app"
	"github.com/cyborgvova/echoprint/config"
)

var (
	text string
	port int
)

func init() {
	flag.StringVar(&text, "text", "Default message", "<text message>")
	flag.IntVar(&port, "port", 8888, "<port number>")
}

func main() {
	flag.Parse()

	cfg := &config.Config{
		Text: text,
		Port: port,
	}

	application := app.New(cfg)

	if err := application.Start(); err != nil {
		log.Fatal("start application:", err)
	}
}
