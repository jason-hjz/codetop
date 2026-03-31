package _5_KGroupReverseList

import (
	"fmt"
	"testing"
)

func TestKGroupreverseList(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	List := initList(nums)
	List = reverseKGroup(List, 3)
	for List != nil {
		fmt.Println(List.Val)
		List = List.Next
	}
	return
}
