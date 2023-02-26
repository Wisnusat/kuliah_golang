package main

import "fmt"

func main() {
	var silver, gold, platinum, gold_user, silver_user, platinum_user, bil, total float64
	for total < 500.0 {
		fmt.Scanln(&bil)
		if bil >= 50.0 && bil <= 99.0 {
			silver += bil
			silver_user += 1.0
			total += bil
		} else if bil >= 100.0 && bil <= 200.0 {
			platinum += bil
			platinum_user += 1.0
			total += bil
		} else if bil > 200.0 {
			gold += bil
			gold_user += 1.0
			total += bil
		}
	}
	fmt.Println("Rata-rata gold user:", gold/gold_user)
	fmt.Println("Rata-rata platinum user:", platinum/platinum_user)
	fmt.Println("Rata-rata silver user:", silver/silver_user)
}
