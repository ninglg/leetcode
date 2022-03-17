type MyLinkedList struct {
	val int
	next *MyLinkedList
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		val:  0,
		next: nil,
	}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	t := this
    
    if index < 0 {
		return -1
	}
    
	for i := -1; i < index; i++ {
		if t.next != nil {
			t = t.next
		} else {
			return -1
		}
	}

	return t.val
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	this.next = &MyLinkedList{
		val:  val,
		next: this.next,
	}

	this.val++
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	t := this

	for t.next != nil {
		t = t.next
	}

	t.next = &MyLinkedList{
		val:  val,
		next: nil,
	}

	this.val++
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	t := this
	for i := 0; i < index; i++ {
		if t.next == nil {
			return
		}

		t = t.next
	}

	t.next = &MyLinkedList{
		val:  val,
		next: t.next,
	}

	this.val++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	t := this
	j := t
	for i := -1; i < index; i++ {
		if t.next == nil {
			return
		}

		j, t = t, t.next
	}

	j.next = t.next

	this.val--
}
