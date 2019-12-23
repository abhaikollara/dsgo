package main

import "fmt"

type ListNode struct {
	num  int
	next *ListNode
}

type List struct {
	head *ListNode
}

func (l *List) add(n int) {
	newNode := ListNode{n, nil}
	if l.head == nil {
		l.head = &newNode
	} else {
		curNode := l.head
		for curNode.next != nil {
			curNode = curNode.next
		}
		curNode.next = &newNode
	}
}

func (l *List) delete(n int) {
	if l.head.num == n {
		l.head = l.head.next
	} else {
		cur := l.head
		for cur.next != nil {
			if cur.next.num == n {
				cur.next = cur.next.next
				break
			} else {
				cur = cur.next
			}
		}
	}
}

func (l List) print() {
	cur := l.head
	for cur != nil {
		fmt.Print(cur.num, "->")
		cur = cur.next
	}
	fmt.Print("nil")
	fmt.Println()
}

func main() {
	n := List{nil}
	n.add(3)
	n.add(7)
	n.add(19)
	n.print()
	n.delete(7)
	n.print()
}
