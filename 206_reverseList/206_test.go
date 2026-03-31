package _06_reverseList

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	val := []int{1, 2, 3, 4, 5, 6}
	List := initList(val)
	traverseList(reverseList(List))
}
