package nfa

import (
	"io"
	"fmt"
	"bytes"
	"strings"
)

const (
	Eps = 'ε'
)

type Nfa struct {
	Accepting bool
	Edges map[rune][]*Nfa
}

func GetEndState(n *Nfa) *Nfa {
	for _, nextList := range n.Edges {
		for _, elem := range nextList {
			if elem.Accepting {
				return elem
			}
			return GetEndState(elem)
		}
		
	}	
	return nil
}

func EpsilonClosure(T []*Nfa) []*Nfa {
	// initialise
	var EpsClosure []*Nfa

	// stack
	stack := make([]*Nfa, 0)

	// push all initial states to EpsClosure and stack 
	for _, Nfa := range T {
		EpsClosure = append(EpsClosure, Nfa)
		stack = append(stack, Nfa) 
	}

	// while stack not empty
	for {
		if len(stack) == 0 {
			break
		}

		// deque last item
		t := stack[len(stack)-1]
		// pop
		stack = stack[:len(stack)-1]

		// iterate all states reachable via Eps
		for _, Nfa := range t.Edges[Eps] {
			EpsClosure = append(EpsClosure, Nfa)
			stack = append(stack, Nfa)
		}
	}

	return EpsClosure 
}

func Move(T []*Nfa, c rune) []*Nfa {
	var res []*Nfa

	for _, Nfa := range T {
		val, ok := Nfa.Edges[c]
		if ok {
			res = append(res, val...)
		}
	}

	return res
}

func (n *Nfa) Simulate(input string) bool {
	buf := bytes.NewBufferString(input)

	S := EpsilonClosure([]*Nfa{n})
	c, _, err := buf.ReadRune()

	for {
		if err == io.EOF {
			break
		}

		S = EpsilonClosure(Move(S, c))
		c, _, err = buf.ReadRune()
	}

	for _, s := range S {
		if s.Accepting {
			return true
		}
	}
	return false
}

func (n *Nfa) PrintNFA() {
	// need to track assigned state numbers
	var seen = make(map[*Nfa]int)
	var levelMap = make(map[*Nfa]int)	

	// queue for bfs
	var queue []*Nfa

	// nextState
	var stateId = 0

	// level
	var level = 1

	// populate level 0  
	for key, nextStates := range n.Edges {	
		for _, nextState := range nextStates {
			// increment stateId
			stateId++

			// print state
			fmt.Printf("[%d]--%s->[%d]\n", 0, string(key), stateId)

			// add to seen map
			seen[nextState] = stateId

			// add to level map
			levelMap[nextState] = level

			// add to queue
			queue = append(queue, nextState)
		}
	}

	for {
		if len(queue) == 0 {
			break
		}

		// deque
		curr := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		// currStateId
		currStateId := seen[curr]

		// currLevel
		currLevel := levelMap[curr]
		
		for char, nextStates := range curr.Edges {			
			for _, nextState := range nextStates {
				val, ok := seen[nextState]

				// we have not seen this state before
				if !ok {
					// increment state
					stateId++

					// add to val
					val = stateId

					// add to seen
					seen[nextState] = stateId	

					// add to level
					levelMap[nextState] = currLevel + 1

					// add to queue
					queue = append(queue, nextState)
				}


				indent := strings.Repeat("\t", currLevel)
				// print
				fmt.Printf("%s[%d]--%s->[%d]\n", indent, currStateId, string(char), val)
			}
		}		 
	}
}

