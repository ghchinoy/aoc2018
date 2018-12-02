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

	// var: hold a list of boxIDs that are different by one letter
	var pals []string

	log.Print("splitting")
	for _, box := range boxIDs {
		letters := strings.Split(box, "")

		log.Printf("%+v", letters)

		// then go through each of them, again
		for _, test := range boxIDs {

			if withinSlice(test, pals) { // no need to reprocess
				break
			}

			testletters := strings.Split(test, "")
			// how many letters are the same, in the same order
			var similaritiescount int
			for i := 0; i < len(letters); i++ {
				if letters[i] == testletters[i] {
					similaritiescount++
				}
			}
			// off by one, keep both of these
			if (len(letters) - 1) == similaritiescount {
				if !withinSlice(box, pals) {
					pals = append(pals, box)
					pals = append(pals, test)
				}
			}
		}
	}
	// show the similar off-by-ones
	log.Print("pals")
	log.Printf("%+v\n", pals)

	var commonletters []string
	lets1 := strings.Split(pals[0], "")
	lets2 := strings.Split(pals[1], "")
	for i := 0; i < len(lets1); i++ {
		if lets1[i] == lets2[i] {
			commonletters = append(commonletters, lets1[i])
		}
	}

	//log.Printf("%+v", common)

	fmt.Println(strings.Join(commonletters, ""))
}

// withinSlice returns whether string test is within string slice s
func withinSlice(test string, s []string) bool {
	for _, i := range s {
		if test == i {
			return true
		}
	}
	return false
}
