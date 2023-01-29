package ast

import (
	"testing"
)

var (
	test1 = buildBaseCase('a')
	
	n1 = buildBaseCase('a')
	n2 = buildBaseCase('b')
	test2 = buildUnion(n1, n2) 
	
	n3 = buildBaseCase('a')
	n4 = buildBaseCase('b')
	n5 = buildUnion(n3, n4) 
	n6 = buildBaseCase('c')
	test3 = buildConcat(n5, n6)

	n7 = buildBaseCase('a')
	test4 = buildClosure(n7)
)

var testCases = []struct{
	nfa *nfa
	inputs []string
	outputs []bool
}{
	{
		test1,
		[]string{"a", "b", "z", "ab"},
		[]bool{ true, false, false, false},
	},
	{
		test2,
		[]string{"a", "b", "c", "aa"},
		[]bool{true, true, false, false},
	},
	{
		test3,
		[]string{"ac", "bc", "c", "aa", "a"},
		[]bool{true, true, false, false, false},
	},
	{
		test4,
		[]string{"a", "aa", "aaa", "ab"},
		[]bool{true, true, true, false},
	},
}

func TestBuildingNFA(t *testing.T) {
	for _, tc := range testCases {
		var resArr[] bool
		for _, input := range tc.inputs {
			res := tc.nfa.Simulate(input)
			resArr = append(resArr, res)
		}

		equal := checkEquality(resArr, tc.outputs)
		if !equal {
			tc.nfa.PrintNFA()
			t.Errorf("expected: %v, got:%v", tc.outputs, resArr)
		}
	}
}

func checkEquality(arr1, arr2 []bool) bool {
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

