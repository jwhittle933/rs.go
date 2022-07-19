package main

import (
	"io/ioutil"
	"os"

	"github.com/jwhittle933/rs.go/result"
)

func main() {
	// Due to the limitation of Go's generics, this cannot be done in a single chain.
	file := result.Match(os.Open("result.txt")).Expect("could not open file")
	result.Match(ioutil.ReadAll(file)).Expect("Could not read file")
}
