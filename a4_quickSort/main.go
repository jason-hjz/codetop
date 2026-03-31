package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strconv"
	"strings"
)

func quickSort(nums []int) []int {
	if slices.IsSorted(nums) {
		return nums
	}

	i := partition(nums)
	quickSort(nums[:i]) //这种写法可以理解成 [）
	quickSort(nums[i+1:])
	return nums
}

func partition(nums []int) int {
	n := len(nums)
	i := rand.Intn(n)

	pivot := nums[i]
	nums[i], nums[0] = nums[0], nums[i]

	i, j := 1, n-1
	for {
		for i <= j && nums[i] < pivot {
			i++
		}
		for i <= j && nums[j] > pivot {
			j--
		}

		if i >= j {
			break
		}

		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	nums[0], nums[j] = nums[j], nums[0]
	return j
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := strings.Fields(scanner.Text())
	nums := make([]int, len(str))
	for i := 0; i < len(str); i++ {
		nums[i], _ = strconv.Atoi(str[i])
	}

	quickSort(nums)
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d ", nums[i])
	}
}
