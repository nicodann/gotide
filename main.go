package main

import (
	"log"

	"github.com/nicodann/gotide/functions"
)

func main() {
	result := functions.Countletters(("amphibian"))
	log.Print(result)
}
