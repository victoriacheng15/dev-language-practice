package main

import "strings"

func Repeat(character string, n int) string {
	// this is running slower compare to the strings.Builder
	// run 'go test -bench=.' to see the benchmark results
	// var repeated string
	// for i := 0; i < n; i++ {
	// 	repeated += character
	// }
	// return repeated

	var repeated strings.Builder
	for i := 0; i < n; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
