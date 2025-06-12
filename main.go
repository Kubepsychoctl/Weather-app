package main

import (
	"flag"
	"fmt"
)

func main() {
	city := flag.String("city", "", "City of the user")
	format := flag.Int("format", 1, "Format of the weather")

	flag.Parse()

	fmt.Println(*city)
	fmt.Println(*format)
}
