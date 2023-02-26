package main

import "fmt"

func main() {
	var n rune
	var a, b, c int
	fmt.Scanf("%c ", &n)
	for n == 65 || n == 66 || n == 67 {
		if n == 65 {
			a += 1
			fmt.Scanf("%c ", &n)
			continue
		} else if n == 66 {
			b += 1
			fmt.Scanf("%c ", &n)
			continue
		} else if n == 67 {
			c += 1
			fmt.Scanf("%c ", &n)
			continue
		}

	}
	fmt.Print("Tipe A = ", a, "\nTipe B = ", b, "\nTipe C = ", c)
}
