package postfix

// Evaluate evalutes the given postfix expression
// for the given values
func Evaluate(exp []rune, evar map[rune]bool, eval []rune) bool {

	// stack for reducing given postfix expression
	var stack [100]rune
	var si int = -1

	// Copying the postfix to temporary variable so that
	// original value in the `exp` does not get modified
	// and will be available for next iteration
	var expression []rune
	for _, e := range exp {
		expression = append(expression, e)
	}

	index := 0
	//mapping data and expression
	for v, found := range evar {
		for i, e := range expression {
			if found && (v == e) {
				expression[i] = eval[index]
			}
		}
		index++
	}

	for _, e := range expression {
		switch {
		case isNot(e):
			s := stack[si]
			stack[si] = 0
			si--

			result := eval_not(s)
			si++
			stack[si] = result
		case isAnd(e):
			s1 := stack[si]
			stack[si] = 0
			si--

			s2 := stack[si]
			stack[si] = 0
			si--

			result := eval_and(s1, s2)
			si++
			stack[si] = result
		case isOr(e):
			s1 := stack[si]
			stack[si] = 0
			si--

			s2 := stack[si]
			stack[si] = 0
			si--

			result := eval_or(s1, s2)
			si++
			stack[si] = result
		case isData(e):
			si++
			stack[si] = e
		}
	}

	var result rune
	if si == 0 {
		result = stack[si]
		stack[si] = 0
		si--
	}

	if result == '0' {
		return false
	}

	return true
}

// eval_not is uesd to evaluate logical NOT operator
func eval_not(r rune) rune {
	if r == '0' {
		return '1'
	}
	return '0'
}

// eval_and is uesd to evaluate logical AND operator
func eval_and(r1, r2 rune) rune {
	if r1 == '1' && r2 == '1' {
		return '1'
	}
	return '0'
}

// eval_or is uesd to evaluate logical OR operator
func eval_or(r1, r2 rune) rune {
	if r1 == '1' || r2 == '1' {
		return '1'
	}
	return '0'
}

// isData finds wheather given value is data or not
func isData(r rune) bool {
	return r == '0' || r == '1'
}
