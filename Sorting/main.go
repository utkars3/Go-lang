package main

import "fmt"

//merge logic for the code
func merge(arr []int, st int, en int, mid int) {
	left := make([]int, mid-st+1)
	right := make([]int, en-mid)

	copy(left, arr[st:mid+1])
	copy(right, arr[mid+1:en+1])

	length1 := mid - st + 1
	length2 := en - mid

	final := st
	i := 0
	j := 0
	for i < length1 && j < length2 {
		if left[i] <= right[j] {
			arr[final] = left[i]
			final++
			i++
		} else {
			arr[final] = right[j]
			final++
			j++
		}
	}

	for i < length1 {
		arr[final] = left[i]
		final++
		i++
	}

	for j < length2 {
		arr[final] = right[j]
		final++
		j++
	}
}

//merge sort function
func mergeSort(arr []int, st int, en int) {
	if st >= en {
		return
	}

	mid := st + (en-st)/2
	mergeSort(arr, st, mid)
	mergeSort(arr, mid+1, en)
	merge(arr, st, en, mid)
}

//main function
func main() {
	arr := []int{5, 3, 4,432,32,544,35436,4}
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
