package ast

import (
	"github.com/claude-betz/mcnaughton-yamada-thompson/nfa"
)

func BuildBaseCase(char rune) *nfa.Nfa {
	startState := &nfa.Nfa{
		Accepting: false,
		Edges: make(map[rune][]*nfa.Nfa),
	}

	endState := &nfa.Nfa{
		Accepting: true,
		Edges: make(map[rune][]*nfa.Nfa),
	}
	
	startState.Edges[char] = []*nfa.Nfa{
		endState,
	}

	return startState 
}

func BuildClosure(n *nfa.Nfa) *nfa.Nfa {
	startState := &nfa.Nfa{
		Accepting: false,
		Edges: make(map[rune][]*nfa.Nfa),
	}

	endState := &nfa.Nfa{
		Accepting: true,
		Edges: make(map[rune][]*nfa.Nfa),
	}

	// add nfa.Epsilon transition from start state of new NFA
	// 1. to start state of N(s) 
	// 2. to end state of new NFA 
	startState.Edges[nfa.Eps] = []*nfa.Nfa{
		n,
		endState,
	}
		
	nfaEndState := nfa.GetEndState(n)
	// set end state as not final
	nfaEndState.Accepting = false

	// add nfa.Epsilon transition from end state of N(s):
	// 1. to start state of N(s)	
	// 2. to end state of new NFA	
	endStateArr := []*nfa.Nfa{
		n,
		endState,
	}
	nfaEndState.Edges[nfa.Eps] = endStateArr

	return startState	
}

func BuildConcat(n1, n2 *nfa.Nfa) *nfa.Nfa {
	// merge end state of N(s) and start state of N(t)
	nfa1EndState := nfa.GetEndState(n1)
	nfa1EndState.Accepting = false	
	nfa1EndState.Edges[nfa.Eps] = []*nfa.Nfa{
		n2,
	}

	return n1
}

func BuildUnion(n1, n2 *nfa.Nfa) *nfa.Nfa {
	startState := &nfa.Nfa{
		Accepting: false,
		Edges: make(map[rune][]*nfa.Nfa),
	}

	endState := &nfa.Nfa{
		Accepting: true,
		Edges: make(map[rune][]*nfa.Nfa),
	}
	
	// add nfa.Epsilon transition from start state of new NFA
	// 1. to start state of N(s)
	// 2. to start state of N(t)
	startState.Edges[nfa.Eps] = []*nfa.Nfa{
		n1,
		n2,
	}

	// add nfa.Epsilon transition from end state of
	// 1. N(s) to end state of new NFA
	// 2. N(t) to end state of new NFA
	nfa1EndState := nfa.GetEndState(n1)
	nfa2EndState := nfa.GetEndState(n2)

	// set end states to false
	nfa1EndState.Accepting = false
	nfa2EndState.Accepting = false

	nfa1EndState.Edges[nfa.Eps] = []*nfa.Nfa{
		endState,
	}
	nfa2EndState.Edges[nfa.Eps] = []*nfa.Nfa{
		endState,
	}
	
	return startState
}

