package main

import "testing"

func TestParseNodeFromSlice(t *testing.T) {
	slices := []int{9,8}
	point := ParseNodeFromSlice(slices)
	t.Log(point)
	if point.Val!=9{
		t.Fail()
	}

	slices2 := []int{1}
	point = ParseNodeFromSlice(slices2)
	t.Log(point)
	if point.Val!=1{
		t.Fail()
	}
}

func TestAddTwoNumbers(t *testing.T){
	point1 := ParseNodeFromSlice([]int{9,8})
	point2 := ParseNodeFromSlice([]int{1})
	point3:=addTwoNumbers(point1,point2)
	PrintNode(point3)
}