package main

import (
	"math/rand"
	"time"
	"fmt"
	"joewt.com/joe/learngo/leetcode/datastruct/sort/quick/quick"
)


func main(){
	const N = 200000
	arr := [N]int{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < N; i++{
		arr[i] = r.Intn(10)
	}
	start := time.Now()
	//quick.QuickSort(arr[:],N)
	//fmt.Println(arr)
	end   := time.Now()
	fmt.Println("quick sort:",end.Sub(start))

	start = time.Now()
	quick.QuickSort2(arr[:],N)
	//fmt.Println(arr)
	end   = time.Now()
	fmt.Println("quick sort2:",end.Sub(start))


	start = time.Now()
	quick.QuickSort3Ways(arr[:],N)
	//fmt.Println(arr)
	end   = time.Now()
	fmt.Println("quick sort2:",end.Sub(start))
}
