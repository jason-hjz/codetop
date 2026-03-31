package _5_3Sum

import "testing"

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	ans := ThreeSum(nums)
	t.Log(ans)
}
