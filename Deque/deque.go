package deque

type Deque struct {
	sentinel *IntNode
	length   int
}

// New initializes new deque.
func New() Deque {
	sentinel := IntNode{0, nil, nil}
	sentinel.prev = &sentinel
	sentinel.next = &sentinel
	return Deque{&sentinel, 0}
}

// FromSlice initilizes new deque and populates it with elements from slice.
func FromSlice(slice []int) Deque {
	deq := New()
	for _, val := range slice {
		deq.Append(val)
	}
	return deq
}

// Append add element to the end of the list.
func (d *Deque) Append(elem int) {
	newNode := IntNode{value: elem, next: d.sentinel, prev: d.sentinel.prev}
	d.sentinel.prev.next = &newNode
	d.sentinel.prev = &newNode
	d.length += 1
}

func (d *Deque) AppendLeft(elem int) {

	d.length += 1
}

func (d *Deque) Pop() int {
	d.length -= 1
	return 0
}

func (d *Deque) PopLeft() int {
	d.length -= 1
	return 0
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

func (d *Deque) String() {

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
