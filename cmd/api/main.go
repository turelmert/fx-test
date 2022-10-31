package main

import (
	"fx-test/di/api"
	"go.uber.org/fx"
)

func main() {
	fx.New(api.API()...).Run()
}
