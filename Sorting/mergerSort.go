package Sorting

import "fmt"

/*
	MERGE SORT
	Time Complexity : n * log n => cn(root) = 1st
	divide time take O(n-1)
	compare and merge time 2-lists takes O(N) -> compare all elements from 2 lists
	it needs to merge log(n) list for merge * compare time = n*log(n)
	total : n-1 + n(total n elements) * log(n)[log n times]
	=>O(n-1 + n * log n) = O(log n)
*/

func MergeSort(list []int, front, end int) {

	if front < end {
		mid := front + (end-front)/2
		//partition

		MergeSort(list, front, mid) //left
		MergeSort(list, mid+1, end) //right
		merge(list, front, mid, end)

	}
}

func merge(x []int, front, mid, end int) {
	leftIndex, rightIndex := 0, 0 //left start index and right start index

	leftSubSize := mid + 1 - front //
	rightSubSize := end - mid

	var leftSub []int //start from 0
	var rightSub []int
	for i := 0; i < leftSubSize; i++ {
		leftSub = append(leftSub, x[front+i])
	}
	for i := 0; i < rightSubSize; i++ {
		rightSub = append(rightSub, x[mid+1+i])
	}

	i := front
	for {
		if leftIndex >= leftSubSize || rightIndex >= rightSubSize {
			break
		}
		if leftSub[leftIndex] <= rightSub[rightIndex] {
			x[i] = leftSub[leftIndex]
			leftIndex++
		} else {
			x[i] = rightSub[rightIndex]
			rightIndex++
		}
		i++
	}
	for {
		if leftIndex >= leftSubSize {
			break
		}
		x[i] = leftSub[leftIndex]
		i++
		leftIndex++
	}

	for {
		if rightIndex >= rightSubSize {

			break
		}
		x[i] = rightSub[rightIndex]
		i++
		rightIndex++
	}

}

//MergeSortWithTopDown Recursive
func MergeSortWithTopDown(list []int, front, end int) {
	if front >= end {
		return
	}

	mid := front + (end-front)/2
	//divide left
	MergeSortWithTopDown(list, front, mid)
	//divide right
	MergeSortWithTopDown(list, mid+1, end)
	//Merge
	topDown(list, front, mid, end)
}

func topDown(list []int, front, mid, end int) {
	//side of 2 list
	leftArrSize := mid + 1 - front
	rightArrSize := end - mid

	//copy the element
	var leftArr []int
	var rightArr []int
	for i := 0; i < leftArrSize; i++ {
		leftArr = append(leftArr, list[front+i])
	}
	for i := 0; i < rightArrSize; i++ {
		rightArr = append(rightArr, list[mid+1+i])
	}
	listIndex := front
	leftIndex, rightIndex := 0, 0

	//Compare and swap
	for {
		if leftIndex >= leftArrSize || rightIndex >= rightArrSize {
			break
		} //out of range

		//compare and swap
		if leftArr[leftIndex] <= rightArr[rightIndex] {
			list[listIndex] = leftArr[leftIndex]
			leftIndex++
		} else {
			list[listIndex] = rightArr[rightIndex]
			rightIndex++
		}
		listIndex++
	}

	for {
		if leftIndex >= leftArrSize {
			break
		}
		list[listIndex] = leftArr[leftIndex]
		listIndex++
		leftIndex++
	}

	for {
		if rightIndex >= rightArrSize {
			break
		}
		list[listIndex] = rightArr[rightIndex]
		listIndex++
		rightIndex++
	}
}

//MergeSortWithBottomUp Iteration
func MergeSortWithBottomUp(list []int, length int) {
	// -- TODO
	var sortArr []int = make([]int, length)
	// Merge i*2 and sort each time
	for i := 1; i < length; i *= 2 { //width of current array -1,2,4,8,16...for one list
		//Merge array with //if i=1 -> there have 2 element ,if i = 2 -> there are 4 elements etc
		for j := 0; j < length; j += 2 * i { //from left :0 to length ->each time compare i*2 and start at j+i*2
			if i > length-j {
				break
			} //current i is larger than length .suppose length is 4,current i is 8,8> 4-0 -> no more elements

			//j = left array stated index
			//i = merge array size
			//j+i = right array stared index
			//min(i,length - j - i).if count out of range else remind size
			BottomUp(list, sortArr, j, i, j+i, min(i, length-j-i))
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func BottomUp(list, newList []int, left, leftCount, right, rightCount int) {
	leftIndex, rightIndex, leftSize, rightSize, index := left, right, left+leftCount, right+rightCount, left
	fmt.Println(left, right)
	for {
		//2 list is already compared!
		if leftIndex >= leftSize && rightIndex >= rightSize {
			break
		}

		//Compare
		if leftIndex < leftSize && rightIndex < rightSize {
			if list[leftIndex] < list[rightIndex] {
				newList[index] = list[leftIndex]
				leftIndex++

			} else {
				newList[index] = list[rightIndex]
				rightIndex++
			}
		} else if leftIndex < leftSize {
			newList[index] = list[leftIndex]
			leftIndex++
		} else {
			newList[index] = list[rightIndex]
			rightIndex++
		}
		index++
	}

	//Put back to original List
	//HERE index is reached to the length
	for i := left; i < index; i++ {
		list[i] = newList[i]
	}
}
