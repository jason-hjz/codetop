package main

import "fmt"

// 输入数组，输出最值
func maxSubArray(nums []int) int {
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1]+nums[i] {
			nums[i] = nums[i-1] + nums[i]
		}

		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
