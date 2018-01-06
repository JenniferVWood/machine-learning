package main

import (
	"fmt"
	"github.com/buskersguidetotheuniverse.org/weather"
)

func main() {
	fmt.Printf("Hello, world.\n")

	weather.Fetch("foo")
}
