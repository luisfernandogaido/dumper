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

//http://localhost:4001/get?token=MMwwQvpYsfg9Xjse4VDqvjfm&file=ect-20180412-142632.sql

