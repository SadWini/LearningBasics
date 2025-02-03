package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}
type multitier struct {
	data    []*Track
	primary string
	second  string
	third   string
}

// ref: https://stackoverflow.com/questions/25824095/order-by-clicking-table-header
func main() {
	multi := NewMultitier(tracks, "Title", "", "")
	sort.Sort(multi)
	printTracks(os.Stdout, multi)
	// start a simple server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			fmt.Printf("ParseForm: %v\n", err)
		}
		for k, v := range r.Form {
			if k == "orderby" {
				ChangePriorities(multi, v[0])
			}
		}
		sort.Sort(multi)
		printTracks(w, multi)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func printTracks(w io.Writer, x sort.Interface) {
	if x, ok := x.(*multitier); ok {
		trackTable.Execute(w, x)
	}
}

func (x *multitier) Len() int {
	return len(x.data)
}

func (x *multitier) Swap(i, j int) {
	x.data[i], x.data[j] = x.data[j], x.data[i]
}

func (x *multitier) Less(i, j int) bool {
	key := x.primary
	for k := 0; k < 3; k++ {
		switch key {
		case "Title":
			if x.data[i].Title != x.data[j].Title {
				return x.data[i].Title < x.data[j].Title
			}
		case "Year":
			if x.data[i].Year != x.data[j].Year {
				return x.data[i].Year < x.data[j].Year
			}
		case "Artist":
			if x.data[i].Artist != x.data[j].Artist {
				return x.data[i].Artist < x.data[j].Artist
			}
		case "Length":
			if x.data[i].Length != x.data[j].Length {
				return x.data[i].Length < x.data[j].Length
			}
		}
		if k == 0 {
			key = x.second
		} else if k == 1 {
			key = x.third
		}
	}
	return false
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func changePriorities(x *multitier, p string) {
	x.primary, x.second, x.third = p, x.primary, x.second
}

func ChangePriorities(x sort.Interface, p string) {
	if x, ok := x.(*multitier); ok {
		changePriorities(x, p)
	}
}

func NewMultitier(data []*Track, prim, sec, th string) sort.Interface {
	return &multitier{
		data:    data,
		primary: prim,
		second:  sec,
		third:   th,
	}
}
