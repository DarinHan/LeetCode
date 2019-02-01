package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node1:=ParseNodeFromSlice([]int{2,4,3})
	node2:=ParseNodeFromSlice([]int{5,6,4})
	node3:=addTwoNumbers(node1,node2)
	PrintNode(node1)
	PrintNode(node2)
	PrintNode(node3)

}

func PrintNode(node *ListNode){
	for{
		if node==nil{
			break
		}
		fmt.Printf("%v ",node.Val)
		node=node.Next
	}
}

func ParseNodeFromSlice(slices []int) *ListNode  {
	if slices==nil ||len(slices)==0{
		return  nil
	}

	root:=&ListNode{Val:slices[0]}
	result:=root
	for i:=1;i<len(slices);i++{
		root.Next=&ListNode{Val:slices[i]}
		root=root.Next
	}
	return result
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1==nil{
		return l2
	}

	if l2==nil{
		return  l1
	}

	root:=&ListNode{}
	result:=root
	base:=0
	ret:=0
	for{
		if l1==nil && l2==nil{
			break
		}

		ret,base=addTwoNumbersInner(l1,l2,base)
		root.Val=ret

		if l1!=nil{
			l1=l1.Next
		}

		if l2!=nil{
			l2=l2.Next
		}

		if base>0 || l1!=nil || l2!=nil {
			//move next
			root.Next=&ListNode{Val:base}
			root=root.Next
		}
	}

	return result
}

func addTwoNumbersInner(l1 *ListNode, l2 *ListNode,base int) (int,int) {
	val1:=0
	val2:=0
	if l1!=nil {
		val1 =l1.Val
	}

	if l2!=nil{
		val2=l2.Val
	}

	val:=val1+val2+base
	plush:=0
	if val>=10{
		val=val-10
		plush=1
	}

	return  val,plush
}
