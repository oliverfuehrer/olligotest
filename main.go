package main

import (
	"fmt"
	"log"
	"time"

	"example.com/olligotest/ollipackage"
	"rsc.io/quote"
)

func main() {

	fmt.Println(ollipackage.ZeigeBeispielA())
	fmt.Println(ollipackage.ZeigeBeispielB())

	fmt.Println(quote.Opt())

	log.Println("Start HTTP-Server ...")

	go runningAndRepeating()

	ollipackage.StartHttp()
}

func runningAndRepeating() {
	for {
		log.Println("Inline thread is running and repeating ...")
		time.Sleep(1 * time.Second)
	}
}
