package main
//插入排序
import (
	"math/rand"
	"time"
	"fmt"
)

func insert(arr []int,n int) {
	for i:=1; i<n;i++{
		t := arr[i]
		j := 0
		for j = i; j > 0 && t < arr[j-1]; j--{
			arr[j] = arr[j-1]
		}
		arr[j] = t
	}
}

func main() {
	const N = 200
	arr := [N]int{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < N; i++{
		arr[i] = r.Intn(N)
	}
	start := time.Now()

	insert(arr[:],N)

	end   := time.Now()
	fmt.Println(end.Sub(start))
	fmt.Println(arr)
	//insert(arr[:],N)
	//fmt.Println(arr)
}
