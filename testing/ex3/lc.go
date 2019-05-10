// Binary lc counts the number of lines in a file.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/szamcsi/golang-elte-2019-public/testing/ex3/lines"
	"log"
	"os"
)

// TODO: move lineCount() to the 'lines' package.
func lineCount(path string) int {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err) // DO NOT PANIC!
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	lc := 0
	for s.Scan() {
		lc++
	}
	return lc
}

func main() {
	flag.Parse()
	for _, path := range flag.Args() {
		lc, err := lines.Count(path)
		if err != nil {
			fmt.Printf("ERROR: %s", err)
			continue
		}
		fmt.Printf("%d\t%s\n", lc, path)
	}
}
