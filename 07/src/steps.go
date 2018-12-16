package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

var datafile string

func init() {
	flag.StringVar(&datafile, "file", "../data/sample.txt", "data file path")
	flag.Parse()
}

func main() {
	// guard against non-openable file
	if _, err := os.Open(datafile); err != nil {
		fmt.Println("Can't open data file", err)
		os.Exit(1)
	}
	// reading everything in then splitting probably
	// isn't efficient, should do a file stream read
	// but oh well
	databytes, err := ioutil.ReadFile(datafile)
	if err != nil {
		fmt.Println("Couldn't read datafile", err)
		os.Exit(1)
	}
	data := fmt.Sprintf("%s", databytes)

	// steps that have been completed
	var completedsteps []string
	// steps and their prerequisites
	steps := make(map[string][]string)
	var incompletesteps []string

	// assemble all step prerequsites
	for _, instruction := range strings.Split(data, "\n") {
		if instruction == "" { // guard against blank lines
			continue
		}
		var prereqStep, postreqStep string
		// conversion here from string to reader to use
		// fmt.Fscanf is idiomatic (rather than string parsing)
		// and should be coupled with a file stream reader
		r := strings.NewReader(instruction)
		fmt.Fscanf(r, "Step %s must be finished before step %s can begin.", &prereqStep, &postreqStep)
		steps[postreqStep] = append(steps[postreqStep], prereqStep)
		if len(steps[prereqStep]) == 0 {
			steps[prereqStep] = []string{}
		}
		// C before A
		//log.Printf("%s before %s\n", prereqStep, postreqStep)
	}

	incompletesteps = stepkeys(steps)

	/* 	for _, i := range incompletesteps {
	   		log.Printf("%s %+v", i, steps[i])
	   	}
	*/

	order := reorder(incompletesteps, steps)

	for len(order) != 0 {
		for k, s := range order {
			steptotal := len(steps[s])
			stepcomplete := completedcount(steps[s], completedsteps)
			//log.Printf("%v %s has %v/%v complete, %s", k, s, stepcomplete, steptotal, steps[s])
			if steptotal-stepcomplete == 0 {
				//log.Printf("%s %v is %v", s, len(steps[s]), steptotal-stepcomplete)
				completedsteps = append(completedsteps, s)
				//order = removeFromSlice(order, s)
				copy(order[k:], order[k+1:])
				order[len(order)-1] = ""
				order = order[:len(order)-1]
				steps = removeFromSteps(s, steps)
				order = reorder(order, steps)
				break
			} else {
				order = reorder(order, steps)
			}
			//log.Printf("✓ %s", completedsteps)
			//log.Printf("x %s", order)

		}

		/* log.Printf("order len %v", len(order))
		log.Printf("✓ %s", completedsteps)
		log.Printf("x %s", order) */
	}

	fmt.Println(strings.Join(completedsteps, ""))
}

func removeFromSteps(x string, steps map[string][]string) map[string][]string {
	for s, todo := range steps {
		steps[s] = removeFromSlice(todo, x)
	}
	return steps
}

func reorder(order []string, steps map[string][]string) []string {
	//log.Println(">order:", order)
	// stepcount, list of steps
	stepsbycount := make(map[int][]string)
	for _, t := range order {
		list := append(stepsbycount[len(steps[t])], t)
		sort.Strings(list)
		stepsbycount[len(steps[t])] = list //append(stepsbycount[len(steps[t])], t)
	}
	var keys []int
	for k := range stepsbycount {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var neworder []string
	for _, v := range keys {
		//log.Printf("%v %s", v, sort.StringSlice(stepsbycount[v]))
		neworder = append(neworder, sort.StringSlice(stepsbycount[v])...)
	}
	//log.Println("<order:", neworder)
	return neworder
}

func completedcount(in []string, c []string) int {
	count := 0
	for _, v := range in {
		for _, i := range c {
			if i == v {
				count++
			}
		}
	}
	return count
}

func removeFromSlice(origin []string, remove string) []string {
	var adjusted []string
	for _, v := range origin {
		if v == remove {
			continue
		}
		adjusted = append(adjusted, v)
	}
	return adjusted
}

// unique returns a slice with only unique items
func unique(in []string) []string {
	if len(in) == 0 {
		return in
	}
	sort.Strings(in)
	j := 0
	for i := 1; i < len(in); i++ {
		if in[j] == in[i] {
			continue
		}
		j++
		in[i], in[j] = in[j], in[i]
	}
	result := in[:j+1]
	return result
}

func in(s []string, a string) bool {
	at := -1
	for k, v := range s {
		if v == a {
			at = k
		}
	}
	if at < 0 {
		return false
	}
	return true
}

func stepkeys(steps map[string][]string) []string {
	var keys []string
	for k := range steps {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

//////

func _sliceContains(s []string, i string) bool {
	sort.Strings(s)
	log.Println("sliceContains", sort.SearchStrings(s, i))
	if sort.SearchStrings(s, i) != 0 {
		return true
	}
	return false
}

// isStepIn checks to see if key is in list
func _isStepIn(step string, list []string) bool {
	for _, v := range list {
		if step == v {
			return true
		}
	}
	return false
}

// addDependency appends only if candidate doesn't already exist
// also, sorts the slice
func _addDependency(dependents []string, candidate string) []string {
	if in(dependents, candidate) {
		return dependents
	}
	dependents = append(dependents, candidate)
	//sort.Strings(dependents)
	return unique(dependents)
}

// remove removes slice x contents from slice in
func _remove(in []string, x []string) []string {
	var r []int
	for k, v := range in {
		for _, d := range x {
			if d == v {
				in = append(in[:k], in[k+1:]...)
			}
		}
	}
	log.Printf("%s - %s = %v", in, x, r)
	return in
}
