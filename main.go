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

//https://dumper.profinanc.com.br/dumpget?host=localhost&user=root&pass=Semaver13&db=atlas&delete=true&token=VjtMvfPVV99fxJ7mqMNGEH29