package main

import (
	"fmt"
	"strconv"
)

func getPermutation(n int, k int) string {
	nums := []int{}

	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}

	return helper(nums, k)
}

func helper(nums []int, k int) string {
	totalNums := 1
	// 计算以每个数字开头的序列个数
	for i := 1; i < len(nums); i++ {
		totalNums *= i
	}

	fmt.Println("totalNum", totalNums)
	for i := range nums {
		// 以当前数字开头的话，前面的序列的总个数
		sum := (i + 1) * totalNums

		sub := k - sum

		// 说明结果是以当前数字开头
		// 去掉当前选中的数字，递归去确定接下来的数字
		if sub <= 0 {
			arr := []int{}
			arr = append(arr, nums[:i]...)
			arr = append(arr, nums[i+1:]...)
			return strconv.Itoa(nums[i]) + helper(arr, k-i*totalNums)
		}
	}

	return ""
}
