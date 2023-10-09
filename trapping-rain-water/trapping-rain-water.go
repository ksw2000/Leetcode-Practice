// https://leetcode.com/problems/trapping-rain-water/
// Time: O(2n) where n is the length of height
// Space: O(1)

package main

func trap(height []int) int {
    n := len(height)
    max := 0
    for i := 1; i < n; i++{
        if height[i] > height[max]{
            max = i
        }
    }

    sum := 0
    // left to max
    lm := height[0]
    for i := 1; i < max; i++{
        if height[i] > lm {
            lm = height[i]
        } else {
            sum += lm - height[i]
        }
    }

    // right to max
    lm = height[len(height) - 1]
    for i := len(height)-2; i > max; i--{
        if height[i] > lm {
            lm = height[i]
        }else{
            sum += lm - height[i]
        }
    }
    return sum
}