package main
import (
	"fmt"
	"sort"
)
func  bubbleSort(data  sort.Interface){
	r := data.Len()-1
	for i := 0; i < r ; i++{
		for j := r; j > i; j--{
			if data.Less(j, j-1){
				data.Swap(j, j-1)
			}
		}
	}
}
func insertSort(data sort.Interface){
	r := data.Len()-1
	for i := 1; i <= r; i++{
		for j := i; j > 0 && data.Less(j, j-1); j--{
			data.Swap(j, j-1)
		}
	}
}
func  selectSort(data sort.Interface){
	r := data.Len()-1
	for i := 0; i < r; i++{
		min := i
		for j:= i+1; j <= r; j++ {
			if data.Less(j, min) {  min = j }
		}
		data.Swap(i, min)
	}
}
func main(){
	nums := []int{120,33,44,1,23,90,87,13,57,43,42}
	//标准库qsort 正序(实际上是qsort结合heapsort，insertsort)
	sort.Ints(nums)
	fmt.Println(nums)
	//反序qsort
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println(nums)
	//正序bubble
	bubbleSort(sort.IntSlice(nums))
	fmt.Println(nums)
	//反序bubble
	bubbleSort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println(nums)
	//正序insert
	insertSort(sort.IntSlice(nums))
	fmt.Println(nums)
	//反序inert
	insertSort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println(nums)
	//正序select
	selectSort(sort.IntSlice(nums))
	fmt.Println(nums)
	//反序select
	selectSort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println(nums)

}
