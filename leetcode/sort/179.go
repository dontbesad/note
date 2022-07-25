package sort

import (
	"strconv"
)

// todo
func largestNumber(nums []int) string {

	for _, item := range nums {
		return strconv.FormatInt(int64(item), 10)
	}
	return ""
}
