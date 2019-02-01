package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T){
	point1 := lengthOfLongestSubstring("abcabcbb")
	if point1!=3{
		t.Fatal("err1")
	}
	point1 = lengthOfLongestSubstring("bbbbb")
	if point1!=1{
		t.Fatal("err2")
	}
	point1 = lengthOfLongestSubstring("pwwkew")
	if point1!=3{
		t.Fatal("err3")
	}
	point1 = lengthOfLongestSubstring("ohomm")
	if point1!=3{
		t.Fatal("err4")
	}
}