package main

import "fmt"

func main() {
	a := [20]int{1,3,4,6,7,8,9,10,11,13,14,17,18,19,21,22,23,24}
	b := [20]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	fmt.Println(len(a), len(b))
}