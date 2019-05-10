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
	if err != nil { // HLnottested
		log.Fatal(err) // HLnottested
	} // HLnottested
	defer f.Close()
	s := bufio.NewScanner(f)
	lc := 0
	for s.Scan() {
		lc++
	}
	return lc
}
func main() { // HLnottested
	flag.Parse()                       // HLnottested
	for _, path := range flag.Args() { // HLnottested
		fmt.Printf("%d\t%s\n", lineCount(path), path) // HLnottested
	} // HLnottested
} // HLnottested
