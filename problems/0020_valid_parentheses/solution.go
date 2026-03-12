package _020_valid_parentheses

type Stack struct {
	elements []rune
}

func (s *Stack) Push(r rune) {
	s.elements = append(s.elements, r)
}

func (s *Stack) Pop() (rune, bool) {
	if len(s.elements) == 0 {
		return 0, false
	}
	index := len(s.elements) - 1
	element := s.elements[index]
	s.elements = s.elements[:index]

	return element, true
}

func IsValid(s string) bool {
	parentheses := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	runeString := []rune(s)
	runeStack := &Stack{
		make([]rune, 0),
	}
	for i := 0; i < len(runeString); i++ {
		_, ok := parentheses[runeString[i]]
		if !ok {
			openParentheses, ok := runeStack.Pop()
			if !ok {
				return false
			}
			if runeString[i] != parentheses[openParentheses] {
				return false
			}
			continue
		}
		runeStack.Push(runeString[i])
	}

	if len(runeStack.elements) == 0 {
		return true
	}

	return false
}
