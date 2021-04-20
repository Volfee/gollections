// Package deque implements deque type and common methods to interact with
// deque.
package deque

import (
	"strconv"
	"strings"
)

// Deque is implemented as circular doubly linked list with sentinel node.
// This allows for O(1) non-amortized add/remove operations for both ends
// of the structure. This makes it a perfect implementation of queue data
// structure.
type Deque struct {
	sentinel *IntNode
	length   int
}

// Constructor for new empty deque.
func New() Deque {
	sentinel := &IntNode{0, nil, nil}
	sentinel.prev = sentinel
	sentinel.next = sentinel
	return Deque{sentinel, 0}
}

// Constructor from new empty deque populated with elements from slice.
func FromSlice(slice []int) Deque {
	deq := New()
	for _, val := range slice {
		deq.Append(val)
	}
	return deq
}

// Append adds element to the end of the deque.
func (d *Deque) Append(elem int) {
	newNode := IntNode{value: elem, next: d.sentinel, prev: d.sentinel.prev}
	d.sentinel.prev.next = &newNode
	d.sentinel.prev = &newNode
	d.length += 1
}

// AppendLeft adds element to the front of the deque.
func (d *Deque) AppendLeft(elem int) {
	newNode := IntNode{value: elem, next: d.sentinel.next, prev: d.sentinel}
	d.sentinel.next.prev = &newNode
	d.sentinel.next = &newNode
	d.length += 1
}

// Pop removes and returns last element in deque.
func (d *Deque) Pop() int {
	if d.Empty() {
		panic("popping from empty list")
	}
	value := d.sentinel.prev.value
	d.sentinel.prev.prev.next = d.sentinel
	d.sentinel.prev = d.sentinel.prev.prev
	d.length -= 1
	return value
}

// Pop removes and returns first element in deque.
func (d *Deque) PopLeft() int {
	if d.Empty() {
		panic("popping from empty list")
	}
	value := d.sentinel.next.value
	d.sentinel.next.next.prev = d.sentinel
	d.sentinel.next = d.sentinel.next.next
	d.length -= 1
	return value
}

// Last return value of the first element in deque.
func (d *Deque) First() int {
	return d.sentinel.next.value
}

// Last return value of the last element in deque.
func (d *Deque) Last() int {
	return d.sentinel.prev.value
}

// Length returns number of elements in deque.
func (d *Deque) Length() int {
	return d.length
}

// Empty checks if there are any elements in deque.
func (d *Deque) Empty() bool {
	return d.length == 0
}

// String outputs string representation of deque.
func (d *Deque) String() string {
	var b strings.Builder
	curr := d.sentinel.next
	b.WriteString("Deque{")
	for i := 0; i < d.Length(); i++ {
		b.WriteString(strconv.Itoa(curr.value))
		b.WriteString(",")
		curr = curr.next
	}
	b.WriteString("}")
	return b.String()
}

// Equals checks if both deques contain same elements in the same order.
func (d *Deque) Equals(other *Deque) bool {
	if d.Length() != other.Length() {
		return false
	}
	curr := d.sentinel.next
	otherCurr := other.sentinel.next
	for i := 0; i < d.Length(); i++ {
		if curr.value != otherCurr.value {
			return false
		}
		curr = curr.next
		otherCurr = otherCurr.next
	}
	return true
}
