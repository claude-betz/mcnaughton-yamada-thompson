package ast

func buildBaseCase(char rune) *nfa {
	startState := &nfa{
		accepting: false,
		edges: make(map[rune][]*nfa),
	}

	endState := &nfa{
		accepting: true,
		edges: make(map[rune][]*nfa),
	}
	
	startState.edges[char] = []*nfa{
		endState,
	}

	return startState 
}

func buildClosure(n *nfa) *nfa {
	startState := &nfa{
		accepting: false,
		edges: make(map[rune][]*nfa),
	}

	endState := &nfa{
		accepting: true,
		edges: make(map[rune][]*nfa),
	}

	// add epsilon transition from start state of new NFA
	// 1. to start state of N(s) 
	// 2. to end state of new NFA 
	startState.edges[eps] = []*nfa{
		n,
		endState,
	}
		
	nfaEndState := GetEndState(n)
	// set end state as not final
	nfaEndState.accepting = false

	// add epsilon transition from end state of N(s):
	// 1. to start state of N(s)	
	// 2. to end state of new NFA	
	endStateArr := []*nfa{
		n,
		endState,
	}
	nfaEndState.edges[eps] = endStateArr

	return startState	
}

func buildConcat(n1, n2 *nfa) *nfa {
	// merge end state of N(s) and start state of N(t)
	nfa1EndState := GetEndState(n1)
	nfa1EndState.accepting = false	
	nfa1EndState.edges[eps] = []*nfa{
		n2,
	}

	return n1
}

func buildUnion(n1, n2 *nfa) *nfa {
	startState := &nfa{
		accepting: false,
		edges: make(map[rune][]*nfa),
	}

	endState := &nfa{
		accepting: true,
		edges: make(map[rune][]*nfa),
	}
	
	// add epsilon transition from start state of new NFA
	// 1. to start state of N(s)
	// 2. to start state of N(t)
	startState.edges[eps] = []*nfa{
		n1,
		n2,
	}

	// add epsilon transition from end state of
	// 1. N(s) to end state of new NFA
	// 2. N(t) to end state of new NFA
	nfa1EndState := GetEndState(n1)
	nfa2EndState := GetEndState(n2)

	// set end states to false
	nfa1EndState.accepting = false
	nfa2EndState.accepting = false

	nfa1EndState.edges[eps] = []*nfa{
		endState,
	}
	nfa2EndState.edges[eps] = []*nfa{
		endState,
	}
	
	return startState
}

