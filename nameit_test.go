package namer

import (
	"strings"
	"testing"
)

func TestShouldNotReturnNil(t *testing.T) {
	name := NameIt(3, " ")
	if name == nil {
		t.Error("NameIt() returned nil")
	}
}

func TestShouldReturnSpecifiedNumberOfWords(t *testing.T) {
	name := NameIt(3, " ")
	if len(name.Words) != 3 {
		t.Errorf("Expected 3 words but got %d", len(name.Words))
	}
}

func TestShouldReturnOneNounAndRestAdjectives(t *testing.T) {
	name := NameIt(4, "-")
	if !doesSliceContainString(nouns, name.Words[3]) {
		t.Errorf("Expected %s to be a noun", name.Words[3])
	}
	for _, word := range name.Words[:3] {
		if !doesSliceContainString(adjectives, word) {
			t.Errorf("Expected %s to be an adjective", word)
		}
	}
}

func TestShouldHaveNameAsWordsWithSeparator(t *testing.T) {
	name := NameIt(5, "_")
	if name.Name != strings.Join(name.Words, "_") {
		t.Errorf("Expected %q to be %v joined with a \"_\"", name.Name, name.Words)
	}
}

func doesSliceContainString(haystack []string, needle string) bool {
	for _, word := range haystack {
		if word == needle {
			return true
		}
	}
	return false
}
