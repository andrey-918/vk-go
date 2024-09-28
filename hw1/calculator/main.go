package main

import (
	"bufio"
	"calculator/calc"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите строку: ")
	for scanner.Scan() {

		expression := scanner.Text()
		expAnswer, err := calc.CalcExpr(expression)
		if err == nil {
			fmt.Println(expAnswer)
		} else {
			fmt.Println(err)
		}
		fmt.Print("Введите строку: ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении:", err)
	}

}
