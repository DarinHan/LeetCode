package main

import "testing"

func TestGetInsertPoint(t *testing.T) {
	slices := [][]int{}
	point := GetInsertPoint(slices, []int{-1, 0, -1})
	t.Log(point)
	if point != 0 {
		t.Fatal("error1", point)
	}
	slices = [][]int{{-3, -1, 4}}
	point = GetInsertPoint(slices, []int{-1, 0, -1})
	t.Log(point)
	if point != 1 {
		t.Fatal("error2", point)
	}

	slices = [][]int{{-3, -1, 4}}
	point = GetInsertPoint(slices, []int{-4, 0, -1})
	t.Log(point)
	if point != 0 {
		t.Fatal("error3", point)
	}
}

func TestAdd(t *testing.T) {
	slices := [][]int{}
	slices = Add(slices, []int{-2, 0, -1})
	if slices[0][0] != -2 {
		t.Fatal("error1")
	}
	t.Log(slices)
	slices = Add(slices, []int{-3, -1, 4})
	if slices[0][0] != -3 {
		t.Fatal("error2")
	}
	t.Log(slices)
	slices = Add(slices, []int{-1, -1, 4})
	if slices[2][0] != -1 {
		t.Fatal("error3")
	}
	t.Log(slices)
}

func TestSortAndRemoveDup(t *testing.T) {
	slices := []int{-1, 0, 1, 2, -1, -4}
	newslice := SortAndRemoveDup(slices, 3)
	t.Log(newslice)
	if newslice[0] != -4 {
		t.Fatal("error1")
	}

	slices = []int{1, -1, 0, -1}
	newslice = SortAndRemoveDup(slices, 3)
	t.Log(newslice)
	t.Log(slices)
	if newslice[0] != -1 {
		t.Fatal("error1")
	}
}

func TestThreeSum4(t *testing.T) {
	slices := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum4(slices)
	t.Log(result)
	if len(result) < 2 {
		t.Fatal("error1")
	}

	slices = []int{0,0,0}
	result = threeSum4(slices)
	t.Log(result)
	if len(result) > 1 {
		t.Fatal("error1")
	}
}
