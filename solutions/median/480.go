package median

import (
	"container/heap"
	"sort"
)

func MedianSlidingWindow(nums []int, k int) []float64 {
	result := make([]float64, 0, len(nums)-k+1)
	cur := make([]int, 0, k)
	cur = append(cur, nums[:k]...)
	first := nums[0]
	mf := MedianFinder{}

	sort.Ints(cur)

	if len(cur)%2 == 0 {
		left := make([]int, 0, len(cur)/2)
		right := make([]int, 0, len(cur)-len(cur)/2)
		left = append(left, cur[:len(cur)/2]...)
		right = append(right, cur[len(cur)/2:]...)
		mf.left, mf.right = &MaxHeap{}, &MinHeap{}
		*mf.left = left
		*mf.right = right
	} else {
		left := make([]int, 0, (len(cur)+1)/2)
		right := make([]int, 0, len(cur)-(len(cur)+1)/2)
		left = append(left, cur[:(len(cur)+1)/2]...)
		right = append(right, cur[(len(cur)+1)/2:]...)
		mf.left, mf.right = &MaxHeap{}, &MinHeap{}
		*mf.left = left
		*mf.right = right
	}
	heap.Init(mf.left)
	heap.Init(mf.right)

	result = append(result, mf.FindMedian())

	for i := 1; i < len(nums)-k+1; i++ {
		mf.Remove(first)
		first = nums[i]
		mf.AddNum(nums[i+k-1])
		result = append(result, mf.FindMedian())
	}
	return result
}
