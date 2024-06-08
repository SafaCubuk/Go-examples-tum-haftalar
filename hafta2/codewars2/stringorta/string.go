package main

import (
	"fmt"
)

func ortaKarakter(s string) string {
	length := len(s)

	if length == 0 {
		return ""
	}

	orta := length / 2
	if length%2 == 0 {
		return s[orta-1 : orta+1]
	} else {

		return string(s[orta])
	}
}

func main() {
	fmt.Println(ortaKarakter("örnek"))
	fmt.Println(ortaKarakter("orta"))
	fmt.Println(ortaKarakter(""))
}
