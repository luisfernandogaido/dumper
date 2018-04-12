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

//http://localhost:4001/get?token=MMwwQvpYsfg9Xjse4VDqvjfm&dump=atlas-20180412-142824.sql

