package sort

func sort(intervals [][]int) [][]int {
	length := len(intervals)
	if length <= 1 {
		return intervals
	}
	leftArr := sort(intervals[0 : length/2])
	rightArr := sort(intervals[length/2 : length])
	resArr := [][]int{}

	for left, right := 0, 0; left < length/2 || right < length-length/2; {

		if left >= length/2 {
			resArr = append(resArr, rightArr[right])
			right++
			continue
		}
		if right >= length-length/2 {
			resArr = append(resArr, leftArr[left])
			left++
			continue
		}

		if leftArr[left][0] < rightArr[right][0] || (leftArr[left][0] == rightArr[right][0] && leftArr[left][1] <= rightArr[right][1]) {
			resArr = append(resArr, leftArr[left])
			left++
		} else {
			resArr = append(resArr, rightArr[right])
			right++
		}
	}
	return resArr
}

func merge(intervals [][]int) [][]int {
	sortedArr := sort(intervals)
	resArr := [][]int{}
	if len(sortedArr) == 0 {
		return resArr
	}
	leftVal, rightVal := sortedArr[0][0], sortedArr[0][1]
	for _, item := range sortedArr {
		if item[0] > rightVal {
			resArr = append(resArr, []int{leftVal, rightVal})
			leftVal, rightVal = item[0], item[1]
		} else if item[1] > rightVal {
			rightVal = item[1]
		}
	}
	resArr = append(resArr, []int{leftVal, rightVal})
	return resArr
}
