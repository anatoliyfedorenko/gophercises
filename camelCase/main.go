package main

import (
	"fmt"
	"unicode"
)

func main() {
	in := "thisIsASimpleSentense"
	fmt.Printf("The sentence %s has %d words\n", in, detectWordsFromCamelCase(in))
}

func detectWordsFromCamelCase(in string) int {
	out := 1
	for _, r := range in {
		if unicode.IsUpper(r) && unicode.IsLetter(r) {
			out++
		}
	}
	return out
}
