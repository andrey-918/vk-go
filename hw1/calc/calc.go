package calc

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"strconv"

	"unicode"
)

func precedence(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	}
	return 0
}

func isValidNumber(num string) bool {
	dotCount := 0
	for _, char := range num {
		if !unicode.IsDigit(char) {
			if char == '.' {
				dotCount++
				if dotCount > 1 {
					return false // больше одной точки
				}
			} else {
				return false // некорректный символ
			}
		}
	}
	return num[len(num)-1] != '.'
}

func infixToPostfix(expression string) ([]string, error) {
	// if !isValidExpression(expression) {
	// 	return nil, errors.New("Invalid expression1")
	// }
	var result []string
	var stack []rune

	var currentNum string

	for _, char := range expression {
		if unicode.IsSpace(char) {
			continue // Игнорируем пробелы
		}

		if unicode.IsDigit(char) || char == '.' {
			currentNum += string(char) // Собираем число
		} else {
			if currentNum != "" {
				if !isValidNumber(currentNum) {
					return nil, errors.New("Invalid number: " + currentNum)
				}
				result = append(result, currentNum)
				currentNum = "" // Сбрасываем текущее число
			}

			switch char {
			case '+', '-', '*', '/':
				for len(stack) > 0 && precedence(stack[len(stack)-1]) >= precedence(char) {
					result = append(result, string(stack[len(stack)-1]))
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, char)
			case '(':
				stack = append(stack, '(')
			case ')':
				for len(stack) > 0 && stack[len(stack)-1] != '(' {
					result = append(result, string(stack[len(stack)-1]))
					stack = stack[:len(stack)-1]
				}
				if len(stack) == 0 {
					return nil, errors.New("Mismatched parentheses")
				}
				stack = stack[:len(stack)-1] // убираем '(' (закрываем скобки)
			default:
				return nil, errors.New("Invalid character: " + string(char))
			}
		}
	}

	if currentNum != "" {
		if !isValidNumber(currentNum) {
			return nil, errors.New("Invalid number: " + currentNum)
		}
		result = append(result, currentNum)
	}

	for len(stack) > 0 {
		result = append(result, string(stack[len(stack)-1]))
		stack = stack[:len(stack)-1]
	}

	return result, nil
}

func isOperator(char rune) bool {
	return char == '+' || char == '-' || char == '*' || char == '/'
}

func calcExpr(expression string) (float64, error) {
	postFixExp, err := infixToPostfix(expression)
	if err != nil {
		return 0, err
	}

	stack := []float64{}

	for _, token := range postFixExp {
		if isOperator(rune(token[0])) {
			if len(stack) <= 1 {
				return 0, errors.New("Invalid expression")
			}
			secondNum := stack[len(stack)-1]
			firstNum := stack[len(stack)-2]
			stack = stack[:len(stack)-2] // Убираем последние два числа

			var result float64
			switch token {
			case "+":
				result = firstNum + secondNum
			case "-":
				result = firstNum - secondNum
			case "*":
				result = firstNum * secondNum
			case "/":
				if secondNum == 0 {
					return 0, errors.New("Division by zero")
				}
				result = firstNum / secondNum
			}
			stack = append(stack, result)
		} else {
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, err
			}
			stack = append(stack, num)
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("Invalid expression")
	}

	return stack[0], nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите строку: ")

	if scanner.Scan() {
		expression := scanner.Text()
		expAnswer, err := calcExpr(expression)
		if err == nil {
			fmt.Println(expAnswer)
		} else {
			fmt.Println(err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении:", err)
	}

}
