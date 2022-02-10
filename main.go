package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	flag.StringVar(&name, "n", "", "")
	flag.Parse()
	fmt.Println(name)
}
