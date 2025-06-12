package main

import (
	"flag"
	"fmt"
	"io"
	"strings"
)

func main() {
	city := flag.String("city", "", "City of the user")
	format := flag.Int("format", 1, "Format of the weather")

	flag.Parse()

	fmt.Println(*city)
	fmt.Println(*format)

	r := strings.NewReader("Hello, World!")
	block := make([]byte, 4)

	for {
		_, err := r.Read(block)
		fmt.Printf("%q\n", block)
		if err == io.EOF {
			break
		}
	}

}
