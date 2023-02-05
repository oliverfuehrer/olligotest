package main

import (
	"fmt"

	ollip "example.com/olligotest/ollipackage"
	"rsc.io/quote"
)

func main() {
	fmt.Println(ollip.ZeigeOlliText2())
	fmt.Println(ollip.ZeigeOlliText3())

	fmt.Println(quote.Opt())

	ollip.StartHttp()
}
