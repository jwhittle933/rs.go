package main

import (
	"log"
	"os"

	"github.com/jwhittle933/rs.go/option"
)

func main() {
	if file := openFile("result.txt"); file.IsNone() {
		log.Fatalln("Could not find file")
	}
}

func openFile(path string) option.Option[*os.File] {
	f, err := os.Open(path)
	if err != nil {
		return option.None[*os.File]()
	}

	return option.Some(f)
}
