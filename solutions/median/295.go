package median

import (
	"container/heap"
)

// An MaxHeap is a min-heap of ints.
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MaxHeap) Root() int {
	if h.Len() > 0 {
		return (*h)[0]
	}
	panic("no root present")
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MinHeap) Root() int {
	if h.Len() > 0 {
		return (*h)[0]
	}
	panic("no root present")
}

type MedianFinder struct {
	left  *MaxHeap
	right *MinHeap
}

func Constructor() MedianFinder {
	left := &MaxHeap{}
	right := &MinHeap{}
	heap.Init(left)
	heap.Init(right)
	return MedianFinder{
		left:  left,
		right: right,
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.left.Len() == 0 {
		heap.Push(this.left, num)
		return
	}

	if num > this.left.Root() {
		heap.Push(this.right, num)
	} else {
		heap.Push(this.left, num)
	}
	this.Balance()
}

func (this *MedianFinder) Remove(num int) {
	if num > this.left.Root() {
		temp := make([]int, 0, this.right.Len())
		for {
			x := heap.Pop(this.right)
			if num == x {
				break
			}
			temp = append(temp, x.(int))
		}
		for _, x := range temp {
			heap.Push(this.right, x)
		}
	} else {
		temp := make([]int, 0, this.left.Len())
		for {
			x := heap.Pop(this.left)
			if num == x {
				break
			}
			temp = append(temp, x.(int))
		}
		for _, x := range temp {
			heap.Push(this.left, x)
		}
	}
	this.Balance()
}

func (this *MedianFinder) Balance() {
	if this.left.Len() > this.right.Len()+1 {
		heap.Push(this.right, heap.Pop(this.left))
	}

	if this.left.Len() < this.right.Len() {
		heap.Push(this.left, heap.Pop(this.right))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.left.Len() == this.right.Len() {
		return float64(this.left.Root()+this.right.Root()) / float64(2)
	} else {
		return float64(this.left.Root())
	}
}

// func main() {
// 	mf := Constructor()
// 	mf.AddNum(1)
// 	mf.AddNum(2)
// 	fmt.Println("Median", mf.FindMedian())

// 	mf.AddNum(3)
// 	fmt.Println("Median", mf.FindMedian())
// }
