// 11. Container With Most Water

// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

// Find two lines that together with the x-axis form a container, such that the container contains the most water.

// Return the maximum amount of water a container can store.

// Notice that you may not slant the container.

// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
// Example 2:

// Input: height = [1,1]
// Output: 1

// Constraints:

// n == height.length
// 2 <= n <= 105
// 0 <= height[i] <= 104

package main

import (
	"fmt"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxArea(height []int) int {
	mArea := 0
	area := 0
	for x := 0; x < len(height); x++ {
		if height[x] > 0 {
			for i := x + 1; i < len(height); i++ {
				if height[i] > 0 {

					minHeight := min(height[x], height[i])
					area = minHeight * (i - x)

					if area > mArea {
						mArea = area
					}

					fmt.Printf("minHeight: %d, i: %d, mArea: %d \n", minHeight, i-x, area)

				}
			}
		}

	}
	return mArea
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	ma := maxArea(height)
	fmt.Printf("ma: %d\n", ma)

	height1 := []int{1, 1}
	ma1 := maxArea(height1)
	fmt.Printf("ma: %d\n", ma1)
	fmt.Println(ma1)

	// const name, dept = "GeeksforGeeks", "CS"
	// fmt.Printf("%s is a %s Portal.\n", name, dept)

}
