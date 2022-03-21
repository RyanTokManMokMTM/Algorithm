package Sorting

/*
	TODO QuickSort - Idea -> partition
	strategy : from the list pick a pivot and make left list last than pivot and right list greater than pivot
	Time : O(n * logN) -> swap and sorting time is O(N) step,for sorting,  it takes logN time,each time O(N)
		=> n*logN

*/

func QuickSort(list []int, front, end int) {
	if front >= end {
		return
	}

	// partition and get pivot
	pivot := partition(list, front, end) //do first then next partition
	QuickSort(list, front, pivot-1)      //left part
	QuickSort(list, pivot+1, end)        //right part

}

//compare and swap the value base on pivot
func partition(list []int, front, end int) int {
	pivot := list[end] //set pivot as the last elements of current list
	i := front - 1
	for j := front; j < end; j++ { //end is pivot ,no need to compare
		if list[j] < pivot {
			i++
			swap(i, j, list)
		}
	}

	//current i is the last elements of left list
	//swap pivot and i and pivot will in the middle of the list
	i++
	swap(i, end, list)
	return i //return pivot
}
