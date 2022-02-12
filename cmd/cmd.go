package main

import (
	"github.com/dkharms/web/api"
	"io"
	"log"
	"os"
)

func main() {
	l := log.New(os.Stdout, "server: ", log.Ldate|log.Lshortfile)
	s := api.NewServer(l)

	s.AddHandler("/index", func(reader io.Reader, writer io.Writer) {

	})
	s.Run(80)
}
