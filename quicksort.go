package main

import (
	"fmt"
)

func main() {
	check := []int {500,700,800,400,100,300,600,200,1000,900}
	fmt.Print("Input: ")
	print_t(check)
	quicksort(check,0,len(check)-1)

	fmt.Print("Output: ")
	print_t(check)
}


func partition(arr []int, p, r int) int {
	  pivot := arr[r]
	  for (p<r){
		  for (arr[p]<pivot){
			  p++
		  }
		  for (arr[r]>pivot){
			  r--
		  }
		  if arr[p]==arr[r]{
			  p++
		  } else if p<r {
			  tmp := arr[p]
			  arr[p] = arr[r]
			  arr[r] = tmp
		  }
	  }
	  return r
}

func quicksort(arr []int, p, r int) {
	if p<r {
		j := partition(arr,p,r)
		quicksort(arr,p,j-1)
		quicksort(arr,j+1,r)
	}
}

func print_t(arr []int){
	for _,x := range arr{
		fmt.Print(x)
		fmt.Print(" ")
	}
	fmt.Println()
}
