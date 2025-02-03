package main

import (
	"fmt"
	"log"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	// a loop
	"linear algebra":        {"calculus"},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	order, cycle := topoSortCycles(prereqs)
	if !cycle {
		log.Fatalln("Cyclic prerequisite")
	}
	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSortCycles(m map[string][]string) ([]string, bool) {
	var order []string
	seen := make(map[string]bool)

	stack := make(map[string]bool)
	var visitAll func(items []string) bool
	visitAll = func(items []string) bool {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true

				stack[item] = true
				if !visitAll(m[item]) {
					return false
				}
				stack[item] = false
				order = append(order, item)
			} else if stack[item] {
				return false
			}
		}
		return true
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	if found := visitAll(keys); found {
		return order, true
	}
	return nil, false
}
