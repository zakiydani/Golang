package main

import "fmt"

func main(){
	angka := [30] int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	var even [] int
	var odd [] int


	for x := 0; x < angka[19]; x++ {
		if angka[x] % 2 == 0 {
			even = append(even, angka[x])
		} else {
			odd = append(odd, angka[x])
		}
	}

	var prime [] int

	for y := 1; y < (angka[19]); y++ {
		z := 0

		for a := 1; a < len(angka); a++ {
			if angka[y] % a == 0 {
				z++
			}
		}
		if z == 2 && y > 0 {
			prime = append(prime, angka[y])
		}
	}
	

	fmt.Println("Bilangan Ganjil:", odd, "\nJumlah bilangan Ganjil:", len(odd))
	fmt.Println("-------------------------------------------")
	fmt.Println("Bilangan Genap:", even, "\nJumlah bilangan Genap:", len(even))
	fmt.Println("-------------------------------------------")
	fmt.Println("Bilangan Prima:", prime, "\nJumlah bilangan Prima:", len(prime))

}


