package main

import (
	RN "Calculator/RomanNumerals"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func decodeNumber(Str string) (int, bool, error) {
	isRoman := false
	result, err := strconv.Atoi(Str)
	if err != nil {
		result, err = RN.Decode(Str)
		if err != nil {
			err = fmt.Errorf("Expression is invalid: unable to decode number %q.", Str)
			return 0, false, err
		} else {
			isRoman = true
		}
	}

	if result <= 0 || result > 10 {
		return 0, false, fmt.Errorf("Operand %q is out of bounds [1, 10].", Str)
	}

	return result, isRoman, err
}

func execExpression(input string) (string, error) {
	// parse input
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		return "", fmt.Errorf("Expression %q is invalid: 3 space separated parts required.", input)
	}

	// decode elements
	left, leftRoman, err := decodeNumber(parts[0])
	if err != nil {
		return "", err
	}

	right, rightRoman, err := decodeNumber(parts[2])
	if err != nil {
		return "", err
	}

	if leftRoman != rightRoman {
		return "", fmt.Errorf("Expression %q is invalid: both parts must use same notation (arabic or roman).", input)
	}

	// execute expression
	var result int
	switch parts[1] {
	case "+":
		result = left + right
	case "-":
		result = left - right
	case "*":
		result = left * right
	case "/":
		result = left / right
	default:
		return "", fmt.Errorf("Invalid operator used (%q), only '+', '-', '*' and '/' are allowed", parts[1])
	}

	var resultRepr string
	if leftRoman {
		resultRepr, err = RN.Encode(result)
		if err != nil {
			return "", fmt.Errorf("%s\n%s", fmt.Sprintf("Error while encoding expression result (%d):", result), err)
		}
	} else {
		resultRepr = strconv.Itoa(result)
	}

	// output result
	return resultRepr, nil
}

func main() {
	fmt.Println("\nCalculator\n")
	fmt.Println("Expression consist of left value, operator and right value. Roman and arabic number notation supported.")
	fmt.Println("Both numbers should be integers in range [1, 10], elements should be separated by space symbol.")
	fmt.Println("Supported operators: '+', '-', '*', '/'. Example: 10 * 9")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nEnter expression (type 'exit' to stop the program):")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			return
		}

		result, err := execExpression(input)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println(result)
		}
	}
}
