package main

/*
	https://leetcode.com/problems/median-of-two-sorted-arrays/
	The overall run time complexity should be O(log (m+n)).
 */

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		return FindMedianForSingleSliceFloat(nums2)
	}

	if len(nums2) == 0 {
		return FindMedianForSingleSliceFloat(nums1)
	}

	for {
		m1 := FindMedianForSingleSliceFloat(nums1)
		m2 := FindMedianForSingleSliceFloat(nums2)

		if m1 > m2 {
			//m1 <-- m2 -->
			max:=Max(nums1[len(nums1)-1],nums2[len(nums2)-1])
			min := Min(nums1[0], nums2[0])

			if nums1[0] == min {
				nums1 = nums1[1:]
			} else if nums2[0] == min {
				nums2 = nums2[1:]
			}

			if nums1[len(nums1)-1] == max {
				nums1 = nums1[:len(nums1)-1]
			} else if nums2[len(nums2)-1] == max {
				nums2 = nums2[:len(nums2)-1]
			}
		}
	}

	return 0
}

func Max(int1, int2 int) int {
	if int1 >= int2 {
		return int1
	} else {
		return int2
	}
}

func Min(int1, int2 int) int {
	if int1 <= int2 {
		return int1
	} else {
		return int2
	}
}

type Median struct {
	Left        int
	Right       int
	MedianValue float64
	Index       int /* right index */
	IsEven      bool
}

type MedianIndex struct{
	IndexLeft int
	IndexRight int
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		return FindMedianForSingleSlice(nums2).MedianValue
	}

	if len(nums2) == 0 {
		return FindMedianForSingleSlice(nums1).MedianValue
	}

	step := 1
	var median1 Median
	var median2 Median

	toward := 0 //1-left,2-right
	for {
		if len(nums1) == 1 {
			if len(nums2) == 1 {
				return float64(nums1[0]+nums2[0]) / 2
			}
			return FindMedianSortedArrary(nums2, nums1[0])
		}
		if len(nums2) == 1 {
			return FindMedianSortedArrary(nums1, nums2[0])
		}
		median1 = FindMedianForSingleSlice(nums1)
		median2 = FindMedianForSingleSlice(nums2)
		position := PositionOfSortedArray(nums2, 0)
		if position < median2.Index {
			nums1 = nums1[step:]
			nums2 = nums2[:len(nums2)-step]
			if toward == 2 {
				step = step * 2
			} else if toward == 1 {
				step = step / 2
			}
		} else if position > median2.Index {
			nums1 = nums1[:len(nums1)-step]
			nums2 = nums2[step:]
			if toward == 1 {
				step = step * 2
			} else if toward == 2 {
				step = step / 2
			}
		} else {
			min := 0
			max := 0
			if median2.Left > median1.Left {
				min = median1.Left
			} else {
				min = median2.Left
			}

			if median2.Right > median1.Right {
				max = median2.Right
			} else {
				max = median1.Right
			}

			if median1.IsEven {
				return median1.MedianValue + median2.MedianValue*2 - float64(min+max)
			} else if median2.IsEven {
				return median1.MedianValue*2 + median2.MedianValue - float64(min+max)
			}
			return (median1.MedianValue*2 + median2.MedianValue*2 - float64(min+max)) / 2
		}
	}

	return 0
}
func FindMedianForSingleSliceFloat(nums []int) float64 {
	if len(nums)%2 == 1 {
		return float64(nums[len(nums)/2])
	} else {
		return float64(nums[len(nums)/2-1]+nums[len(nums)/2]) / 2
	}
}

func FindMedianForSingleSlice(nums []int) Median {
	if len(nums)%2 == 1 {
		return Median{
			Left:        nums[len(nums)/2],
			Right:       nums[len(nums)/2],
			MedianValue: float64(nums[len(nums)/2]),
			Index:       len(nums) / 2,
			IsEven:      true}
	} else {
		return Median{
			Left:        nums[len(nums)/2-1],
			Right:       nums[len(nums)/2],
			MedianValue: (float64(nums[len(nums)/2-1]+nums[len(nums)/2]) / 2),
			Index:       len(nums) / 2,
			IsEven:      false}
	}
}

func FindMedianSortedArrary(nums []int, singlenum int) float64 {
	index := PositionOfSortedArray(nums, singlenum)
	if index == 0 {
		length := len(nums)
		if length%2 == 1 {
			return float64(nums[length/2-1]+nums[length/2]) / 2
		} else {
			return float64(nums[length/2-1])
		}
	} else if index == len(nums) {
		length := len(nums)
		if length%2 == 1 {
			return float64(nums[length/2 ]+nums[length/2+1]) / 2
		} else {
			return float64(nums[length/2 ])
		}
	} else {
		length := len(nums)
		if length%2 == 1 {
			if index == length/2 || index == (length/2+1) {
				return float64(singlenum+nums[length/2]) / 2
			} else if index < length/2 {
				return float64(nums[length/2-1 ]+nums[length/2]) / 2
			} else {
				return float64(nums[length/2+1 ]+nums[length/2]) / 2
			}
		} else {
			if index == (length/2 - 1) {
				return float64(nums[length/2-1])
			} else if index == (length / 2) {
				return float64(singlenum)
			} else {
				return float64(nums[length/2])
			}
		}
	}
	return 0
}

func FindMedianIndexForArray(nums []int) (int, int) {
	length := len(nums)
	if length%2 == 1 {
		return (length - 1) / 2, nums[(length-1)/2]
	} else {
		return length/2 - 1, nums[(length)/2-1]
	}
}

/*
	给定一个排序的数组Arr和一个整数target，找出整数在数组中的位置
	如果所有数都大于给定的数，返回0
	如果所有数都小于给定的数，返回数组的长度
	如果数组中存在Index,当Arr[Index]=target 返回Index
	如果数组中存在Index，当Arr[Index]<target<Arr[index+1] 返回Index+1
 */
func PositionOfSortedArray(slices []int, target int) int {
	if slices == nil || len(slices) == 0 {
		return 0
	}

	if slices[0] >= target {
		return 0
	}

	if slices[len(slices)-1] == target {
		return len(slices) - 1
	}

	if slices[len(slices)-1] < target {
		return len(slices)
	}

	point := 0
	step := len(slices) / 2
	for {
		if slices[point] == target {
			return point
		}

		if slices[point] < target && target < slices[point+1] {
			return point + 1
		}
		if slices[point] > target {
			point = point - step
		} else {
			point = point + step
		}

		step = step / 2
	}
	return -1
}
