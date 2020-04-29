package main

import "fmt"

func main() {
	 search := find("Charger")

	if search {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}
}

func find(value string) bool {
	array := [5]string{"Money", "Earphone", "Wallet", "Handphone", "Charger"}

	for _, item := range array {
		if item == value {
			return true
		}
	}
	return false
}