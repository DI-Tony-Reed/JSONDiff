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
		log.Fatal(err)
	}

	log.Print(output)
}
