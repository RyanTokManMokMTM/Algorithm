package Sorting

// build the head tree O(N)

//sorting the head

/*
	Note Heap: root(i) -left(2i)-right(2i+1)
	Max Heap : root valur is greater than its child values
	Min Heap : root value is less than its child values

	//Suppose
	index:|0|1|2|3|4|5|6|7|8|
	value:|9|7|8|6|4|2|3|5|1|

	Max Heap: |X|9|7|8|6|4|2|3|5|1| (size+1)
					(root i=1)
						9
		2i|			7		 8			|2i + 1
	4i|			6	  4	   2	3			|4i + 1
8i|			 5	   1							|8i + 1

*/
//func HeadSort(list []int) {
//	buildMaxHead(list)
//}
//
//func maxHeapify() {
//	left, root, right :=
//}
//
//func buildMaxHead(list []int) {
//	maxHead := make([]int, len(list)+1) //start from index 1
//
//	for i := 0; i < len(list); i++ {
//
//	}
//}
