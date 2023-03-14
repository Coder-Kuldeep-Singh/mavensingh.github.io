package main

import (
	"fmt"
)

func sumFor(arr []int) int {
	total := 0
	for _, num := range arr {
		total += num
	}
	return total
}

func sumReduce(arr []int) int {
	return reduceInt(func(r, v int) int { return r + v }, arr)
}

func reduceInt(f func(int, int) int, arr []int) int {
	r := arr[0]
	for _, v := range arr[1:] {
		r = f(r, v)
	}
	return r
}

func main() {
	smallArr := []int{1, 2, 3, 4, 5}
	mediumArr := make([]int, 1000)
	for i := range mediumArr {
		mediumArr[i] = i + 1
	}

	largeArr := make([]int, 100000)
	for i := range largeArr {
		largeArr[i] = i + 1
	}

	fmt.Println("Small Array:")
	fmt.Println("For Loop: ", sumFor(smallArr))
	fmt.Println("Reduce: ", sumReduce(smallArr))

	fmt.Println("\nMedium Array:")
	fmt.Println("For Loop: ", sumFor(mediumArr))
	fmt.Println("Reduce: ", sumReduce(mediumArr))

	fmt.Println("\nLarge Array:")
	fmt.Println("Reduce: ", sumReduce(largeArr))
	fmt.Println("For Loop: ", sumFor(largeArr))
}
