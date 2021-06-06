package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/add-two-numbers/
//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

//请你将两个数相加，并以相同形式返回一个表示和的链表。

//你可以假设除了数字 0 之外，这两个数都不会以 0 开头。


 type ListNode struct {
 		Val int
     Next *ListNode
 }

func addTwoNumbers(l1c *ListNode, l2c *ListNode) *ListNode {

	a:=make([]int,5)
	b:=make([]int,5)


	one:=l1c
	a= append(a, one.Val)
	for{
		fmt.Print(one.Val)
		one=one.Next
		if one==nil {
			break
		}
		a= append(a, one.Val)
	}

	two:=l2c
	b= append(b, two.Val)
	for{
		fmt.Print(two.Val)
		two=two.Next
		if two==nil {
			break
		}
		b= append(b, two.Val)
	}

	alen:=len(a)
	blen:=len(b)
	var length int
	if alen>=blen{
		length= alen
	}else {
		length=blen
	}


	c:=make([]int,0)
	for i:=length-1;i>=0;i--{
		var a1 = 0
		if i<length {
			a1=a[i]
		}else {
			a1=0
		}

		var b1 = 0
		if i<length {
			b1=b[i]
		}else {
			b1=0
		}
		c1:=a1+b1
		if c1!=0{
			c=append(c,a1+b1)
		}
	}

	var last = false
	for index, i2 := range c {
		if i2 >=10 {
			if index<len(c)-1{
				c[index]=i2%10
				c[index+1]=c[index+1]+1
			}else{
				c[index]=i2%10
				last=true
			}
		}
	}

	var resultLast ListNode


	var head *ListNode
	if last {
		resultLast =ListNode{Val: 1,Next: nil}
		head=&resultLast
	}

	for i:=len(c)-1;i>=0;i--{
		temp:=head
		newNode :=ListNode{Val: c[i],Next: nil}
		head=&newNode
		newNode.Next=temp
	}
	return head
}

func main() {
	l1a:=ListNode{Val: 1}
	l1b:=ListNode{Val: 1,Next: &l1a}
	l1c:=ListNode{Val:9,Next: &l1b}

	l2a:=ListNode{Val: 1}
	l2b:=ListNode{Val: 8,Next: &l2a}
	l2c:=ListNode{Val:1,Next: &l2b}

	addTwoNumbers(&l1c,&l2c)



}


