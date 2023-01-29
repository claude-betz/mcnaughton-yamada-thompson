package parser

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/claude-betz/mcnaughton-yamada-thompson/ast"
	"github.com/claude-betz/mcnaughton-yamada-thompson/nfa"
)

type eof struct {}

func (e *eof) Error() string {
	return "eof"
}

func matchCharacter(c string, opts ...string) (bool, error) {
	s := strings.Join(opts, "")
	if strings.Index(s, c) >= 0 {
		return true, nil
	}
	return false, fmt.Errorf("unknown token %s", c)
}

func Parse(s *string) (*nfa.Nfa, error) {
	return union(s)
}

func union(s *string) (*nfa.Nfa, error) {
	n1, err := concat(s)
	if err != nil {
		return nil, err
	}

	n2, err := unionTail(s)
	if err != nil {
		return n1, nil
	}

	return ast.BuildUnion(n1, n2), nil
}

func unionTail(s *string) (*nfa.Nfa, error) {
	next, err := peekNext(s)
	if err != nil {
		return nil, err
	}
	
	_, err = matchCharacter(next, "|")
	if err != nil {	
		return nil, fmt.Errorf("not union")
	}

	_, err = readNext(s)
	if err != nil {
		return nil, err
	}

	n1, err := concat(s)
	if err != nil {
		return nil, err
	}
	
	n2, err := unionTail(s)
	if err != nil {
		return n1, nil
	}

	return ast.BuildUnion(n1, n2), nil
}

func concat(s *string) (*nfa.Nfa, error) {
	n1, err := closure(s)
	if err != nil {
		return nil, err
	}

	n2, err := concatTail(s)
	if err != nil {
		return n1, nil
	}

	return ast.BuildConcat(n1, n2), nil
}

func concatTail(s *string) (*nfa.Nfa, error) {
	next, err := peekNext(s)
	if err != nil {
		return nil, err
	}
	
	if !unicode.IsLetter(rune(next[0])) {
		return nil, fmt.Errorf("not concat") 
	}

	n1, err := closure(s)
	if err != nil {
		return nil, err
	}

	n2, err := concatTail(s)
	if err != nil {
		return n1, nil
	}

	return ast.BuildConcat(n1, n2), nil
}

func closure(s *string) (*nfa.Nfa, error) {
	n1, err := value(s)
	if err != nil {
		return nil, err
	}
	
	next, err := peekNext(s)
	if err != nil {
		_, ok := err.(*eof)
		if ok {
			return n1, nil
		}
	}

	_, err = matchCharacter(next, "*")
	if err != nil {	
		return n1, nil
	}

	return ast.BuildClosure(n1), err
}

func value(s *string) (*nfa.Nfa, error) {
	next, err := readNext(s)
	if err != nil {
		err, ok := err.(*eof)
		if ok {
			return nil, err
		}
		return nil, err
	}
	
	if unicode.IsLetter(rune(next[0])) {
		// construct value NFA
		return ast.BuildBaseCase(rune(next[0])), nil 
	}

	_, err = matchCharacter(next, "(")
	if err != nil {
		return nil, err
	}	
	
	n, err := union(s)
	if err != nil {
		return nil, err 
	}

	next, err = readNext(s)
	if err != nil {
		return nil, err
	}
	
	_, err = matchCharacter(next, ")")
	if err != nil {
		return nil, err
	}

	return n, nil
}

func peekNext(s *string) (string, error) {
	if len(*s) > 0 {
		return string((*s)[0]), nil
	}
	return "", &eof{}
}

func readNext(s *string) (string, error) {
	if len(*s) > 0 {
		next := string((*s)[0])
		*s = (*s)[1:]
		return next, nil
	}
	return "", &eof{} 
}
