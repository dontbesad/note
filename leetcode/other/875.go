package main

import "fmt"

func checkSpeed(piles []int, speed int, h int) bool {
	sum := 0
	for _, v := range piles {
		sum += v / speed
		if v%speed > 0 {
			sum++
		}
	}
	if sum <= h {
		return true
	}
	return false
}

func minEatingSpeed(piles []int, h int) int {
	var minSpeed, maxSpeed = int(1), int(1e9)

	for minSpeed < maxSpeed {
		midSpeed := (minSpeed + maxSpeed) / 2
		if checkSpeed(piles, midSpeed, h) {
			maxSpeed = midSpeed
		} else {
			minSpeed = midSpeed + 1
		}
	}

	return minSpeed
}

func main() {
	piles, h := []int{3, 6, 7, 11}, 100
	fmt.Println(minEatingSpeed(piles, h))
}
