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

// rec is a record
type rec struct {
	GuardID int
	Time    time.Time
	Awake   bool
	Entry   string
}

// GuardMeta is metadata about a Guard
type GuardMeta struct {
	GuardID              int
	TotalSleepingMinutes int
	MinuteCount          map[int]int
}

func main() {

	// guard against non-existent / unopenable data file
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
	log.Println("Entries", len(t))

	vis(t)
	calculateAsleepyness(t)
}

func calculateAsleepyness(records []rec) {
	// who slept the most
	var sleepiestGuard int
	var sleepFrom, sleepTo, guard int
	asleep := make(map[int]int)
	for _, r := range records {
		if strings.Contains(r.Entry, "begins shift") {
			guard = r.GuardID
		} else if strings.Contains(r.Entry, "falls asleep") {
			sleepFrom = r.Time.Minute()
		} else if strings.Contains(r.Entry, "wakes") {
			sleepTo = r.Time.Minute()
			length := sleepTo - sleepFrom
			asleep[guard] += length
			if asleep[guard] > asleep[sleepiestGuard] {
				sleepiestGuard = guard
			}
		}
	}
	fmt.Printf("Guard #%v slept %vm\n", sleepiestGuard, asleep[sleepiestGuard])

	// minute where sleeping occured the most
	minutes := make(map[int]int)
	var mostSleepyMinute int
	for _, r := range records {
		if r.GuardID != sleepiestGuard { // only sleepy guard's records
			continue
		}
		if strings.Contains(r.Entry, "begins shift") { // only sleeps and wakes needed
			guard = r.GuardID
			continue
		}
		if strings.Contains(r.Entry, "falls asleep") {
			sleepFrom = r.Time.Minute()
		} else if strings.Contains(r.Entry, "wakes") {
			sleepTo = r.Time.Minute()
			for i := sleepFrom; i < sleepTo; i++ {
				minutes[i]++
				if minutes[i] > minutes[mostSleepyMinute] {
					mostSleepyMinute = i
				}
			}
		}
	}

	fmt.Printf("Sleepiest minute for Guard %v is %v\n", guard, mostSleepyMinute)
	fmt.Printf("Answer 04.01 (%v x %v): %v\n", guard, mostSleepyMinute, guard*mostSleepyMinute)
}

// ingest does ugly parsing
func ingest(datafilepath string) ([]rec, error) {
	var records []rec

	// read in file
	databytes, err := ioutil.ReadFile(datafilepath)
	if err != nil {
		return records, err
	}
	// jan 2, 2003 4:05:6
	format := "2006-01-02 15:04"

	// since records may be out of order, get the time
	for _, v := range strings.Split(fmt.Sprintf("%s", databytes), "\n") {
		if v == "" { // guard against blank entries
			continue
		}
		parts := strings.Split(v, "]")
		stamp := parts[0][1:]
		eventTime, err := time.Parse(format, stamp)
		if err != nil {
			log.Println(err)
			continue
		}
		// "Because all asleep/awake times are during the midnight hour (00:00 - 00:59),
		// only the minute portion (00 - 59) is relevant for those events."
		// Let's convert day with hour 23 to day+1 00:00
		if eventTime.Hour() != 0 {
			zero := 0
			eventTime = time.Date(eventTime.Year(), eventTime.Month(), eventTime.Day()+1, zero, zero, eventTime.Second(), eventTime.Nanosecond(), eventTime.Location())
			log.Printf("Forcing time change for \"%s\" to %s", v, eventTime.Format(format))
		}
		// construct the record
		r := rec{Time: eventTime, Entry: parts[1]}
		// append the record
		records = append(records, r)
	}

	// sort by time
	// "your entries are in the order you found them.
	// You'll need to organize them before they can be analyzed."
	sort.Slice(records, func(i, j int) bool { return records[i].Time.Before(records[j].Time) })

	// parse the entry, update the record
	var guardID int
	for k, r := range records {
		awake := false
		action := r.Entry
		// parse action
		if strings.Contains(action, "#") {
			guardIDs := strings.Split(strings.Split(action, "#")[1], " ")[0]
			guardID, err = strconv.Atoi(guardIDs)
			if err != nil {
				log.Println(err)
				continue
			}
			awake = true
		} else if strings.Contains(action, "wakes") {
			awake = true
		}
		r.GuardID = guardID
		r.Awake = awake
		records[k] = r
	}

	return records, nil
}

