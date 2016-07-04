package postfix

// Convert converts the given infix expression to postfix expression
func Convert(expression string) ([]rune, map[rune]bool) {

	// queue, stack for handling operands, operator respectively
	var queue []rune
	var stack [100]rune

	// initialisng the queue, stack indexes
	var qi, si int = -1, -1

	// variables holds list of variables
	// that we find in the given infix expression
	variables := make(map[rune]bool)

	// loop through UTF-8 string to find runes(~character) for building AST
	for _, e := range expression {
		switch {
		case isSpace(e):
			continue
		case isNot(e) || isAnd(e) || isOr(e):
			for (si != -1) && (stack[si] == '!' || stack[si] == '&' || stack[si] == '|') {
				if precedence(e) <= precedence(stack[si]) {
					s := stack[si]
					stack[si] = 0
					si--

					queue = append(queue, s)
					qi++
				}
			}
			si++
			stack[si] = e
		case isOpenParentheses(e):
			si++
			stack[si] = e
		case isCloseParentheses(e):
			for stack[si] != '(' {
				s := stack[si]
				stack[si] = 0
				si--

				queue = append(queue, s)
				qi++
			}
			stack[si] = 0
			si--
		case isAlphabet(e):
			queue = append(queue, e)
			qi++

			if !variables[e] {
				variables[e] = true
			}
		}
	}

	for si != -1 {
		s := stack[si]
		stack[si] = 0
		si--

		queue = append(queue, s)
		qi++
	}

	return queue, variables
}

// isNot finds wheather given value is logical `NOT` or not
func isNot(r rune) bool {
	if r == '!' {
		return true
	}
	return false
}

// isAnd finds wheather given value is logical `AND` or not
func isAnd(r rune) bool {
	if r == '&' {
		return true
	}
	return false
}

// isOr finds wheather given value is logical `OR` or not
func isOr(r rune) bool {
	if r == '|' {
		return true
	}
	return false
}

// isAlphabet finds wheather given value is alphabet or not
func isAlphabet(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	return false
}

// isSpace finds wheather given value is space or not
func isSpace(r rune) bool {
	if r == ' ' {
		return true
	}
	return false
}

// isOpenParenthesis finds wheather given value is left parenthesis or not
func isOpenParentheses(r rune) bool {
	if r == '(' {
		return true
	}
	return false
}

// isCloseParenthesis finds wheather given value is right parenthesis or not
func isCloseParentheses(r rune) bool {
	if r == ')' {
		return true
	}
	return false
}

// precedence is a local function that helps dealing with operator precedence
func precedence(r rune) int {
	var p int

	switch r {
	case '!':
		p = 3
	case '&':
		p = 2
	case '|':
		p = 1
	}

	return p
}
