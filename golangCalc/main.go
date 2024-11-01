package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Addition(arg1, arg2 float64) float64 {
	return arg1 + arg2
}

func Subtraction(arg1, arg2 float64) float64 {
	return arg1 - arg2
}

func Multiply(arg1, arg2 float64) float64 {
	return arg1 * arg2
}

func Separation(arg1, arg2 float64) float64 {
	return arg1 / arg2
}

func Remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func Solver(output []string) []float64 {
	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	operators := []string{"*", "/", "+", "-"}
	solution := []float64{}
	for i := 0; i < len(output); i++ {

		switch {
		case slices.Contains(digits, output[i]):
			arg, _ := strconv.ParseFloat(output[i], 64)
			solution = append(solution, arg)
		case slices.Contains(operators, output[i]):
			switch {
			case output[i] == "+":
				Arg1 := solution[len(solution)-2]
				Arg2 := solution[len(solution)-1]
				solution = solution[:len(solution)-2]
				solution = append(solution, Addition(Arg1, Arg2))

			case output[i] == "-":
				Arg1 := solution[len(solution)-2]
				Arg2 := solution[len(solution)-1]
				solution = solution[:len(solution)-2]
				solution = append(solution, Subtraction(Arg1, Arg2))

			case output[i] == "*":
				Arg1 := solution[len(solution)-2]
				Arg2 := solution[len(solution)-1]
				solution = solution[:len(solution)-2]
				solution = append(solution, Multiply(Arg1, Arg2))

			case output[i] == "/":
				Arg1 := solution[len(solution)-2]
				Arg2 := solution[len(solution)-1]
				solution = solution[:len(solution)-2]
				solution = append(solution, Separation(Arg1, Arg2))
			}

		}
	}
	return solution
}

func Calc(expression string) (float64, error) {

	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	operators := []string{"*", "/", "+", "-"}
	priority := map[string]int{
		"*": 3,
		"/": 4,
		"+": 2,
		"-": 1,
	}
	output := []string{}
	stack := []string{}

	expressionSlice := strings.Split(expression, "")

	switch {
	case len(expression) == 0:
		return 0, fmt.Errorf("empty expression")
	case slices.Contains(operators, expressionSlice[0]) || slices.Contains(operators, expressionSlice[len(expressionSlice)-1]):
		return 0, fmt.Errorf("syntax error")
	}

	for i := 0; i < len(expression); {
		switch {
		case len(expressionSlice) == 0 && len(stack) != 0:
			for len(stack) != 0 {
				if stack[0] == "(" {
					return 0, fmt.Errorf("missed: '('")
				}
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
		case len(expressionSlice) == 0 && len(stack) == 0:
			solve := Solver(output)[0]
			return solve, nil
		case !slices.Contains(digits, expressionSlice[i]) && !slices.Contains(operators, expressionSlice[i]) && expressionSlice[i] != ")" && expressionSlice[i] != "(":
			return 0, fmt.Errorf("wrong symbol")

		case slices.Contains(digits, expressionSlice[i]):
			output = append(output, expressionSlice[i])
			expressionSlice = expressionSlice[1:]

		case slices.Contains(operators, expressionSlice[i]):
			if len(stack) > 0 {
				if slices.Contains(operators, stack[0]) {
					key1 := priority[expressionSlice[i]]
					key2 := priority[stack[len(stack)-1]]
					for key2 >= key1 {
						output = append(output, stack[(len(stack)-1)])
						stack = stack[:len(stack)-1]
						if len(stack) > 0 && slices.Contains(operators, stack[0]) {
							key2 = priority[stack[(len(stack)-1)]]
						} else {
							key2 = key1 - 1
						}
					}
				}

			}
			if slices.Contains(operators, expressionSlice[i+1]) {
				return 0, fmt.Errorf("syntax error")
			}
			stack = append(stack, expressionSlice[i])
			expressionSlice = expressionSlice[1:]

		case expressionSlice[i] == "(":
			if expressionSlice[i+1] == ")" {
				return 0, fmt.Errorf("syntax error")
			}
			stack = append(stack, "(")
			expressionSlice = expressionSlice[1:]
		case expressionSlice[i] == ")":
			for stack[len(stack)-1] != "(" {
				output = append(output, stack[len(stack)-1])
				if len(stack) == 1 {
					return 0, fmt.Errorf("missed: '('")
				}
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			expressionSlice = expressionSlice[1:]
			if len(stack) > 0 && slices.Contains(operators, stack[len(stack)-1]) {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}

		}

	}
	return 0, fmt.Errorf("something went wrong")
}

func main() {
	fmt.Println(Calc("1+1"))
}
