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
	{{if .web}}
	if err := server.New(server.FromEnv(), server.Web("/", "web", "index.html")).ListenAndServe(srv); err != nil {
		panic(err)
	}
	{{else}}
	if err := server.ListenAndServe(srv); err != nil {
		panic(err)
	}
	{{end}}
}
	
`
)
