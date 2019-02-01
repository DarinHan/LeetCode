package main

import "testing"

func TestFindMedianIndexForArray(t *testing.T){
	index,value := FindMedianIndexForArray([]int{1})
	t.Log(index)
	t.Log(value)
	if index!=0 || value!=1{
		t.Fatal("err1")
	}

	index,value = FindMedianIndexForArray([]int{1,2,3})
	t.Log(index)
	t.Log(value)
	if index!=1 || value!=2{
		t.Fatal("err1")
	}

	index,value = FindMedianIndexForArray([]int{1,2,2,3})
	t.Log(index)
	t.Log(value)
	if index!=1 || value!=2{
		t.Fatal("err1")
	}
}


func TestPositionOfSortedArray(t *testing.T)  {
	source:=[]int{1, 3,5,8,10,12,15,84}
	value := PositionOfSortedArray(source,1)
	t.Log(value)
	if value!=0{
		t.Fatal("err1")
	}
	value= PositionOfSortedArray(source,46)
	t.Log(value)
	if value!=7{
		t.Fatal("err5")
	}
	value= PositionOfSortedArray(source,3)
	t.Log(value)
	if value!=1{
		t.Fatal("err2")
	}
	value= PositionOfSortedArray(source,10)
	t.Log(value)
	if value!=4{
		t.Fatal("err3")
	}
	value= PositionOfSortedArray(source,84)
	t.Log(value)
	if value!=7{
		t.Fatal("err4")
	}
}

func TestFindMedianSortedArrary(t *testing.T)  {
	source:=[]int{1, 3,5,8,10}
	value := FindMedianSortedArrary(source,1)
	t.Log(value)
	if value!=4{
		t.Fatal("err1")
	}

	source=[]int{1, 3,5,8,10}
	value = FindMedianSortedArrary(source,15)
	t.Log(value)
	if value!=6.5{
		t.Fatal("err2")
	}

	source=[]int{1, 3}
	value = FindMedianSortedArrary(source,2)
	t.Log(value)
	if value!=2{
		t.Fatal("err3")
	}
}

func TestFindMedianSortedArrays(t *testing.T)  {
	source1:=[]int{1, 3}
	source2:=[]int{2}
	value := findMedianSortedArrays2(source1,source2)
	t.Log(value)
	if value!=2{
		t.Fatal("err1")
	}
	source1=[]int{}
	source2=[]int{1}
	value = findMedianSortedArrays2(source1,source2)
	t.Log(value)
	if value!=1{
		t.Fatal("err4")
	}

	source1=[]int{1, 2}
	source2=[]int{3,4}
	value = findMedianSortedArrays2(source1,source2)
	t.Log(value)
	if value!=2.5{
		t.Fatal("err2")
	}

	source1=[]int{1, 9,10,12,15,23,29,41}
	source2=[]int{2,81}
	value = findMedianSortedArrays2(source1,source2)
	t.Log(value)
	if value!=13.5{
		t.Fatal("err3")
	}

	source1=[]int{1, 2,3,5}
	source2=[]int{4}
	value = findMedianSortedArrays2(source1,source2)
	t.Log(value)
	if value!=3{
		t.Fatal("err5")
	}

	source1=[]int{2,3}
	source2=[]int{1,4,5}
	value = findMedianSortedArrays2(source1,source2)
	t.Log(value)
	if value!=3{
		t.Fatal("err5")
	}

	source1=[]int{3,5}
	source2=[]int{1,2,4}
	value = findMedianSortedArrays2(source1,source2)
	t.Log(value)
	if value!=3{
		t.Fatal("err6")
	}

	source1=[]int{1,2,6}
	source2=[]int{3,4,5}
	value = findMedianSortedArrays2(source1,source2)
	t.Log(value)
	if value!=3.5{
		t.Fatal("err6")
	}

	source1=[]int{1,3,4}
	source2=[]int{2,5,6}
	value = findMedianSortedArrays2(source1,source2)
	t.Log(value)
	if value!=3.5{
		t.Fatal("err6")
	}
}