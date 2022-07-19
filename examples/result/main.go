package main

import (
	"os"

	"github.com/jwhittle933/rs.go/result"
)

func main() {
	file := openFile("./result.txt")
	file.Expect("Could not open file")
}

func openFile(path string) result.Result[*os.File, error] {
	f, err := os.Open(path)
	if err != nil {
		return result.Err[*os.File](err)
	}

	return result.Ok(f)
}
