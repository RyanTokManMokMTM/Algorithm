package main

import (
	"BasicAlgorithm/Sorting"
	"fmt"
)

func main() {
	sortingList := []int{
		10, 1, 6, 5, 7, 9, 2,
	}
	Sorting.QuickSort(sortingList, 0, len(sortingList)-1)

	fmt.Println(sortingList)
}
