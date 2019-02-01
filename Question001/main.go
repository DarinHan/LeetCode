package main

import (
	"fmt"
	"time"
)

func main() {

	nums:=[]int{3,2,4}
	result:=twoSum(nums,6)

	fmt.Println(result)
	time.Sleep(time.Minute * 2)
}

func twoSum(nums []int, target int) []int {
	maps:=make(map[int]int,len(nums))
	for index,val:= range nums{
		maps[val]=index
	}
	for index,val:= range nums {
		var target2 = target - val
		if _,ok:= maps[target2];ok  {
			if maps[target2] !=index{
				return []int{index,maps[target2]}
			}
		}
	}

	return nil
}