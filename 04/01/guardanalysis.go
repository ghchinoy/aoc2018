package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var datafile string

func init() {

	flag.StringVar(&datafile, "file", "../data/sample.txt", "filepath to data file")
	flag.Parse()
}

type rec struct {
	GuardID int
	Time    time.Time
	Awake   bool
}

type Timeseries []rec

func main() {

	// guard, non-existent / unopenable data file
	if _, err := os.Open(datafile); err != nil {
		fmt.Println("can't open datafile", err)
		os.Exit(1)
	}
	// create timeseries from datafile
	t, err := ingest(datafile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Entries", len(t))

}

func ingest(datafilepath string) (Timeseries, error) {
	var t Timeseries
	databytes, err := ioutil.ReadFile(datafilepath)
	if err != nil {
		return t, err
	}
	// jan 2, 2003 4:05:6
	format := "2006-01-02 15:04"

	var guardID int
	for _, v := range strings.Split(fmt.Sprintf("%s", databytes), "\n") {
		if v == "" {
			continue
		}
		parts := strings.Split(v, "]")
		//log.Println(strings.Split(v, "]"))
		stamp := parts[0][1:]
		eventTime, err := time.Parse(format, stamp)
		if err != nil {
			log.Println(err)
			continue
		}
		awake := false

		action := parts[1]
		if strings.Contains(action, "#") {
			//log.Println(strings.Split(action, "#"))
			guardIDs := strings.Split(strings.Split(action, "#")[1], " ")[0]
			guardID, err = strconv.Atoi(guardIDs)
			if err != nil {
				log.Println(err)
				continue
			}
		} else if strings.Contains("awake", action) {
			awake = true
		}
		r := rec{guardID, eventTime, awake}
		//log.Printf("%s -> %s %v %v", stamp, eventTime, guardID, awake)
		log.Printf("%+v", r)
		t = append(t, r)
	}

	//time.Parse(format)
	return t, nil
}
