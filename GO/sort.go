package main

import (
	"fmt"
	"math/rand"
	"time"
)

type T int

type Heap struct {
	data     []T
	count    int
	capacity int
}

func GenArr(n int, maxn int) []T {
	arr := make([]T, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		arr[i] = T(r.Intn(maxn))
	}
	return arr[:]
}

func BubbleSort(arr []T, n int) {
	var swap bool = true
	for swap {
		swap = false
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swap = true
			}
		}
		n--
	}
}

func InsertSort(arr []T, n int) {
	for i := 1; i < n; i++ {
		e := arr[i]
		j := 0
		for j = i; j > 0 && arr[j-1] > e; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = e
	}
}

func SelectSort(arr []T, n int) {
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func merge(arr []T, l int, mid int, r int) {
	aux := make([]T, r-l+1)
	for i := l; i <= r; i++ {
		aux[i-l] = arr[i]
	}
	i, j := l, mid+1
	for k := l; k <= r; k++ {
		if i > mid {
			arr[k] = aux[j-l]
			j++
		} else if j > r {
			arr[k] = aux[i-l]
			i++
		} else if aux[i-l] < aux[j-l] {
			arr[k] = aux[i-l]
			i++
		} else {
			arr[k] = aux[j-l]
			j++
		}
	}
}

func mergeSort(arr []T, l int, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)/2
	mergeSort(arr, l, mid)
	mergeSort(arr, mid+1, r)
	merge(arr, l, mid, r)
}

func MergeSort(arr []T, n int) {
	mergeSort(arr, 0, n-1)
}

func quickSort(arr []T, l int, r int) {
	if l >= r {
		return
	}
	ran := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr[l], arr[ran.Intn(r-l+1)+l] = arr[ran.Intn(r-l+1)+l], arr[l]
	v := arr[l]
	lt, gt, i := l, r+1, l+1
	for i < gt {
		if arr[i] < v {
			arr[i], arr[lt+1] = arr[lt+1], arr[i]
			lt++
			i++
		} else if arr[i] > v {
			arr[i], arr[gt-1] = arr[gt-1], arr[i]
			gt--
		} else {
			i++
		}
	}
	arr[l], arr[lt] = arr[lt], arr[l]
	quickSort(arr, l, lt-1)
	quickSort(arr, gt, r)
}

func QuickSort(arr []T, n int) {
	quickSort(arr, 0, n-1)
}

func (r *Heap) Size() int {
	return r.count
}

func (r *Heap) IsEmpty() bool {
	return r.count == 0
}

func (r *Heap) Insert(item T) {

	if r.capacity <= r.count {
		newdata := r.data
		r.data = make([]T, (r.count+1)*2)
		copy(r.data, newdata)
	}

	r.data[r.count+1] = item
	r.count++
	r.shiftUp(r.count)
}

func (r *Heap) shiftUp(k int) {
	for k > 1 && r.data[k/2] < r.data[k] {
		r.data[k/2], r.data[k] = r.data[k], r.data[k/2]
		k /= 2
	}
}

func (r *Heap) ExtractMax() T {
	ret := r.data[1]
	r.data[1], r.data[r.count] = r.data[r.count], r.data[1]
	r.count--
	r.shiftDown(1)
	return ret
}

func (r *Heap) shiftDown(k int) {
	for 2*k <= r.count {
		j := 2 * k
		if j+1 <= r.count && r.data[j+1] > r.data[j] {
			j += 1
		}

		if r.data[j] > r.data[k] {
			r.data[j], r.data[k] = r.data[k], r.data[j]
			k = j
		} else {
			break
		}
	}
}

func HeapSort(arr []T, n int) {
	heap := Heap{
		data:     make([]T, n+1),
		count:    0,
		capacity: n,
	}
	for i := 0; i < n; i++ {
		heap.Insert(arr[i])
	}
	for i := n - 1; i >= 0; i-- {
		arr[i] = heap.ExtractMax()
	}
}

func TimeFunc(arr []T, sort func(arr []T, n int)) time.Duration {
	start := time.Now()
	sort(arr, len(arr))
	end := time.Now()
	return end.Sub(start)
}

func main() {
	const N = 10000
	const MAX = 1000

	arr := GenArr(N, MAX)
	t := TimeFunc(arr, BubbleSort)
	fmt.Printf("BubbleSort time (%d) : %v\n", len(arr), t)

	arr = GenArr(N, MAX)
	t = TimeFunc(arr, BubbleSort)
	fmt.Printf("InsertSort time (%d) : %v\n", len(arr), t)

	arr = GenArr(N, MAX)
	t = TimeFunc(arr, SelectSort)
	fmt.Printf("SelectSort time (%d) : %v\n", len(arr), t)

	arr = GenArr(N, MAX)
	t = TimeFunc(arr, MergeSort)
	fmt.Printf("MergeSort time (%d) : %v\n", len(arr), t)

	arr = GenArr(N, MAX)
	t = TimeFunc(arr, QuickSort)
	fmt.Printf("QuickSort time (%d) : %v\n", len(arr), t)

	arr = GenArr(N, MAX)
	t = TimeFunc(arr, HeapSort)
	fmt.Printf("HeapSort time (%d) : %v\n", len(arr), t)
}
