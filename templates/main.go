package templates

const (
	// Main is main.go templates.
	Main = `package main

import (
	"github.com/pthethanh/micro/server"
	"{{.module_name}}/internal/app/user"
)

func main() {
	srv := user.New()
	if err := server.ListenAndServe(srv); err != nil {
		panic(err)
	}
}
	
`
)
