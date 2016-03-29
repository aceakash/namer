package main

import (
	"fmt"

	"github.com/aceakash/namer"
)
import "flag"

func main() {
	n := flag.Int("n", 3, "Number of words")
	s := flag.String("s", " ", "Separator")
	flag.Parse()

	name := namer.NameIt(*n, *s)
	fmt.Println(name.Name)
}
