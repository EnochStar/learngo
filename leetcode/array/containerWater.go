package array

import "math"

func maxArea(height []int) int {
	left := 0
	right := len(height)
	var max float64
	for left < right {
		high := math.Min(float64(height[left]), float64(height[right]))
		max = math.Max(max, high*float64((right-left)))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return int(max)
}
