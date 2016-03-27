# name-it

Generate a random, colourful name like "invincible miscreant cheese" or "whimsical habitual gnarly cat"

Only usable as a command-line tool for now. Programmatic access coming soon...

## Pre-requisites

Go 1.3+ installed, with $GOPATH set

## Build
```
go get github.com/aceakash/name-it
cd $GOPATH/src/github.com/aceakash/name-it
./build.sh
```

## Usage

```
name-it             # synonymous filthy sheep
name-it -n 4        # black decorous flying thought
name-it -n 7        # spiritual lackadaisical sweltering animated panicky illustrious plastic
name-it -n 2 -s _   # tremendous_cemetery
name-it -n 4 -s --  # berserk--repulsive--stingy--trucks
```

If you haven't guessed it by now, `name-it` takes two optional arguments:

* *-n [number]* : how many words to generate (default 3)
* *-s [separator]* : what separator to use between words (default " ")