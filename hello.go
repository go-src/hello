package main

import (
	"fmt"
	"github.com/go-src/stringutil"
)

func main() {
	fmt.Printf("Hello, world.")
	fmt.Printf("\n")
	fmt.Printf(stringutil.Reverse("Hello, world."))
	fmt.Printf("\n")
}
