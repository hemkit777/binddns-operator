package main

import "github.com/hemkit777/binddns-operator/cmd/webhook/app"

func main() {
	c := app.NewCommand()
	if err := c.Execute(); err != nil {
		panic(err)
	}
}
