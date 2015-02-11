package main

import (
	"fmt"
	"github.com/go-src/hello/stringutil"
)

func main() {
	fmt.Printf("Hello, world. 你好，世界")
	fmt.Printf("\n")
	fmt.Printf(stringutil.Reverse("Hello, world. 你好，世界"))
	fmt.Printf("\n")
}
