package linkedlist

type LinkedList struct {
	Node *LinkedNode
}

type LinkedNode struct {
	Data interface{}
	Next *LinkedNode
}

func Reverse(list *LinkedList) *LinkedList {
	// 处理边界条件
	if list.Node == nil || list.Node.Next == nil {
		return list
	}

	// 遍历
	var pre, cur, next *LinkedNode

	pre = nil
	cur = list.Node
	next = cur.Next

	cur.Next = pre

	for {
		if next.Next == nil {
			next.Next = cur
			list.Node = next
			return list
		}

		pre = cur
		cur = next
		next = next.Next

		cur.Next = pre
	}
}

// pre cur next
// nil- n1-> n2-> n3
// nil<-n1-> n2-> n3

//      pre  cur  next
// nil<-n1-> n2-> n3
// nil<-n1<- n2-> n3
