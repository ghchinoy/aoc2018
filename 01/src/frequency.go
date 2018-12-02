package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var freqfile string

func init() {
	flag.StringVar(&freqfile, "file", "../data/input.txt", "Frequency file path")
	flag.Parse()
}
func main() {
	// read in frequency file
	_, err := os.Open(freqfile)
	if err != nil {
		fmt.Println("Can't find frequency file")
		os.Exit(1)
	}
	freqb, err := ioutil.ReadFile(freqfile)
	fs := strings.Split(fmt.Sprintf("%s", freqb), "\n")
	var f int
	var e error
	for _, change := range fs {
		if change[:1] == "#" { // skip lines with "#"
			continue
		}
		f, e = action(f, change)
		if e != nil {
			log.Println(e.Error())
		}
	}
	fmt.Println("Frequency", f)

}

func action(state int, change string) (int, error) {
	op := change[:1]
	amt, err := strconv.Atoi(change[1:])
	if err != nil {
		return 0, err
	}
	if op == "+" {
		state = state + amt
	} else {
		state = state - amt
	}
	return state, nil
}
