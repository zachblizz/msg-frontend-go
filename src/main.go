package main

import (
	"fmt"

	"go.uber.org/dig"

	svr "bitbucket/zblizz/jwt-go/pkg/server"
)

func buildContainer() *dig.Container {
	container := dig.New()

	container.Provide(svr.NewServer)

	return container
}

func main() {
	container := buildContainer()
	utils.LoadProps()

	// TODO: might want to try to use wire for this
	// REF: https://github.com/google/wire
	err := container.Invoke(func(server *svr.Server) {
		server.Run()
	})

	if err != nil {
		panic(err)
	}
}
