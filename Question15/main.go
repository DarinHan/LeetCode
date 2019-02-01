package main

import (
	"fmt"
)

func main() {
	nums := []int{-1, -2, -3, 4, 1, 3, 0, 3, -2, 1, -2, 2, -1, 1, -5, 4, -3}
	result := threeSum4(nums)
	fmt.Printf("%v", result)
}

func threeSum2(nums []int) [][]int {
	//prepare
	nums=SortAndRemoveDup(nums,3)


	length := len(nums)
	maps := make(map[int][]int, length)
	result := make(map[string][]int, length)
	for index, val := range nums {
		sli := make([]int, 0)
		if sli2, ok := maps[val]; ok {
			sli = sli2
		}
		sli = append(sli, index)
		maps[val] = sli
	}

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			target := 0 - nums[i] - nums[j]
			if sli, ok := maps[target]; ok {
				for _, valindex := range sli {
					if valindex > j {
						min := nums[i]
						max := nums[i]
						med := 0
						if nums[j] < min {
							min = nums[j]
						}

						if nums[valindex] < min {
							min = nums[valindex]
						}

						if nums[j] > max {
							max = nums[j]
						}

						if nums[valindex] > max {
							max = nums[valindex]
						}
						med = 0 - min - max

						key := fmt.Sprintf("%d%d%d", min, med, max)
						fmt.Println(key)
						if _, ok := result[key]; ok != true {
							result[key] = []int{min, med, max}
						}
					}
				}
			}
		}
	}

	ret := make([][]int, 0)

	for _, val := range result {
		ret = append(ret, val)
	}
	return ret
}

func threeSum(nums []int) [][]int {
	negative := make(map[int]int, len(nums))
	postive := make(map[int]int, len(nums))
	zero := make(map[int]int, len(nums))

	resultmap := map[string][]int{}
	for _, val := range nums {
		if val < 0 {
			if count, ok := negative[val]; ok {
				count = count + 1
				negative[val] = count
			} else {
				negative[val] = 1
			}
		} else if val > 0 {
			if count, ok := postive[val]; ok {
				count = count + 1
				postive[val] = count
			} else {
				postive[val] = 1
			}
		} else {
			if count, ok := zero[val]; ok {
				count = count + 1
				zero[val] = count
			} else {
				zero[val] = 1
			}
		}
	}

	for key, val := range negative {
		for pkey, pval := range postive {
			target := 0 - pkey - key
			if target < 0 {
				if _, ok := negative[target]; ok {
					if target == key {
						if val > 1 {
							data := []int{target, target, pkey}
							datakey := fmt.Sprintf("%v", data)
							if _, ok := resultmap[datakey]; !ok {
								resultmap[datakey] = data
							}
						}
					} else {
						if target < key {
							data := []int{target, key, pkey}
							datakey := fmt.Sprintf("%v", data)
							if _, ok := resultmap[datakey]; !ok {
								resultmap[datakey] = data
							}
						} else {
							data := []int{key, target, pkey}
							datakey := fmt.Sprintf("%v", data)
							if _, ok := resultmap[datakey]; !ok {
								resultmap[datakey] = data
							}
						}
					}
				}
			} else if target > 0 {
				if _, ok := postive[target]; ok {
					if target == pkey {
						if pval > 1 {
							data := []int{key, target, target}
							datakey := fmt.Sprintf("%v", data)
							if _, ok := resultmap[datakey]; !ok {
								resultmap[datakey] = data
							}
						}
					} else {
						if target < pkey {
							data := []int{key, target, pkey}
							datakey := fmt.Sprintf("%v", data)
							if _, ok := resultmap[datakey]; !ok {
								resultmap[datakey] = data
							}
						} else {
							data := []int{key, pkey, target}
							datakey := fmt.Sprintf("%v", data)
							if _, ok := resultmap[datakey]; !ok {
								resultmap[datakey] = data
							}
						}
					}
				}
			} else {
				if len(zero) > 0 {
					data := []int{key, 0, pkey}
					datakey := fmt.Sprintf("%v", data)
					if _, ok := resultmap[datakey]; !ok {
						resultmap[datakey] = data
					}
				}
			}
		}
	}

	if count, ok := zero[0]; ok && count > 2 {
		data := []int{0, 0, 0}
		datakey := fmt.Sprintf("%v", data)
		if _, ok := resultmap[datakey]; !ok {
			resultmap[datakey] = data
		}
	}

	ret := [][]int{}

	for _, val := range resultmap {
		ret = append(ret, val)
	}
	return ret
}

func Add(slices [][]int, target []int) [][]int {
	if len(slices) == 0 {
		slices = append(slices, target)
		return slices
	}

	point := GetInsertPoint(slices, target)

	if point == len(slices) {
		slices = append(slices, target)
		return slices
	}

	oldpoint := slices[point]
	if CompareSlice(oldpoint, target) == 0 {
		return slices //existed
	}

	slices = append(slices, target)
	for i := len(slices) - 1; i >= point+1; i-- {
		slices[i] = slices[i-1]
	}
	slices[point] = target

	return slices
}

