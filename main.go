package main

import (
	"flag"
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
}
