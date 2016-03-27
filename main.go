package main

import (
	"fmt"
	"strings"
)
import "flag"

func main() {
	n := flag.Int("n", 3, "Number of words")
	flag.Parse()
	randomAdjectives := getRandomAdjectives(*n - 1)
	randomNoun := getRandomNoun()
	fmt.Printf("%v %s\n", strings.Join(randomAdjectives, " "), randomNoun)
}
