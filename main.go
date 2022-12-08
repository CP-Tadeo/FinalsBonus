package main

import (
	"fmt"
	//"math/rand"
	//"time"
)

func main() {
	givensize := 0
	fmt.Println("Indicate size:")
	fmt.Scanln(&givensize)
	if givensize == 0 {
	    fmt.Println("Size cannot be 0. Setting size to default: 10")
		givensize = 10
	}

	slice := make([]int, givensize, givensize) //makes slice to put elements to sort
	for i := 0; i < givensize; i++ {
		var element int
		fmt.Println("Value",i+1)
		fmt.Scanln(&element)
		slice[i] = element
	}
	fmt.Println("\n --- unsorted --- \n\n", slice)
	fmt.Println("\n--- sorted ---\n\n", MergeSort(slice), "\n")
}


// runs mergesort algorithm
func MergeSort(sourceslice []int) []int {

	if len(sourceslice) < 2 {
		return sourceslice
	}
	middle := (len(sourceslice)) / 2
	return Merger(MergeSort(sourceslice[:middle]), MergeSort(sourceslice[middle:]))
}

//merges left and right
func Merger(left, right []int) []int {

	size, leftcounter, rightcounter := len(left)+len(right), 0, 0
	slice := make([]int, size, size)
	count := 0

	for leftcounter < len(left) && rightcounter < len(right) {
		if left[leftcounter] <= right[rightcounter] {
			slice[count] = left[leftcounter]
			count, leftcounter = count+1, leftcounter+1
		} else {
			slice[count] = right[rightcounter]
			count, rightcounter = count+1, rightcounter+1
		}
	}
	for leftcounter < len(left) {
		slice[count] = left[leftcounter]
		count, leftcounter = count+1, leftcounter+1
	}
	for rightcounter < len(right) {
		slice[count] = right[rightcounter]
		count, rightcounter = count+1, rightcounter+1
	}

	return slice
}
