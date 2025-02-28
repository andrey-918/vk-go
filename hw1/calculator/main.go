package main

import (
	"calculator/calc"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Пожалуйста, введите выражение.")
		return
	}

	expression := os.Args[1]
	expAnswer, err := calc.CalcExpr(expression)
	if err == nil {
		fmt.Println(expAnswer)
	} else {
		fmt.Println(err)
	}
}
