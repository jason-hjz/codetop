package _15_KthLargest

// 找第K大数、建大根堆、大根堆化

func maxHeapify(a []int, i, heapsize int) {
	l, r, largest := i*2+1, i*2+2, i // 左右子节点下标和最大值下标，i为父节点下标
	if l < heapsize && a[l] > a[largest] {
		largest = l
	}
	if r < heapsize && a[r] > a[largest] {
		largest = r
	}

	if largest != i {
		a[largest], a[i] = a[i], a[largest]
		maxHeapify(a, largest, heapsize) //递归调整最大子节点对应的子树
	}
}

func buildmaxHeap(a []int, heapsize int) {
	// 最后一个非叶子节点的下标：heapSize/2 - 1
	// 从该节点向前遍历到根节点（下标 0），逐个调整堆
	for i := heapsize/2 - 1; i >= 0; i-- {
		maxHeapify(a, i, heapsize)
	}
}

func findKthLargest2(nums []int, k int) int {
	heapsize := len(nums)
	buildmaxHeap(nums, heapsize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapsize--
		maxHeapify(nums, 0, heapsize)
	}
	return nums[0]
}
