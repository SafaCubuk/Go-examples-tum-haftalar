package main

import (
	"fmt"
)

func main() {
	var koyun int

	for {
		fmt.Print("Bir sayı girin (negatif olmayan): ")
		fmt.Scanln(&koyun)

		if koyun >= 0 {
			break
		}
	}

	for i := 1; i <= koyun; i++ {
		fmt.Print(i, " sheep... ")
	}
}
