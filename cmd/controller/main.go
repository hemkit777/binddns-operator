package main

import "github.com/hemkit777/binddns-operator/cmd/controller/app"

func main() {
	c := app.NewCommand()
	if err := c.Execute(); err != nil {
		panic(err)
	}
}
