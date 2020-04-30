package main

import "fmt"

// type Product struct {
// 	 Name string
// 	 Price int
// 	 Stock int
// }

func main() {

	nama := make(map[int]string)
	harga := make(map[int]float32)
	stok := make(map[int]int)

	nama[0] = "Masker"
	harga[0] = 10000
	stok[0] = 9

	nama[1] = "Sarung Tangan"
	harga[1] = 20000
	stok[1] = 5

	nama[2] = "Hand Sanitizer"
	harga[2] = 30000
	stok[2] = 3

	nama[3] = "Sabun"
	harga[3] = 5000
	stok[3] = 12

	nama[4] = "Shampoo"
	harga[4] = 17000
	stok[4] = 13
	
	fmt.Println("Barang yang memiliki stok dibawah 10 adalah :")

	for index, item := range stok {
		if item < 10 {
			fmt.Println(nama[index], "Rp",harga[index], "dengan stok tersisa", stok[index])

		}
	}


	// var Item = make(map[int]Product)

	// Item[0] = Product{Name:"Masker", Price:10000, Stock:9}
	// Item[1] = Product{Name:"Sarung Tangan", Price:20000, Stock:5}
	// Item[2] = Product{Name:"Hand Sanitizer", Price:25000, Stock:8}
	// Item[3] = Product{Name:"Sabun", Price:3000, Stock:30}
	// Item[4] = Product{Name:"Shampoo", Price:17000, Stock:50}

	// fmt.Println("Barang yang memiliki Stok di bawah 10 pcs adalah :")
	
	// for _, qty := range Item {
	// 	if qty.Stock < 10 {
	// 		fmt.Println(">", qty.Name, qty.Stock)
	// 	}
	// }
}