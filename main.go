package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randSource)
	fmt.Printf("%s %s\n", adjectives[random.Int31n(int32(len(adjectives)))], nouns[random.Int31n(int32(len(nouns)))])
}
