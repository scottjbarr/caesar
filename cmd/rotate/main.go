package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/scottjbarr/caesar"
)

func main() {
	offset := flag.Int("offset", 13, "Character offset e.g. 13 for rot13 :)")
	flag.Parse()

	// read stdin
	in, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}

	// trim the string input
	data := strings.Trim(string(in), "\n")

	// transpose the text against a caesar cipher
	out := caesar.Transpose([]byte(data), *offset)

	fmt.Printf("%v\n", string(out))
}