// vis shows the records
func vis(records []rec) {

	format := "%5s %5s %s\n"

	// header
	fmt.Printf(format, "Date", "ID", "Minute")
	// create stacked time output
	var f, s []string
	for i := 0; i < 6; i++ {
		for j := 0; j < 10; j++ {
			f = append(f, strconv.Itoa(i))
			s = append(s, strconv.Itoa(j))
		}
	}
	fmt.Printf(format, "   ", "   ", strings.Join(f, ""))
	fmt.Printf(format, "   ", "   ", strings.Join(s, ""))

	// Gather the records by day
	//var timeline string
	//byday := make(map[time.Time][]rec)
	byday := make(map[string][]rec)
	//byday := make(ByShiftTime)
	sort.Slice(records, func(i, j int) bool { return records[i].Time.Before(records[j].Time) })
	for _, v := range records {
		// just the day, please
		day := v.Time.Format("01-02")
		//day := time.Date(v.Time.Year(), v.Time.Month(), v.Time.Day(), 0, 0, 0, 0, v.Time.Location())
		byday[day] = append(byday[day], v)
	}

	// Process the shifts
	var days []string
	for d := range byday {
		days = append(days, d)
	}
	sort.Strings(days)

	mark := "."
	sleepMin := 0
	guardlist := make(map[int]GuardMeta)

	for _, d := range days {
		s := byday[d]
		var date string
		var guardID int
		timeline := makeAwakeLine()

		for _, v := range s {
			if date != v.Time.Format("01-02") {
				date = v.Time.Format("01-02")
			}
			// set up guardmetadata
			if _, ok := guardlist[v.GuardID]; !ok {
				guardlist[v.GuardID] = GuardMeta{GuardID: v.GuardID, MinuteCount: make(map[int]int)}
			}
			if guardID != v.GuardID {
				guardID = v.GuardID
			}
			g := guardlist[v.GuardID]
			minute := v.Time.Minute()

			if strings.Contains(v.Entry, "begins") { // reset to conscious
				mark = "." // ">"
			} else if strings.Contains(v.Entry, "asleep") { // sleep starts here
				mark = "#"
				sleepMin = minute
				minuteList := g.MinuteCount
				if minuteList == nil {
					minuteList = make(map[int]int)
					minuteList[minute] = 1
				} else {
					minuteList[minute] = minuteList[minute] + 1
				}
				g.MinuteCount = minuteList
				guardlist[v.GuardID] = g
			} else if strings.Contains(v.Entry, "wakes") {
				mark = "."

				minuteList := g.MinuteCount
				for x := sleepMin + 1; x < minute; x++ {
					minuteList[x] = minuteList[x] + 1
					//log.Println(x, minuteList[x])
					timeline[x] = "#"
				}
				g.MinuteCount = minuteList
				g.TotalSleepingMinutes = len(minuteList) + 1
				guardlist[guardID] = g
				sleepMin = 0
			}
			//log.Printf("%+v", guardAsleepMinutes[v.GuardID])

			timeline[minute] = mark
			//log.Println(v.Time.Format("04"), i)
		}

		fmt.Printf(format, date, fmt.Sprintf("#%v", guardID), strings.Join(timeline, ""))
		//fmt.Println("asleep for", asleepMin)

	}

}

func makeAwakeLine() []string {
	var timeline []string
	for i := 0; i < 60; i++ {
		timeline = append(timeline, ".")
	}
	return timeline
}
