package main

import (
	"errors"
	"fmt"
	"github.com/bmwadforth/drift/src"
	"log"
	"os"
)

func main() {
	src.SetWD()
	src.SetMigrationPath()

	args := os.Args[1:]

	if len(args) > 0 {
		fmt.Println(args)
		switch args[0] {
		case "init": {
			_, err := src.Initialise(); if err != nil {
				log.Fatal(err)
			}
		}
		case "add": {
			_, err := src.AddMigration(args[1]); if err != nil {
				log.Fatal(err)
			}
		}
		case "remove": {

		}
		default:
			log.Fatal(errors.New("invalid argument supplied"))
		}
	} else {
		//TODO: Binary called with new arguments, process.exit
		//Print CLI HELP
		log.Fatal(errors.New("an argument must be supplied"))
	}
}
