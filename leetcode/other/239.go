package other

import (
	"container/list"
)

type segmentTree struct {
	l, r     int
	maxValue int
}

func initTree(l, r, p int, tree []segmentTree, nums []int) int {
	if l == r {
		tree[p] = segmentTree{
			l:        l,
			r:        r,
			maxValue: nums[l],
		}
		return nums[l]
	}
	mid, maxValue := (l+r)/2, 0
	lValue := initTree(l, mid, p<<1, tree, nums)
	rValue := initTree(mid+1, r, p<<1+1, tree, nums)
	if lValue > rValue {
		maxValue = lValue
	} else {
		maxValue = rValue
	}
	tree[p] = segmentTree{
		l:        l,
		r:        r,
		maxValue: maxValue,
	}
	return maxValue
}

func queryTree(l, r, p int, tree []segmentTree) int {
	if l == tree[p].l && r == tree[p].r {
		return tree[p].maxValue
	}
	mid := (tree[p].l + tree[p].r) / 2
	if r <= mid {
		return queryTree(l, r, p<<1, tree)
	} else if l > mid {
		return queryTree(l, r, p<<1+1, tree)
	} else {
		lValue := queryTree(l, mid, p<<1, tree)
		rValue := queryTree(mid+1, r, p<<1+1, tree)
		if lValue > rValue {
			return lValue
		}
		return rValue
	}

}

func maxSlidingWindow(nums []int, k int) []int {
	length := len(nums)
	var tree = make([]segmentTree, length*3)
	var res = []int{}
	initTree(0, length-1, 1, tree, nums)
	for i := 0; i <= length-k; i++ {
		res = append(res, queryTree(i, i+k-1, 1, tree))
	}
	return res
}

func maxSlidingWindow2(nums []int, k int) []int {

	length := len(nums)
	if k <= 1 || length <= 1 {
		return nums
	}
	queue, res := list.New(), []int{}
	queue.PushBack(nums[0])

	for i := 1; i < length; i++ {

		if nums[i] >= queue.Front().Value.(int) {
			queue.Init()
			queue.PushBack(nums[i])
		} else {
			queue.PushBack(nums[i])
		}
		if queue.Len() > k {
			queue.Remove(queue.Front())
			currentValue := queue.Front().Value.(int)
			for q := queue.Front(); q != nil; q = q.Next() {
				if q.Value.(int) > currentValue {
					currentValue = q.Value.(int)
					for queue.Front().Value.(int) < currentValue {
						queue.Remove(queue.Front())
					}
				}
			}
		}
		if i >= k-1 {
			res = append(res, queue.Front().Value.(int))
		}
	}

	return res
}
