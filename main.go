package main

import (
	"log"
	"os"
)

func main() {
	runner := Runner{
		Arguments: os.Args[1:],
	}

	output, err := runner.Run(OSFileReader{})
	if err != nil {
		log.Fatal(output, err)
	}

	log.Print(output)
}
