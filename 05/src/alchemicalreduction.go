package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var datafile string
var debug bool
var shortest bool

const alpha = "abcdefghijklmnopqrstuvwxyz"

func init() {
	flag.StringVar(&datafile, "file", "", "input file")
	flag.BoolVar(&debug, "debug", false, "show debug log")
	flag.BoolVar(&shortest, "shortest", false, "05.02 find the unit whose removal produces the shortest reacted polymer")
	flag.Parse()
}

func main() {
	var data string

	// allow for a sequence on stdin or a file
	// regardless, should process as stream instead of
	// this reading all at once
	if datafile == "" {
		data = flag.Args()[0]
	} else {
		databytes, err := ioutil.ReadFile(datafile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		data = fmt.Sprintf("%s", databytes)
	}

	if shortest != true { // 05.01
		finalarray := react(strings.Split(data, ""))
		fmt.Printf("final %v units\n", len(finalarray))
		os.Exit(0)
	}

	var shortestLetter string
	shortestUnitLength := len(data)
	for _, letter := range strings.Split(alpha, "") {
		testdata := data
		testdata = strings.Replace(testdata, letter, "", -1)
		testdata = strings.Replace(testdata, strings.ToUpper(letter), "", -1)
		units := len(react(strings.Split(testdata, "")))
		if units < shortestUnitLength {
			shortestUnitLength = units
			shortestLetter = letter
		}
	}
	fmt.Printf("Shortest reacted polymer length %v produced when removing '%s'\n", shortestUnitLength, shortestLetter)

}

// questions
// should these be chars rather than strings?
// should unicode codepoints be considered?
// input and output slices, up to 2x memory
// how could this be done within the same slice?

// react iterates through unit reactions
func react(input []string) []string {
	var finalarray []string
	prevUnit := input[:1][0]
	finalarray = append(finalarray, prevUnit)
	if debug {
		log.Printf("%v units, starting with prevunit '%s'", len(input), prevUnit)
	}
	for idx := 1; idx < len(input); idx++ {
		unitUnderTest := input[idx]
		var n string
		if idx+1 < len(input) {
			n = input[idx+1]
		}
		if debug {
			log.Printf("* (%s %s) > %s", prevUnit, unitUnderTest, n)
		}
		if strings.ToLower(unitUnderTest) == strings.ToLower(prevUnit) { // similar, curious
			if unitUnderTest == prevUnit { // unit polarities match, bail
				if debug {
					log.Printf("check: %s %s", unitUnderTest, prevUnit)
				}
				finalarray = append(finalarray, unitUnderTest)
				continue
			}
			finalarray = reduce(append(finalarray, unitUnderTest)) //largearray[:idx+1])
			if len(finalarray) > 0 {
				prevUnit = finalarray[len(finalarray)-1]
			} else {
				prevUnit = ""
			}
		} else {
			finalarray = append(finalarray, unitUnderTest)
			prevUnit = unitUnderTest
		}
		if debug {
			log.Printf("%s %v", prevUnit, len(finalarray))
		}
		//log.Printf("%s", finalarray)
	}
	return finalarray
}

// reduce lops off last two units that have reacted
func reduce(input []string) []string {
	//log.Printf("> %s", input)
	//log.Printf("< %s", input[:len(input)-2])
	reduced := input[:len(input)-2]
	if debug {
		log.Printf("- %s %v", input[len(input)-2:len(input)], len(reduced))
	}
	return reduced
}
