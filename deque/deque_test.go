package deque

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEquals(t *testing.T) {
	for _, tc := range TestCasesEquals {
		if tc.this.Equals(&tc.other) != tc.expected {
			t.Errorf(tc.msg)
		}
	}
}

func TestAppend(t *testing.T) {
	for _, tc := range TestCasesAppend {
		actual := tc.actual()
		if !actual.Equals(&tc.expected) {
			t.Errorf(tc.msg)
		}
	}
}

func TestPop(t *testing.T) {
	// Main test cases
	for _, tc := range TestCasesPop {
		deque, popped := tc.actual()
		if !deque.Equals(&tc.expectedDeque) || popped != tc.expectedPopped {
			t.Errorf(tc.msg)
		}
	}

	// Popping from empty list
	poppingEmpty := func() {
		d := FromSlice([]int{})
		d.Pop()
	}
	assert.Panics(t, poppingEmpty, "popping from empty deque")
}

func TestPopLeft(t *testing.T) {
	// Main test cases
	for _, tc := range TestCasesPopLeft {
		deque, popped := tc.actual()
		if !deque.Equals(&tc.expectedDeque) || popped != tc.expectedPopped {
			t.Errorf(tc.msg)
		}
	}

	// Popping from empty list
	poppingEmpty := func() {
		d := FromSlice([]int{})
		d.Pop()
	}
	assert.Panics(t, poppingEmpty, "popping from empty deque")
}
