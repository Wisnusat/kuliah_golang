package main

import "fmt"

func deteksi_prima() {
	var x, i int
	var isPrime bool

	// input the number we want to check
	fmt.Print("Masukkan Bilangan: ")
	fmt.Scanln(&x)

	// declare i, i start from 2 because prime number is a number that doenst have any divisible
	// other than 1 and the number itself
	i = 2
	isPrime = true

	// if the number is less than 1 or equal to one, then it is not prime
	if x <= 1 {
		isPrime = false
	}

	// This for loop is checking that the current number have any divisible other than 1
	for i <= x-1 && isPrime {
		// if the value is equal zero (habis dibagi 2) then isPrime value will return false
		// and the loop will stop
		isPrime = x%i != 0
		i++
	}
	fmt.Println(isPrime)
}

func generate_prima() {
	var i, x int
	var isPrime bool

	// running the for loop from 1 to N (n value is 10, u can change to whatever u want)
	for i = 2; i <= 10; i++ {

		// flag which will confirm that the number is Prime or not
		isPrime = true

		// This for loop is checking that the current number have any divisible other than 1
		for x = 2; x <= i-1; x++ {
			if i%x == 0 {
				isPrime = false
				break
			}
		}
		// if the value of the flag is false then the number is not prime and we are not printing it.
		if isPrime {
			fmt.Println(i)
		}
	}
}

func main() {
	deteksi_prima()
	generate_prima()
}
