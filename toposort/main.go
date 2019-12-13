package main

import (
	"fmt"
	"log"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
	"linear algebra":        {"calculus"}, // 循環を発生させる為に入れた
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string][]string)

	visitAll = func(items map[string][]string) {
		for key, value := range items {
			if !seen[key] {
				seen[key] = true

				for _, v := range value {
					for _, w := range m[v] {
						if w == key {
							// 循環を報告し終了する
							log.Fatalf("error: circulation [%s: %v]\n", v, m[v])
						}
					}

					visitAll(map[string][]string{v: m[v]})
				}

				order = append(order, key)
			}
		}
	}

	visitAll(m)
	return order
}
