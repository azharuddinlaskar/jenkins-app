package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var matrix = make([][]int, 10)
	for i := range matrix {
		matrix[i] = make([]int, 10)
		for j := range matrix[i] {
			matrix[i][j] = rand.Intn(100)
		}
	}

	for i := range matrix {
		fmt.Println(matrix[i])
	}
}
