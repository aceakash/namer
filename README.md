# namer

Generate a random, colourful name like "invincible miscreant cheese" or "whimsical habitual gnarly cat".

## Pre-requisites

Go 1.3+ installed, with $GOPATH set

## Install
```
go get github.com/aceakash/namer/...
```

## Command-Line Usage
```
nameit             # synonymous filthy sheep
nameit -n 4        # black decorous flying thought
nameit -n 7        # spiritual lackadaisical sweltering animated panicky illustrious plastic
nameit -n 2 -s _   # tremendous_cemetery
nameit -n 4 -s --  # berserk--repulsive--stingy--trucks
```

If you haven't guessed it by now, `nameit` takes two optional arguments:

* *-n [number]* : how many words to generate (default 3)
* *-s [separator]* : what separator to use between words (default " ")

## Programmatic Usage
```
package main

import "fmt"
import "github.com/aceakash/namer" // <--------- bring it in

func main() {
  name := namer.NameIt(7, "--") // <--------- use it
  fmt.Println(name.Name)
  fmt.Println(name.Words)
}
```

## Build
```
cd $GOPATH/src/github.com/aceakash/namer
./build.sh
```
