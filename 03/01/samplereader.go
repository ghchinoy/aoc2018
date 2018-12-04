package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Claim is an Elf's claim on the santacloths
type Claim struct {
	ID     int
	Origin []int
	Size   []int
}

// ByXOrigin is a sorting type for Claims
type ByXOrigin []Claim

func (a ByXOrigin) Len() int           { return len(a) }
func (a ByXOrigin) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByXOrigin) Less(i, j int) bool { return a[i].Origin[0] < a[j].Origin[0] }

// ByYOrigin is a sorting type for Claims
type ByYOrigin []Claim

func (a ByYOrigin) Len() int           { return len(a) }
func (a ByYOrigin) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByYOrigin) Less(i, j int) bool { return a[i].Origin[1] < a[j].Origin[1] }

var inputfile string
var vis bool

func init() {
	flag.StringVar(&inputfile, "file", "../data/sample.txt", "file with list of claim IDs")
	flag.BoolVar(&vis, "vis", false, "show a visualization")
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
	claimsdata := strings.Split(fmt.Sprintf("%s", filebytes), "\n")
	var claims []Claim
	var xmax, ymax int
	for k, v := range claimsdata {
		//fmt.Println(k, v)
		if strings.HasPrefix(v, "-") {
			continue
		}
		// convert to Claim
		c, err := newClaim(v)
		if err != nil {
			log.Println("Can't convert to claim", v)
		}
		claims = append(claims, c)
		log.Printf("%v %+v", k, v)

		xtot := c.Origin[0] + c.Size[0]
		ytot := c.Origin[1] + c.Size[1]
		if xtot > xmax {
			xmax = xtot
		}
		if ytot > ymax {
			ymax = ytot
		}
	}

	sort.Sort(ByYOrigin(claims))
	sort.Sort(ByXOrigin(claims))

	// visualize blank grid
	// visualizeBlankGrid(xmax, ymax)

	// initialize a grid with 0's
	grid := make([][]int, ymax)
	for i := range grid {
		grid[i] = make([]int, xmax)
	}

	for _, c := range claims {
		//fmt.Printf("Grid %v - %+v\n", k, c)
		for y := c.Origin[1]; y < c.Origin[1]+c.Size[1]; y++ {
			for x := c.Origin[0]; x < c.Origin[0]+c.Size[0]; x++ {
				// how many claims are in this spot
				gridval := grid[y][x]
				grid[y][x] = gridval + 1
			}
		}
		//visualizeGrid(xmax, ymax, grid)
		//fmt.Println()
	}

	if vis {
		visualizeGrid(xmax, ymax, grid)
	}
	fmt.Println(countInches(xmax, ymax, grid))
}

// countInches returns the number of coordinates with 2 or more claims
func countInches(xmax, ymax int, grid [][]int) int {
	var inches int
	for y := 0; y < ymax; y++ {
		for x := 0; x < xmax; x++ {
			if grid[y][x] >= 2 {
				inches++
			}
		}
	}
	return inches
}

// visualizeGrid outputs a text grid with the marks within the given grid
func visualizeGrid(xmax, ymax int, grid [][]int) {
	for y := 0; y < ymax; y++ {
		for x := 0; x < xmax; x++ {
			fmt.Print(grid[y][x])
		}
		fmt.Print("\n")
	}
}

func visualizeBlankGrid(xmax, ymax int) {
	for y := 0; y < ymax; y++ {
		for x := 0; x < xmax; x++ {
			fmt.Print(".")
		}
		fmt.Print("\n")
	}
}

// newClaim parses a string into a Claim
func newClaim(claimstr string) (Claim, error) {
	var c Claim
	// #1242 @ 746,650: 10x13
	parts := strings.Split(claimstr, " ")
	id, err := strconv.Atoi(parts[0][1:])
	if err != nil { // cant convert ID to int
		return c, err
	}

	origin, err := coordsplit(strings.TrimSuffix(parts[2], ":"), ",")
	if err != nil { // can't convert origin coords to int
		return c, err
	}

	patchcoords, err := coordsplit(parts[3], "x")
	if err != nil { // can't convert patch coords to int
		return c, err
	}

	//log.Printf("%+v", parts)

	return Claim{id, origin, patchcoords}, nil
}

func coordsplit(coordstr string, sep string) ([]int, error) {
	var coords []int
	ords := strings.Split(coordstr, sep)
	for _, v := range ords {
		ord, err := strconv.Atoi(v)
		if err != nil { // can't convert a coord to int
			return coords, err
		}
		coords = append(coords, ord)
	}
	return coords, nil
}
