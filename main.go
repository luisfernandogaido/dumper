package main

import (
	"dumper/app"
	"dumper/conf"
	"log"
)

func main() {
	if err := conf.Load("./conf.json"); err != nil {
		log.Fatal(err)
	}
	app.Ini(":4001")
}
