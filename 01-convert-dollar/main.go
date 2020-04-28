package main

import "fmt"


func main() {
	var dollar uint16 = 2

	var rupiah uint16 = 15518

	var konversi = dollar * rupiah
	
	fmt.Printf("2 dollar : Rp %d\n",konversi)
}