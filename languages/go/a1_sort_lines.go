package main

import (
	"fmt"
	"os"
)

func main() {
  fmt.Printf("P1 - Go\n")
  fmt.Printf("Reading file");
  var path = "problems/assets/p1_lines.txt"
  fmt.Printf(path)
  readLines(path)
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	return lines, nil
}
