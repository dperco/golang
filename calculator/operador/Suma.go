package operador

import (
	"fmt"
	"strconv"
	"strings"
)

func Suma(operacion string) int {
	valores := strings.Split(operacion, "+")
	result := 0

	for i := 0; i < len(valores); i++ {
		num, error := strconv.Atoi(valores[i])
		if error != nil {
			fmt.Println(error)
			fmt.Println("ingresar numero entero")
		}
		result += num
	}

	return result
}