/*
	返回需要插入的数组下标位置(位置可能会超出数组长度)
 */
func GetInsertPoint(slices [][]int, target []int) int {
	if slices == nil {
		return 0
	}

	length := len(slices)
	if length == 0 {
		return 0
	}

	point := 0

	for i := 0; i < length; i++ {
		com := CompareSlice(slices[i], target)
		if com < 0 {
			point = i + 1
		} else if com > 0 {
			break
		} else if com == 0 {
			point = i
			break
		}
	}

	return point
}

func CompareSlice(targetA []int, targetB []int) int {
	length := len(targetA)
	for i := 0; i < length; i++ {
		if targetA[i] > targetB[i] {
			return 1
		} else if targetA[i] < targetB[i] {
			return -1
		}
	}
	return 0
}

func threeSum3(nums []int) [][]int {
	nums=SortAndRemoveDup(nums,3)
	l := len(nums)
	var result [][]int
	for i := 0; i < l-1 && nums[i] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := l - 1
		target := 0 - nums[i]
		for ; j < k; {
			sum := nums[j] + nums[k]
			if sum == target {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				for j < k && nums[j-1] == nums[j] {
					j++
				}
				k--
				for j < k && nums[k+1] == nums[k] {
					k--
				}
			} else if sum > target {
				k--
				for j < k && nums[k+1] == nums[k] {
					k--
				}
			} else {
				j++
				for j < k && nums[j-1] == nums[j] {
					j++
				}
			}
		}
	}
	return result
}

func threeSum4(nums []int) [][]int {
	nums=SortAndRemoveDup(nums,3)
	l := len(nums)

	if l==3 {
		if nums[0]+nums[1]+nums[2]==0{
			return [][]int{nums}
		}

		return  [][]int{}
	}
	lastnegativeindex:=-1
	firstpostiveindex:=-1
	for index,value:=range nums{
		if value==0{
			if lastnegativeindex == -1{
				lastnegativeindex = index - 1
			}
		}else if value > 0 {
			if lastnegativeindex == -1{
				lastnegativeindex = index - 1
			}
			firstpostiveindex = index
			break
		}
	}

	fmt.Println("lastnegativeindex",lastnegativeindex)
	fmt.Println("firstpostiveindex",firstpostiveindex)

	var negative []int
	var zero []int
	var postive []int

	if lastnegativeindex==-1{
		//no negative values
		if firstpostiveindex>2{
			return [][]int{{0,0,0}}
		}else{
			return [][]int{}
		}
	}else{
		negative=nums[0:lastnegativeindex+1]
	}

	if firstpostiveindex==-1{
		if len(nums)-lastnegativeindex>2{
			return [][]int{{0,0,0}}
		}else{
			return [][]int{}
		}
	}else {
		postive=nums[firstpostiveindex:]
	}

	zero=nums[lastnegativeindex+1:firstpostiveindex]

	fmt.Println("negative",negative)
	fmt.Println("zero",zero)
	fmt.Println("postive",postive)
	var result [][]int
	lastvalue:=0
	for index,value:=range negative{
		if value==lastvalue{
			continue
		}
		lastpvalue:=0
		for pindex,pvalue:=range postive{
			if lastpvalue==pvalue{
				continue
			}
			target:=0-pvalue-value
			if target==0 && len(zero)>0{
				result=append(result, []int{value,0,pvalue})
			}else if target<0{
				if target>=value{
					for t:=index+1;t<len(negative);t++{
						if negative[t]==target{
							result=append(result, []int{value,target,pvalue})
							break;
						}
					}
				}
			}else{
				if target>=pvalue{
					for t:=pindex+1;t<len(postive);t++{
						if postive[t]==target{
							result=append(result, []int{value,target,pvalue})
							break;
						}
					}
				}
			}
			lastpvalue=pvalue
		}
		lastvalue=value
	}

	if len(zero)>2{
		result=append(result, []int{0,0,0})
	}

	return result
}

func SortAndRemoveDup(nums []int, maxdup int) []int{
	length := len(nums)
	//freshSlice := make([]int, length)

	if length <= 1 {
		return nums
	}

	dupcount := 0
	dupvalue := 0
	freshIndex := 0
	for i := 0; i < length; i++ {
		for j := length - 1; j > i; j-- {
			if nums[j] < nums[i] {
				tmp := nums[i]
				nums[i] = nums[j]
				nums[j] = tmp
			}
		}

		if i == 0 || (nums[i] != dupvalue && i > 0) {
			dupvalue = nums[i]
			dupcount = 1
		} else {
			if nums[i] == dupvalue {
				dupcount++
			}
		}

		if dupcount <= maxdup {
			nums[freshIndex]= dupvalue
			freshIndex++
		}
	}
	return nums[0:freshIndex]
}
