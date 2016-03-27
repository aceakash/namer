package main

import (
	"fmt"
	"strings"
)
import "flag"

func main() {
	n := flag.Int("n", 3, "Number of words")
	s := flag.String("s", " ", "Separator")
	flag.Parse()
	randomAdjectives := getRandomAdjectives(*n - 1)
	randomNoun := getRandomNoun()
	fmt.Printf("%s%s%s\n", strings.Join(randomAdjectives, *s), *s, randomNoun)
}
