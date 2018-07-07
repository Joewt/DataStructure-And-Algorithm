package maxheap

type MaxHeap struct {
	data []int
	count int
	capacity int
}


func (r *MaxHeap) Construction(n int) {
	r.data = make([]int, n+1)
	r.count = 0
	r.capacity = n
}


func (r *MaxHeap) Construction2(arr []int, n int) {
	r.data = make([]int,n+1)
	r.capacity = n
	for i := 0; i < n; i++ {
		r.data[i+1] = arr[i]
	}
	r.count = n

	for i := r.count/2; i > 0; i-- {
		r.shiftDown(i)
	}
}


func (r MaxHeap) Size() int {
	return r.count
}


func (r MaxHeap) IsEmpty() bool {
	return r.count == 0
}


func (r *MaxHeap) Insert(item int) {
	if r.capacity <= r.count {
		newdata := r.data
		r.data = make([]int, (r.count+1)*2)
		copy(r.data,newdata)
	}
	r.data[r.count+1] = item
	r.count++
	r.shiftUp(r.count)
}


func (r *MaxHeap) shiftUp(k int) {
	for k > 1 && r.data[k/2] < r.data[k]  {
		r.data[k/2],r.data[k] = r.data[k],r.data[k/2]
		k /= 2
	}
}


func (r *MaxHeap) ExtractMax() int{
	ret := r.data[1]
	r.data[1],r.data[r.count] = r.data[r.count],r.data[1]
	r.count--
	r.shiftDown(1)
	return ret
}


func (r MaxHeap) shiftDown(k int) {
	for 2*k <= r.count {
		j := 2*k
		if j+1 <= r.count && r.data[j+1] > r.data[j] {
			j+=1
		}
		if r.data[j] > r.data[k] {
			r.data[j],r.data[k]=r.data[k],r.data[j]
			k = j
		} else {
			break
		}
	}
}