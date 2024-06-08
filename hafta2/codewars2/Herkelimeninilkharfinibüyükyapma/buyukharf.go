package main

import (
	"fmt"
	"strings"
	"unicode"
)

func CapitalizeWords(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		runes := []rune(word)
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
		}
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}

func main() {
	text := "her kelimenin ilk harfini büyük yapma"
	result := CapitalizeWords(text)
	fmt.Println(result)
}
