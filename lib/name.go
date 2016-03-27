package lib

import "strings"

// Name is a name generated by name-it in several common forms
type Name struct {
	Words []string
	Name  string
}

// NameIt returns a name
func NameIt(number int, separator string) *Name {
	name := new(Name)
	name.Words = append(getRandomAdjectives(number-1), getRandomNoun())
	name.Name = strings.Join(name.Words, separator)
	return name
}
