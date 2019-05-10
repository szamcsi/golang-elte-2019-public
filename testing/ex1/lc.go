// Binary lc counts the number of lines in a file.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func lineCount(path string) int {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
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
		fmt.Printf("%d\t%s\n", lineCount(path), path)
	}
}
