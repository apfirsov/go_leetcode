//https://leetcode.com/problems/reverse-linked-list/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}

func makeData(s []int) *ListNode {
	var cur *ListNode
	for i := len(s) - 1; i >= 0; i-- {
		node := &ListNode{s[i], cur}
		cur = node
	}
	return cur
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	head := makeData(slice)
	newHead := reverseList(head)
	fmt.Println(newHead)
}
