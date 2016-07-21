package main

import (
	"log"

	"github.com/sevlyar/go-daemon"
	"github.com/ryuen1111/go-daemon/app"
)

func main() {
	cntxt := &daemon.Context{
		PidFileName: "pid",
		PidFilePerm: 0644,
		LogFileName: "log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{"[go-daemon]"},
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Println(err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()

	app.MessageLoop()
}
