package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dperco/calculator/operador"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("La Calculadora")
	scanner.Scan()
	operation := scanner.Text()

	result := 0

	if strings.Contains(operation, "+") {
		result = operador.Suma(operation)
	} else if strings.Contains(operation, "-") {
		result = operador.Suma(operation)
	} else if strings.Contains(operation, "*") {
		result = operador.Suma(operation)
	} else if strings.Contains(operation, "/") {
		result = operador.Suma(operation)
	} else {
		fmt.Println("el operador es incorrecto")
	}
	fmt.Println(result)
}
