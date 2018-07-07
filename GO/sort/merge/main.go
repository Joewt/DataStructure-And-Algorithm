package main

import (
	"math/rand"
	"time"
	"fmt"
	"joewt.com/joe/learngo/leetcode/datastruct/sort/merge/mergesort"
)

func main() {
	const N = 1000000
	arr := [N]int{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < N; i++{
		arr[i] = r.Intn(N)
	}
	start := time.Now()

	mergesort.MergeSort(arr[:],N)
	//fmt.Println(arr)
	end   := time.Now()
	fmt.Println("merge sort:",end.Sub(start))

}
