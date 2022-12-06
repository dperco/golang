package mensage

import "fmt"

func functionPrivete() {
	fmt.Println("soy funcion private")

}

func FunctionPublic() {
	fmt.Println("soy function public")

	functionPrivete()
}
