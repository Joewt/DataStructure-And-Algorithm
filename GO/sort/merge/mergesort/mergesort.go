package mergesort

//将arr[r,mid] 和 arr[mid...r)进行归并
func merge(arr []int, l int, mid int, r int){
	aux := make([]int,r-l+1)
	for i := l; i <= r; i++{
		aux[i-l] = arr[i]
	}
	i,j := l,mid+1
	for k := l; k <= r; k++{
		if i > mid {
			arr[k] = aux[j-l]
			j++
		}else if j > r {
			arr[k] = aux[i-l]
			i++
		} else if aux[i-l] < aux[j-l]{
			arr[k] = aux[i-l]
			i++
		}else {
			arr[k] = aux[j-l]
			j++
		}
	}
}

//递归使用归并排序，arr[l,r]
func mergeSort(arr []int, l int, r int){
	if l >= r {
		return
	}
	mid := l + (r-l)/2
	mergeSort(arr, l,mid)
	mergeSort(arr,mid+1,r)
	//一个优化方法，当排序的是近乎有序的
	if arr[mid] > arr[mid+1]{
		merge(arr, l, mid, r)
	}

}

func MergeSort(arr []int, n int){
	mergeSort(arr, 0, n-1)
}
