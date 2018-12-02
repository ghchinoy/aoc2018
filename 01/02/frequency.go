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
	fmonitor := make(map[int]int)

	found := false
	var rot int

	for found != true {
		for _, change := range fs {
			if change[:1] == "#" { // skip lines with "#"
				continue
			}
			// fmt.Println(k, f, change, fmonitor[f])
			fmonitor[f] = fmonitor[f] + 1
			if fmonitor[f] == 2 {
				log.Println("second instance of frequency", f)
				found = true
				break
			}
			/*
				if val, ok := fmonitor[f]; ok {
					fmonitor[f] = val + 1
					if fmonitor[f] == 2 {
						log.Println("second instance of frequency", fmonitor[f])
					}
				} else {
					fmonitor[f] = fmonitor[f] + 1
				}
			*/
			f, e = action(f, change)
			if e != nil {
				log.Println(e.Error())
			}
		}
		rot++
	}

	fmt.Println("Frequency", f)
	fmt.Println("in this many rotations", rot)

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
