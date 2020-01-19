package main

import (
	"go.uber.org/dig"

	svr "github.com/msg-frontend-go/pkg/server"
)

func buildContainer() *dig.Container {
	container := dig.New()

	container.Provide(svr.NewServer)

	return container
}

func main() {
	container := buildContainer()

	// TODO: might want to try to use wire for this
	// REF: https://github.com/google/wire
	err := container.Invoke(func(server *svr.Server) {
		server.Run()
	})

	if err != nil {
		panic(err)
	}
}
