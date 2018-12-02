package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var inputfile string

func init() {
	flag.StringVar(&inputfile, "file", "../data/input.txt", "file of list of candidates")
	flag.Parse()
}
func main() {
	_, err := os.Open(inputfile)
	if err != nil {
		fmt.Println("can't find inputfile", err)
		os.Exit(1)
	}
	filebytes, err := ioutil.ReadFile(inputfile)
	candidates := strings.Split(fmt.Sprintf("%s", filebytes), "\n")

	var lettersAppearTwice, lettersAppearThrice int

	for _, v := range candidates {
		letters := strings.Split(v, "")
		counter := make(map[string]int)
		// count frequency of letters in string
		for _, letter := range letters {
			counter[letter] = counter[letter] + 1
		}
		// check for first appearance of letters counting twice
		for _, c := range counter {
			if c == 2 {
				lettersAppearTwice++
				break // first appearance
			}
		}
		// check for first appearance of letters counting three times
		for _, c := range counter {
			if c == 3 {
				lettersAppearThrice++
				break
			}
		}
		log.Printf("%s +%v", v, counter)
	}

	fmt.Println("checksum", lettersAppearThrice*lettersAppearTwice)
}
