package main

import (
	"math/rand"
	"time"
	"fmt"
)

func bubbleSort(arr []int, n int) {
	var swapped bool = true
	for swapped {
		swapped = false
		for i := 1; i < n; i++{
			if arr[i-1] > arr[i] {
				arr[i-1],arr[i] = arr[i],arr[i-1]
				swapped = true
			}
		}
		n--
	}
}

func main(){

	const N = 2000
	arr := [N]int{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < N; i++{
		arr[i] = r.Intn(N)
	}
	start := time.Now()

	bubbleSort(arr[:],N)
	//fmt.Println(arr)
	end   := time.Now()
	fmt.Println(end.Sub(start))

}
