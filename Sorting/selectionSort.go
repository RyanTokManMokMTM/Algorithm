package Sorting

/*
	TODO SelectSort
	Time : O(N^2) //  find the min in n + n-1 + n-2 .... + 1 time
	analyze : total finding step : (n*n-1/2) + total swap time n(each time find a min value and swap ,total n elements) = n*(n-1/2) + n = O(n^2)
	Space:O(N)

	How It work: Find the min value in the unsorted list and swap to the end of sorted List

*/
func SelectSort(list []int) {
	for i := 0; i < len(list); i++ {
		//Find the mind
		minIndex := i
		for j := i + 1; j < len(list); j++ {
			if list[j] < list[minIndex] {
				minIndex = j
			}
		}

		//swap value to current i
		list[i], list[minIndex] = list[minIndex], list[i]
	}
}
