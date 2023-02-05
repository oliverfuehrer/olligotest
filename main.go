package main

import (
	"fmt"
	"log"

	"example.com/olligotest/ollipackage"
	"rsc.io/quote"
)

func main() {
	fmt.Println(ollipackage.ZeigeBeispielA())
	fmt.Println(ollipackage.ZeigeBeispielB())

	fmt.Println(quote.Opt())

	log.Println("Start HTTP-Server ...")
	ollipackage.StartHttp()
}
