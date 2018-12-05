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
	Entry   string
}

// ByTime is a sorting type for records
type ByTime []rec

func (a ByTime) Len() int           { return len(a) }
func (a ByTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTime) Less(i, j int) bool { return (a[i].Time).Before(a[j].Time) }

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

// ingest does ugly parsing
func ingest(datafilepath string) ([]rec, error) {
	var records []rec

	databytes, err := ioutil.ReadFile(datafilepath)
	if err != nil {
		return records, err
	}
	// jan 2, 2003 4:05:6
	format := "2006-01-02 15:04"

	for _, v := range strings.Split(fmt.Sprintf("%s", databytes), "\n") {
		if v == "" { // guard against blank entries
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
		r := rec{Time: eventTime, Entry: parts[1]}
		records = append(records, r)
	}
	sort.Sort(ByTime(records))

	var guardID int
	for k, r := range records {
		awake := false
		action := r.Entry
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
		r.GuardID = guardID
		r.Awake = awake
		// rec{GuardID: guardID, Time: eventTime, Awake: awake}
		//log.Printf("%s -> %s %v %v", stamp, eventTime, guardID, awake)
		log.Printf("%+v", r)
		records[k] = r
		//records = append(records, r)
	}

	//time.Parse(format)
	return records, nil
}
