package main

import (
	"fmt"
	"math/rand"
	"time"
)

func selectSort(arr []int,n int){
	for i := 0; i < n; i++{
		var minIndex int = i
		for j := i+1; j < n; j++{
			if arr[j] < arr[minIndex]{
				minIndex = j
			}
		}
		arr[i],arr[minIndex] = arr[minIndex],arr[i]
	}
}

func main(){
	const N = 20000
	arr := [N]int{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < N; i++{
		arr[i] = r.Intn(N)
	}
	start := time.Now()

	selectSort(arr[:],N)

	end   := time.Now()
	fmt.Println(end.Sub(start))

}
