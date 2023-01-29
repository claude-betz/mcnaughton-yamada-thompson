package main

import (
	"testing"
)

/*
	[0]-ε->[1]
	[0]-ε->[2]
	[1]-b->[3]
	[1]-a->[3]
	[2]-c->[3]
*/
var (
	end = &nfa{
		true,
		map[rune][]*nfa{},
	}

	nfa3 = &nfa{
		false,
		map[rune][]*nfa{
			'a': []*nfa{
				end,
			},
			'b': []*nfa{
				end,
			},
		},
	}
	
	nfa2 = &nfa{
		false,
		map[rune][]*nfa{
			'c': []*nfa{
				end,
			},
		},
	}

	nfa1 = &nfa{
		false,
		map[rune][]*nfa{
			'a': []*nfa{
				nfa3,
			},
		},
	}

	start = &nfa{
		false,
		map[rune][]*nfa{
			eps: []*nfa{
				nfa1,
				nfa2,
			},
		},
	}
)

var testCases2 = []struct {
	nfas []*nfa
	expectedEpsClosure []*nfa
	expectedMoves []*nfa
}{
	{
		[]*nfa{
			start,
		},
		[]*nfa{
			start,
			nfa1,
			nfa2,
		},
		[]*nfa{
			nfa1,
			nfa2,
		},
	},
}

func TestEpsClosure(t *testing.T) {
	for _, tc := range testCases2 {
		epsClosure := epsilonClosure(tc.nfas)		
		expectedEpsClosure := tc.expectedEpsClosure

		equal := checkNfaEquality(epsClosure, expectedEpsClosure)
		if !equal {
			t.Errorf("epsClosure: %v\n, expectedEpsClosure: %v\n", epsClosure, expectedEpsClosure)
		}
	}
}

func TestMove(t *testing.T) {
	for _, tc := range testCases2 {
		moves := Move(tc.nfas, eps)
		expectedMoves := tc.expectedMoves

		equal := checkNfaEquality(moves, expectedMoves)
		if !equal {
			t.Errorf("moves: %v\n, expectedmoves: %v\n", moves, expectedMoves)
		}
	}
}

var simTestCases = []struct {
	nfas []*nfa
	input [][]string
	expected [][]bool
}{
	{
		[]*nfa{
			start,
		},
		[][]string{
			[]string{
				"a",
				"b",
				"c",
				"aa",
				"ab",
			},
		},
		[][]bool{
			[]bool {
				false,
				false,
				true,
				true,
				true,
			},
		},
	},
}

func TestSimulation(t *testing.T) {

	var res []bool
	for _, tc := range simTestCases {
		for i, nfa := range tc.nfas {
			for _, val := range tc.input[i] {
				accepts := nfa.Simulate(val)
				res = append(res, accepts)
			}

			// compare
			equal := checkBoolEquality(res, tc.expected[i])
			if !equal{
				t.Errorf("got: %v, expected: %v", res, tc.expected[i])
			}
		}
	}	
}

func checkBoolEquality(res []bool, expected []bool) bool {
	if len(res) != len(expected) {
		return false
	}

	for i, _ := range res {
		if res[i] != expected[i] {
			return false
		}
	}
	return true
}

func checkNfaEquality(res []*nfa, expected []*nfa) bool {
	if len(res) != len(expected) {
		return false
	}

	for i, _ := range res {
		if res[i] != expected[i] {
			return false
		}
	}
	return true
}
	
