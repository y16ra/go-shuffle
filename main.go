package main

import (
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type stringSlice []string

func (s *stringSlice) Set(value string) error {
	if len(*s) > 0 {
		return errors.New("interval flag already set")
	}
	for _, dt := range strings.Split(value, ",") {
		*s = append(*s, dt)
	}
	return nil
}

func (s *stringSlice) String() string {
	return fmt.Sprintf("%v", *s)
}

var items stringSlice

func main() {
	// Parse command line arguments
	flag.Parse()
	// Shuffle items
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	r.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})
	for _, item := range items {
		fmt.Println(item)
	}
}

func init() {
	flag.Var(&items, "items", "comma separated list of items")
}
