package main

import "fmt"

func main() {
	var x byte
	fmt.Scanf("%c", &x)
	if x > 64 && x <= 90 {
		fmt.Println(string(x), "adalah huruf kapital")
	} else if x > 96 && x <= 122 {
		fmt.Println(string(x), "bukan huruf kapital")
	} else {
		fmt.Println("Inputan bukan huruf!")
	}
}
