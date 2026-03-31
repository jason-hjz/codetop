package _5_3Sum

import "sort"

// 排序，固定a值，双指针找b、c,
// c从数组末尾往前，b从a往后，重复值跳过

func ThreeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0) //记录结果三元组

	for a := 0; a < n; a++ {
		if nums[a] > 0 {
			continue
		}
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}

		c := n - 1
		for b := a + 1; b < n; b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}

			for b < c && nums[b]+nums[c] > (-1*nums[a]) {
				c--
			}

			if b == c {
				break
			}

			if nums[b]+nums[c] == (-1 * nums[a]) {
				ans = append(ans, []int{nums[a], nums[b], nums[c]})
			}
		}
	}

	return ans
}
