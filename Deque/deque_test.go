package deque

import "testing"

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
