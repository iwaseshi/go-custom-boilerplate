package main

import (
	"custom-boilerplate/pkg/restapi"

	_ "custom-boilerplate/pkg/di"
	_ "custom-boilerplate/services/healthcheck/adapter/controller"
	_ "custom-boilerplate/services/healthcheck/usecase"
)

func main() {
	restapi.Run()
}
