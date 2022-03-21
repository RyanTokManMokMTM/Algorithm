package Sorting

import "fmt"

/*
	InsertionSort TODO
	Base case is O(n)
	Time Complexity ： O(N^2) -> each time compare with sorted list 1 + 2 + 3 + 4...n  for n elements(each time pick one and compare ,for all elements is n)
		Total time :n*(n-1)/2 + n = O(n^2)
	Space Complexity : O(N)
	i = current sorting elements
	j = the last index of  sorted list
	ensure that before j+1 is sorted list
	sorting the list after j
	key current sorting element
	GIVEN A LIST = [5,3,1,2],j = -1 and i = 0
	5,3 -> 3,5,1,2
	5,1 -> 1,3,5,2
	5,2 -> 1,2,3,5

	AFTER comparison -> let key in  j+1 index

*/

func InsertionSort(list []int, sortByLessThan bool) {
	if !sortByLessThan {
		insertionGreater(list)
	} else {
		insertionLess(list)
	}
	fmt.Println(list)
}

func insertionGreater(list []int) {
	for i := 1; i < len(list); i++ {
		key := list[i]
		j := i - 1
		for {
			if j >= 0 && key > list[j] {
				list[j+1] = list[j]
				j--
			} else {
				break
			}
		}
		list[j+1] = key
	}
}

func insertionLess(list []int) {
	for i := 1; i < len(list); i++ {
		key := list[i]
		j := i - 1
		for {
			if j >= 0 && key < list[j] {
				list[j+1] = list[j]
				j--
			} else {
				break
			}
		}
		list[j+1] = key
	}
}
