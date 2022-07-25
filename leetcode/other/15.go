package other

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	res := [][]int{}
	for lIdx := 0; lIdx+1 < length; lIdx++ {
		for rIdx := length - 1; lIdx+1 < rIdx; rIdx-- {

			l, r, idx := lIdx+1, rIdx-1, -1
			for l <= r {
				mid := (r + l) / 2
				if nums[lIdx]+nums[rIdx]+nums[mid] > 0 {
					r = mid - 1
				} else if nums[lIdx]+nums[rIdx]+nums[mid] < 0 {
					l = mid + 1
				} else {
					idx = mid
					break
				}
			}
			if idx != -1 {
				res = append(res, []int{nums[lIdx], nums[idx], nums[rIdx]})
			}
			for lIdx+1 < rIdx && nums[rIdx-1] == nums[rIdx] {
				rIdx--
			}
		}
		for lIdx+1 < length && nums[lIdx+1] == nums[lIdx] {
			lIdx++
		}
	}
	return res
}

func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	res := [][]int{}
	for i := 0; i < length-1; i++ {
		l, r, flag := i+1, length-1, false
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for l < r {
			if nums[l]+nums[r]+nums[i] == 0 && !flag {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				flag = true
			} else if nums[i]+nums[l]+nums[r] > 0 {
				r--
				flag = false
				continue
			} else if nums[i]+nums[l]+nums[r] < 0 {
				l++
				flag = false
				continue
			}
			l++
		}
	}
	return res
}
