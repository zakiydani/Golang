package main

import "fmt"

type Product struct {
	 Name string
	 Price int
	 Stock int
}

func main() {
	var Item = make(map[int]Product)

	Item[0] = Product{Name:"Masker", Price:10000, Stock:9}
	Item[1] = Product{Name:"Sarung Tangan", Price:20000, Stock:5}
	Item[2] = Product{Name:"Hand Sanitizer", Price:25000, Stock:8}
	Item[3] = Product{Name:"Sabun", Price:3000, Stock:30}
	Item[4] = Product{Name:"Shampoo", Price:17000, Stock:50}

	fmt.Println("Barang yang memiliki Stok di bawah 10 pcs adalah :")
	
	for _, qty := range Item {
		if qty.Stock < 10 {
			fmt.Println(">", qty.Name, qty.Stock)
		}
	}
}