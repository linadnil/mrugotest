package main

import (
	"fmt"
	"os"
	"bufio"
)

func countLines(f *os.File, counter int) int{
	input := bufio.NewScanner(f)
	var tmp int
	for input.Scan(){
		tmp++
	}
	return tmp
}

func main() {
	var counter int
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprint(os.Stderr, "mywc: %v\n", err)
	}
	tmp := countLines(f, counter)
	fmt.Printf("%d\n", tmp)
	f.Close()
}
