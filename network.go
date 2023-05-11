package main

import (
	"fmt"
	"math"
)

func main() {
	towers := [][]int{{1, 2, 13}, {2, 1, 7}, {0, 1, 9}}
	radius := 2
	fmt.Print(bestCoordinate(towers, radius))
}

func bestCoordinate(towers [][]int, radius int) []int {
	var (
		max int   = 0
		ans []int = []int{0, 0}
	)

	for x := 0; x <= 50; x++ {
		for y := 0; y <= 50; y++ {
			sum := 0
			for _, v := range towers {
				d := math.Sqrt(float64((x-v[0])*(x-v[0]) + (y-v[1])*(y-v[1])))
				if d <= float64(radius) {
					sum += int(math.Floor(float64(v[2])) / (1 + d))
				}
				if sum > max {
					max = sum
					ans = []int{x, y}
				}
			}
		}
	}
	return ans
}
