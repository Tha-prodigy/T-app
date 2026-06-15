package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i, j := 0, 0
	comMedian := []int{}
	for i < len(nums1) || j < len(nums2) {
		if i >= len(nums1) {
			comMedian = append(comMedian, nums2[j])
			j++

		} else if j >= len(nums2) {
			comMedian = append(comMedian, nums1[i])
			i++

		} else if nums1[i] < nums2[j] {
			comMedian = append(comMedian, nums1[i])

			i++

		} else {
			comMedian = append(comMedian, nums2[j])
			j++
		}
	}

	// fmt.Println(comMedian)
	n := len(comMedian)
	if n%2 == 0 {

		return float64((comMedian[n/2-1] + comMedian[n/2])) / 2
	}

	return float64(comMedian[n/2])

}

func main() {
	y := "bbaabb"
	fmt.Println(isPalindrome(y))
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i <j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--

	}
	return true

}
