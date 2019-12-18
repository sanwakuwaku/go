package main

import (
	"flag"
	"fmt"

	"sanwakuwaku/go/exercise7/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")
var termo = tempconv.KelvinFlag("termo", 293.15, "the thermodynamic temperture")

func main() {
	flag.Parse()
	fmt.Println(*temp)
	fmt.Println(*termo)
}
