package quick

import (
	"math/rand"
	"time"
)

func swap(i *int, j *int){
	*i, *j = *j, *i
}

var ran = rand.New(rand.NewSource(time.Now().UnixNano()))


func quickSort3(arr []int, l int, r int){
	if l >= r{
		return
	}
	swap(&arr[l],&arr[ran.Intn(r-l+1)+l])
	v := arr[l]
	lt,gt,i := l,r+1,l+1
	for i < gt {
		if arr[i] < v{
			swap(&arr[i],&arr[lt+1])
			lt++
			i++
		}else if arr[i] > v{
			swap(&arr[i],&arr[gt-1])
			gt--
		}else{
			i++
		}
	}
	swap(&arr[l],&arr[lt])
	quickSort3(arr, l, lt-1)
	quickSort3(arr, gt, r)

}

func partition2(arr []int, l int, r int) int {
	swap(&arr[l],&arr[ran.Intn(r-l+1)+l])
	v := arr[l]
	i,j := l+1, r
	for {
		for i <= r && arr[i] < v{
			i++
		}
		for j >= l+1 && arr[j] > v{
			j--
		}
		if i > j{
			break
		}
		swap(&arr[i],&arr[j])
		i++
		j--
	}
	swap(&arr[j],&arr[l])
	return j
}

func partition(arr []int, l int, r int) int{
	swap(&arr[l],&arr[ran.Intn(r-l+1)+l])
	v := arr[l]
	j := l
	for i := l+1; i <= r; i++{
		if(arr[i] < v){
			j++
			swap(&arr[i],&arr[j])
		}
	}
	swap(&arr[l],&arr[j])
	return j
}

func quickSort(par func(arr []int, l int, r int) int,arr []int, l int, r int){
	if l >= r {
		return
	}
	p := par(arr, l, r)
	quickSort(par,arr, l, p-1)
	quickSort(par,arr, p+1, r)
}

func QuickSort(arr []int, n int) {
	quickSort(partition,arr, 0, n-1)
}

func QuickSort2(arr []int, n int){
	quickSort(partition2,arr, 0, n-1)
}

func QuickSort3Ways(arr []int, n int){
	quickSort3(arr , 0, n-1)
}