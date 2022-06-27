// Package main
// @Time : 2022/6/27 5:55 下午
// @User : liangxiaoqiang
// @Description :

package main

func main() {
	findMedianSortedArrays1([]int{1, 3}, []int{2})
}

// 二分法
// O(log(m+n))
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	if length%2 == 1 {
		return float64(getK(nums1, nums2, length/2+1))
	} else {
		return float64(getK(nums1, nums2, length/2)+getK(nums1, nums2, length/2+1)) / 2.0
	}
}

// 转化为寻找第k小数问题
func getK(nums1 []int, nums2 []int, k int) float64 {
	// 用来记录nums1 / nums2 的偏移位置
	index1, index2 := 0, 0
	for {
		// nums1已经遍历完
		if index1 == len(nums1) {
			return float64(nums2[index2+k-1])
		}
		// nums2已经遍历完
		if index2 == len(nums2) {
			return float64(nums1[index1+k-1])
		}
		// k为1 说明已经排除完k-1个比目标值小的数， 因此只需要返回两个数组的最小值
		if k == 1 {
			return float64(min(nums1[index1], nums2[index2]))
		}
		// 偏移值
		half := k / 2
		// 防止越界
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		// 如果是nums1更小 那就排除nums1中newIndex1左边所有值
		if nums1[newIndex1] <= nums2[newIndex2] {
			k -= newIndex1 - index1 + 1
			index1 = newIndex1 + 1
		} else {
			k -= newIndex2 - index2 + 1
			index2 = newIndex2 + 1
		}
	}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
