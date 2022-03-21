package Sorting

import "fmt"

/*
	BubbleSort -TODO
	Time complexity : o(N^2) comparison Time,compare time n + n-1 +n-2+...1,for all elements (n)
	=> compare and swap (n*n-1/2) + n(each elements)
	Space Complexity : o(N) -> total array size

	//suppose input list [4,5,1,3]
	i loop over n-1s time
	j loop over n-1-i time -> all elements after n-1-i is sorted

	4 -> 5 , 5 -> 1(swap) -> 5 -> 3(swap) => 4,1,3,5 j=4
	4 -> 1(swap), 4 -> 3(swap) => 1,3,4,5 j=3
	1-> 3  => 1 ,3,4,5 j = 2
	1 3 4 5 j = 1
*/
//True
func BubbleSort(list []int, sortByLessThan bool) {
	if !sortByLessThan {
		bubbleGreater(list)
	} else {
		bubbleLess(list)
	}
	fmt.Println(list)
}

func bubbleGreater(list []int) {
	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-1-i; j++ {
			if list[j] <= list[j+1] {
				swap(j, j+1, list)
			}
		}
	}
}

func bubbleLess(list []int) {
	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-1-i; j++ {
			if list[j] > list[j+1] {
				swap(j, j+1, list)
			}
		}
	}
}

func swap(x, y int, arr []int) {
	arr[x], arr[y] = arr[y], arr[x]
}
