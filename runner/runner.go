package main

import (
	// "github.com/pcrawfor/fayego/fayeserver"
	"../fayeserver"
)

func main() {
	fayeserver.Start(":3000")
}
