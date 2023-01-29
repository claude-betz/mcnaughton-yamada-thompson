package nfa

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
	end = &Nfa{
		true,
		map[rune][]*Nfa{},
	}

	Nfa3 = &Nfa{
		false,
		map[rune][]*Nfa{
			'a': []*Nfa{
				end,
			},
			'b': []*Nfa{
				end,
			},
		},
	}
	
	Nfa2 = &Nfa{
		false,
		map[rune][]*Nfa{
			'c': []*Nfa{
				end,
			},
		},
	}

	Nfa1 = &Nfa{
		false,
		map[rune][]*Nfa{
			'a': []*Nfa{
				Nfa3,
			},
		},
	}

	start = &Nfa{
		false,
		map[rune][]*Nfa{
			Eps: []*Nfa{
				Nfa1,
				Nfa2,
			},
		},
	}
)

var testCases2 = []struct {
	Nfas []*Nfa
	expectedEpsClosure []*Nfa
	expectedMoves []*Nfa
}{
	{
		[]*Nfa{
			start,
		},
		[]*Nfa{
			start,
			Nfa1,
			Nfa2,
		},
		[]*Nfa{
			Nfa1,
			Nfa2,
		},
	},
}

func TestEpsClosure(t *testing.T) {
	for _, tc := range testCases2 {
		EpsClosure := EpsilonClosure(tc.Nfas)		
		expectedEpsClosure := tc.expectedEpsClosure

		equal := checkNfaEquality(EpsClosure, expectedEpsClosure)
		if !equal {
			t.Errorf("EpsClosure: %v\n, expectedEpsClosure: %v\n", EpsClosure, expectedEpsClosure)
		}
	}
}

func TestMove(t *testing.T) {
	for _, tc := range testCases2 {
		moves := Move(tc.Nfas, Eps)
		expectedMoves := tc.expectedMoves

		equal := checkNfaEquality(moves, expectedMoves)
		if !equal {
			t.Errorf("moves: %v\n, expectedmoves: %v\n", moves, expectedMoves)
		}
	}
}

var simTestCases = []struct {
	Nfas []*Nfa
	input [][]string
	expected [][]bool
}{
	{
		[]*Nfa{
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
		for i, Nfa := range tc.Nfas {
			for _, val := range tc.input[i] {
				accepts := Nfa.Simulate(val)
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

func checkNfaEquality(res []*Nfa, expected []*Nfa) bool {
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
	
