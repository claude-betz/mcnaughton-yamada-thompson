package main

import (
	"testing"
	"github.com/claude-betz/mcnaughton-yamada-thompson/parser"
)

var r2nTestCases= []struct {
	regex string
	values []string
	expected []bool
}{
	{
		"a",
		[]string{"a", "aa", "aab", "c"},
		[]bool{true, false, false, false},
	},
	{
		"a|b",
		[]string{"a", "b", "c", "aa"},
		[]bool{true, true, false, false},
	},
	{
		"c*",
		[]string{"c", "ccc", "ac", "cccb"},
		[]bool{true, true, false, false},
	},
	{
		"(a|b)*",
		[]string{"aa", "bb", "ab", "c"},
		[]bool{true, true, true, false},
	},
}

func TestRegexToNFA(t *testing.T) {
	for _, tc := range r2nTestCases {
		nfa, _ := parser.Parse(&tc.regex)
		results := make([]bool, 0)
		for _, val := range tc.values {
			res := nfa.Simulate(val)
			results = append(results, res) 
		}

		eq := checkBoolEquality(tc.expected, results)
		
		if !eq {
			t.Fail()
		}
	}
}

func checkBoolEquality(arr1, arr2 []bool) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	
	for i, val := range arr1{
		if val != arr2[i] {
			return false
		}
	}
	return true
}
