package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var inputfile string

func init() {
	flag.StringVar(&inputfile, "file", "../data/input.txt", "file with list of candidate boxes")
	flag.Parse()
}

func main() {
	if _, err := os.Open(inputfile); err != nil {
		fmt.Println("can't open inputfile", err)
		os.Exit(1)
	}

	filebytes, err := ioutil.ReadFile(inputfile)
	if err != nil {
		fmt.Println("can't read file", err)
		os.Exit(1)
	}
	boxIDs := strings.Split(fmt.Sprintf("%s", filebytes), "\n")

	for k, v := range boxIDs {
		fmt.Println(k, v)
	}
}
