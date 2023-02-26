package main

import "fmt"

func main() {
	var hp int
	var isSwarm bool
	fmt.Scanln(&hp)
	fmt.Scanln(&isSwarm)
	if hp >= 0 && hp <= 100 {
		if hp < 50 && hp > 0 {
			if isSwarm {
				fmt.Println("Membuat dinding")
				fmt.Println("Tambah HP")
			} else {
				fmt.Println("Tambah HP")
			}
		} else if hp == 0 {
			fmt.Println("Revive")
			fmt.Println("Membuat dinding")
		} else if isSwarm {
			fmt.Println("Membuat dinding")
		} else {
			fmt.Println("Tidak perlu mengeluarkan skill")
		}
	} else {
		fmt.Println("Range HP melebihi batas")
	}
}
