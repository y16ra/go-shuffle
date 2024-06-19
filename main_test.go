package main

import "flag"

func Example() {
	flag.CommandLine.Set("items", "a,b,c,d,e")
	main()
	// Unordered output:
	// e
	// a
	// c
	// b
	// d
}
