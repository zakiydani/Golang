package main

import "fmt"

func main() {
	s := make([]int, 0, 30)
	y := 30

	for x := 1; x < y; x++ {
		z := 0

		for a := 1; a < y; a++ {
			if x % a == 0 {
				z++
			}
		}
		if z == 2 && y > 1 {
			s = append(s, x)
		}
	}

	fmt.Println(s)
	
}