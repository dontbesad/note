package sort

import (
	"fmt"
	"testing"
)

func Test56(t *testing.T) {
	res := merge([][]int{
		{1, 2},
		{0, 9},
		{9, 10},
		{11, 99},
	})
	fmt.Println(res)
}

func Test179(t *testing.T) {
	res := largestNumber([]int{
		9,
		10,
	})
	fmt.Println(res)
}

func TestNums(t *testing.T) {
	res := findDuplicates([]int{
		1,
		3,
		3,
	})
	fmt.Println(res)
}
