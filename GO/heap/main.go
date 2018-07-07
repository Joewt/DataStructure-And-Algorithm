package main

import (
	"math/rand"
	"time"
	"joewt.com/joe/learngo/leetcode/datastruct/heap/maxheap"
	"fmt"
)

type MaxHeap interface {
	Size() int
	IsEmpty() bool
	Construction(n int)
	Insert(item int)
	ExtractMax() int
	Construction2(arr []int, n int)
}


func HeapSort1(arr []int, n int) {
	var max MaxHeap
	maxheap := maxheap.MaxHeap{}
	max = &maxheap
	max.Construction(n)

	for i := 0; i < n; i++ {
		max.Insert(arr[i])
	}

	for i := n-1; i >= 0; i-- {
		arr[i] = max.ExtractMax()
	}

}

func HeapSort2(arr []int, n int) {
	maxheap := maxheap.MaxHeap{}
	maxheap.Construction2(arr, n)
	for i := n-1; i >= 0; i-- {
		arr[i] = maxheap.ExtractMax()
	}
}

func main() {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))


	const N = 20
	arr := [N]int{}
	for i := 0; i < N; i++{
		arr[i] = r.Intn(N)
	}

/*
	start := time.Now()
	HeapSort1(arr[:], N)
	end   := time.Now()
	fmt.Println(arr)
	fmt.Println("heap1 sort:",end.Sub(start))
*/



	start := time.Now()
	HeapSort2(arr[:], N)
	end   := time.Now()
	fmt.Println(arr)
	fmt.Println("heap2 sort:",end.Sub(start))



	/*
	max := maxheap.MaxHeap{}
	max.Construction(100)
	for i := 0; i < 107; i++ {
		max.Insert(r.Intn(100))
	}
	*/



}