package _15_KthLargest

import (
	"fmt"
	"testing"
)

func TestKthLargest(t *testing.T) {
	nums := []int{7, 4, 5, 6, 3, 8, 9}
	fmt.Println(findKthLargest(nums, 2))
}
