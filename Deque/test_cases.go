package deque

type TestCase struct {
	this     Deque
	other    Deque
	expected bool
	msg      string
}

var TestCasesEquals []TestCase = []TestCase{
	{
		this:     New(),
		other:    New(),
		expected: true,
		msg:      "two empty deques",
	},
	{
		this:     FromSlice([]int{1, 2, 3}),
		other:    FromSlice([]int{1, 2, 3}),
		expected: true,
		msg:      "two identical, 3-element deques",
	},
	{
		this:     FromSlice([]int{1, 2, 3}),
		other:    FromSlice([]int{1, 2, 4}),
		expected: false,
		msg:      "same length, last element mismatching",
	},
	{
		this:     FromSlice([]int{1, 2, 3}),
		other:    FromSlice([]int{1, 2}),
		expected: false,
		msg:      "different length, elements matching",
	},
}

type TestCaseAppend struct {
	actual   func() Deque
	expected Deque
	msg      string
}

var TestCasesAppend = []TestCaseAppend{
	{
		actual: func() Deque {
			return New()
		},
		expected: New(),
		msg:      "2 empty deques",
	},
	{
		actual: func() Deque {
			d := New()
			d.Append(1)
			return d
		},
		expected: FromSlice([]int{1}),
		msg:      "adding one element to deque",
	},
	{
		actual: func() Deque {
			d := New()
			d.Append(1)
			d.Append(3)
			d.Append(5)
			d.Append(19)
			return d
		},
		expected: FromSlice([]int{1, 3, 5, 19}),
		msg:      "adding multiple elements to deque",
	},
}
