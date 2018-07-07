package main

import "fmt"

func binarySearch(arr []int, n int, target int)int{
	l, r := 0, n-1
	for l <= r{
		mid := l + (r-l)/2
		if arr[mid] == target{
			return mid
		}
		if arr[mid] < target {
			l = mid+1
		}
		if arr[mid] > target {
			r = mid-1
		}
	}
	return -1
}

func main(){
	arr := make([]int,10)
	for i := 0; i < 10; i++{
		arr[i] = i
	}

	arr1 := []int{1,2,3,4,5,6,7,9,10}

	fmt.Println(binarySearch(arr,10,4))
	fmt.Println(binarySearch(arr1,10,4))

}
