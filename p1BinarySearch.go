package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{2, 5, 7, 10, 12}
	var size = float64(len(arr))
	fmt.Println(binarySearch(size, arr, 2))
}


func binarySearch (n float64, a []int, key int) bool{
	// n = size(a) , a is the sorted array, key is the traget value
	var l float64 = 0
	var r float64 = (n-1.0)

	for l <= r {
		mid := math.Floor((l+r)/2.0)  
		var intmid = int(mid)

		if a[intmid] < key {
			l = mid + 1
		} else if a[intmid] > key {
			r = mid - 1
		} else {
			return true
		}
	}

	return false
}

