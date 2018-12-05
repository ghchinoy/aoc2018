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

	// initialize a grid of appropriate size with 0's
	grid := make([][]int, ymax+1)
	for i := range grid {
		grid[i] = make([]int, xmax+1)
	}

	gridclaimants := make(map[string][]int)

	// in the grid, keep a counter of how many claims overlap
	for _, c := range claims {
		//fmt.Printf("Grid %v - %+v\n", k, c)
		for y := c.Origin[1]; y < c.Origin[1]+c.Size[1]; y++ {
			for x := c.Origin[0]; x < c.Origin[0]+c.Size[0]; x++ {
				// how many claims are in this spot
				gridval := grid[y][x]
				grid[y][x] = gridval + 1

				// save the ID in the list of coords
				coord := fmt.Sprintf("%vx%v", x, y)
				claimants := gridclaimants[coord]
				claimants = append(claimants, c.ID)
				gridclaimants[coord] = claimants
			}
		}
		//visualizeGrid(xmax, ymax, grid)
		//fmt.Println()
	}

	if vis {
		visualizeGrid(xmax, ymax, grid)
	}

	// output the answer for 03.01
	fmt.Printf("%v: %s\n", countInches(xmax, ymax, grid), "with >= 2 overlapping inches")

	// output the answer for 03.02
	// list all the coordinates found for each of the narrowed list of claims
	for id, grids := range gridsWithCount(xmax, ymax, grid, gridclaimants) {
		//fmt.Println(id, grids)

		// find the claim by ID
		var c Claim
		for _, v := range claims {
			if v.ID == id {
				c = v
				break
			}
		}
		// if the length of the coords equals the length of the coords
		// of the Claim, this claim has no other claims on top of it
		//fmt.Printf("  %s\n", listAllCoords(c))
		if len(grids) == len(listAllCoords(c)) {
			fmt.Printf("Unoverlapping ID: %v\n", id)
		}
	}
}

func listAllCoords(c Claim) []string {
	var coords []string
	for cy := c.Origin[1]; cy < c.Origin[1]+c.Size[1]; cy++ {
		for cx := c.Origin[0]; cx < c.Origin[0]+c.Size[0]; cx++ {
			coords = append(coords, fmt.Sprintf("%vx%v", cx, cy))
		}
	}
	return coords
}

// isPointWithinClaim returns whether a particular point is within a claim
func isPointWithinClaim(x, y int, c Claim) bool {
	for cy := c.Origin[1]; cy < c.Origin[1]+c.Size[1]; cy++ {
		for cx := c.Origin[0]; cx < c.Origin[0]+c.Size[0]; cx++ {
			if x == cx {
				if y == cy {
					return true
				}
			}
		}
	}
	return false
}

// gridsWithCount is an attempt to narrow the list of coords to check
func gridsWithCount(xmax, ymax int, grid [][]int, gridclaimants map[string][]int) map[int][]string {
	claimGrids := make(map[int][]string)
	// go through each point and ...
	for y := 0; y < ymax; y++ {
		for x := 0; x < xmax; x++ {
			coord := fmt.Sprintf("%vx%v", x, y)
			//log.Printf("%s %+v", coord, gridclaimants[coord])
			// ... determine whether that point has a single claimaint ...
			if len(gridclaimants[coord]) == 1 {
				coords := claimGrids[gridclaimants[coord][0]]
				coords = append(coords, coord)
				// ... and gather all the points, by claim ID
				claimGrids[gridclaimants[coord][0]] = coords
			}
		}
	}
	return claimGrids
}

// countInches returns the number of coordinates with 2 or more claims
// This yields the answer for 03.01
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
	for y := 0; y < ymax+1; y++ {
		for x := 0; x < xmax+1; x++ {
			fmt.Print(grid[y][x])
		}
		fmt.Print("\n")
	}
}

func visualizeBlankGrid(xmax, ymax int) {
	for y := 0; y < ymax+1; y++ {
		for x := 0; x < xmax+1; x++ {
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
