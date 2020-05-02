package main

import (
	"fmt"
	// "math/rand"
)

func main() {
	array1 := [20]int{1,3,4,6,7,8,9,10,11,13,14,17,18,19,21,22,23,24,30,31}
	array2 := [20]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	
	// array1 := make([]int, 20)
	// array2 := make([]int, 20)

	array := []int{}

	for i := 0; i < len(array1); i++ {
		// array1[i]= rand.Intn(20)
		status := false
		for j := 0; j < len(array2); j++ {
			if array1[i] == array2[j] {
				status = true
			}
		}
		if !status {
			array = append(array, array1[i])
		}
	}


	for i := 0; i < len(array2); i++ {
		// array2[i] = rand.Intn(20)
		status := false
		for j := 0; j < len(array1); j++ {
			if array2[i] == array1[j] {
				status = true
			}
		}
		if !status {
			array = append(array, array2[i])
		}
	}

	fmt.Println("Array 1:", array1)
	fmt.Println("\nArray 2:", array2)

		fmt.Println("\nNilai yang berbeda dari Array 1 dan Array 2:", array)
	}