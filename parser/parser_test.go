package parser

import (
	"testing"
)

var readNextTestCases = []struct{
	regex string
	expectedVal string
	expectedPeek string
	expectedNext string	
}{
	{
		"abc",
		"a",
		"abc",
		"bc",
	},
}

func TestReadNext(t *testing.T) {
	for _, tc := range readNextTestCases {
		pointer := &tc.regex
		val, _ := readNext(pointer)
		if val != tc.expectedVal {
			t.Fail()
		}
		if *pointer != tc.expectedNext {
			t.Fail()
		}
	}
}

func TestPeekNext(t *testing.T) {
	for _, tc := range readNextTestCases {
		pointer := &tc.regex
		val, _ := peekNext(pointer)
		if val != tc.expectedVal { 
			t.Fail()
		}
		if *pointer != tc.expectedPeek {
			t.Fail()
		}
	}
}
