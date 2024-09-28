package calc

import (
	"calculator/stack"
	"errors"

	"math"
	"strconv"
	"unicode"
)

func precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	}
	return 0
}

func isValidNumber(num string) bool {
	dotCount := 0
	for index, char := range num {
		switch char {
		case '+', '-':
			if index != 0 {
				return false
			}
		case '.':
			dotCount++
			if dotCount > 1 {
				return false // Больше одной точки
			}
		default:
			if !unicode.IsDigit(char) {
				return false //Некорректный символ
			}
		}
	}
	return num[len(num)-1] != '.'
}

func isOperator(char rune) bool {
	return char == '+' || char == '-' || char == '*' || char == '/'
}

func handlePlusAndMinus(char string, lastWasOperator *bool, currentNum *string, resultStack *stack.Stack, operStack *stack.Stack) error {
	if *lastWasOperator { // Обработка унарного оператора
		if *currentNum == "" {
			*currentNum = string(char)
		} else {
			return errors.New("Invalid input")
		}
		*lastWasOperator = true
		return nil
	}

	topValue, ok := operStack.Pop()
	for ok {
		// Проверка типа перед приведением
		if strValue, ok := topValue.(string); ok {
			if precedence(strValue) >= precedence(char) {
				resultStack.Push(strValue)
			} else {
				// Если приоритет не подходит, возвращаем элемент обратно в стек
				operStack.Push(topValue)
				break
			}
		} else {
			return errors.New("Invalid type in stack: expected string")
		}
		topValue, ok = operStack.Pop()
	}

	operStack.Push(string(char))
	*lastWasOperator = true 
	return nil
}

func handleMultiplicationAndDivision(char string, lastWasOperator *bool, resultStack *stack.Stack, operStack *stack.Stack) error {
	topValue, ok := operStack.Pop()
	for ok {
		if strValue, ok := topValue.(string); ok {
			if precedence(strValue) >= precedence(char) {
				resultStack.Push(strValue)
			} else {
				// Если приоритет не подходит, возвращаем элемент обратно в стек
				operStack.Push(topValue)
				break
			}
		} else {
			return errors.New("Invalid type in stack: expected string")
		}
		topValue, ok = operStack.Pop()
	}

	operStack.Push(string(char))
	*lastWasOperator = true 
	return nil
}

func handleOpenBracket(char string, lastWasOperator *bool, resultStack *stack.Stack, operStack *stack.Stack) error {
	if char != "(" {
		return errors.New("Invalid character: expected '('")
	}
	if !*lastWasOperator { // Если перед скобкой нет оператора, добавляем '*'
		topValue, ok := operStack.Pop()
		for ok {
			if strValue, ok := topValue.(string); ok {
				if precedence(strValue) >= precedence("*") {
					resultStack.Push(strValue)
				} else {
					// Если приоритет не подходит, возвращаем элемент обратно в стек
					operStack.Push(topValue)
					break
				}
			} else {
				return errors.New("Invalid type in stack: expected string")
			}
			topValue, ok = operStack.Pop()
		}
		operStack.Push("*")
	}
	operStack.Push("(")
	*lastWasOperator = true // '(' может быть перед унарным оператором

	return nil
}

func handleCloseBracket(char string, lastWasOperator *bool, resultStack *stack.Stack, operStack *stack.Stack) error {
	if char != ")" {
		return errors.New("Invalid character: expected ')'")
	}

	if operStack.IsEmpty() {
		return errors.New("Mismatched parentheses")
	}

	topValue, ok := operStack.Pop()
	for ok {
		strValue, isString := topValue.(string)
		if !isString {
			return errors.New("Invalid type in stack: expected string")
		}

		if strValue == "(" {
			break // Закрываем скобку
		}

		resultStack.Push(strValue)
		topValue, ok = operStack.Pop()
	}

	if topValue == nil && !ok { // Если стек пуст и не было найдено открывающей скобки
		return errors.New("Mismatched parentheses")
	}

	*lastWasOperator = false 
	return nil
}

func infixToPostfix(expression string) (stack.Stack, error) {
	resultStack := stack.New()
	operStack := stack.New()
	if len(expression) == 0 {
		return *resultStack, nil
	}

	var currentNum string
	lastWasOperator := true // Указывает, был ли последний символ оператором или скобкой
	if (expression[0] == '-' || expression[0] == '+') && (len(expression) > 1 && expression[1] == '(') {
		expression = "0" + expression
	}
	for _, char := range expression {
		if unicode.IsSpace(char) {
			continue // Игнорируем пробелы
		}
		token := string(char)
		if unicode.IsDigit(char) || char == '.' {
			currentNum += token     // Собираем число
			lastWasOperator = false 
		} else {
			if currentNum != "" {
				if !isValidNumber(currentNum) {
					return *resultStack, errors.New("Invalid number: " + currentNum)
				}
				resultStack.Push(currentNum)
				currentNum = ""
			}
			var err error
			switch token {
			case "+", "-":
				err = handlePlusAndMinus(token, &lastWasOperator, &currentNum, resultStack, operStack)
			case "*", "/":
				err = handleMultiplicationAndDivision(token, &lastWasOperator, resultStack, operStack)
			case "(":
				err = handleOpenBracket(token, &lastWasOperator, resultStack, operStack)
			case ")":
				err = handleCloseBracket(token, &lastWasOperator, resultStack, operStack)
			default:
				return *resultStack, errors.New("Invalid character: " + string(char))
			}
			if err != nil {
				return *resultStack, err
			}
		}
	}

	if currentNum != "" {
		if !isValidNumber(currentNum) {
			return *resultStack, errors.New("Invalid number: " + currentNum)
		}
		resultStack.Push(currentNum)
	}

	topValue, ok := operStack.Pop()
	for ok {
		resultStack.Push(topValue.(string))
		topValue, ok = operStack.Pop()
	}
	return *resultStack, nil
}

func calculate(firstNum, secondNum interface{}, token string) (float64, error) {
	var result float64

	fNum, ok1 := firstNum.(float64)
	sNum, ok2 := secondNum.(float64)

	if !ok1 || !ok2 {
		return 0, errors.New("both numbers must be of type float64")
	}

	switch token {
	case "+":
		result = fNum + sNum
	case "-":
		result = fNum - sNum
	case "*":
		result = fNum * sNum
	case "/":
		if sNum == 0 {
			return 0, errors.New("division by zero")
		}
		result = fNum / sNum
	default:
		return 0, errors.New("invalid operator")
	}

	return result, nil
}

func CalcExpr(expression string) (float64, error) {
	postFixExp, err := infixToPostfix(expression)
	if err != nil {
		return 0, err
	}

	NumStack := stack.New()

	for token := range postFixExp.Iterate() {

		char, ok := token.(string)
		if ok && isOperator(rune(char[len(char)-1])) {
			secondNum, ok2 := NumStack.Pop()
			firstNum, ok1 := NumStack.Pop()
			if !ok1 || !ok2 {
				return 0, errors.New("Invalid expression")
			}

			var result float64
			result, err = calculate(firstNum, secondNum, char)
			if err != nil {
				return 0, err
			}
			NumStack.Push(result)
		} else {
			num, err := strconv.ParseFloat(char, 64)
			if err != nil {
				return 0, err
			}
			NumStack.Push(num)
		}
	}
	answer, ok := NumStack.Pop()
	if !NumStack.IsEmpty() || !ok {
		return 0, errors.New("Invalid expression")
	}
	roundTo := 10000
	return math.Ceil(answer.(float64)*float64(roundTo)) / float64(roundTo), nil
}
