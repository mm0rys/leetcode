// Package main
// @Time : 2022/6/27 2:33 下午
// @User : liangxiaoqiang
// @Description :

package main

import "fmt"

// best
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		// 先判断目标值是否在hash table中
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		// 不在的话就将当前值和索引位置存入hash table
		m[nums[i]] = i
	}
	return nil
}

// O(n^2)
// 暴力破解
func twoSum2(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	//result := twoSum([]int{2, 7, 11, 15}, 9)
	result := twoSum2([]int{2, 7, 11, 15}, 9)
	fmt.Println(result)
}
