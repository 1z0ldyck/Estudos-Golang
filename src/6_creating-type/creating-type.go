package main

import (
	"fmt"
)

type hotdog int

var b hotdog
var c int = 123

func main() {

	fmt.Printf("%T", b)

	// b = c - não é possível fazer a inicialização utilizando o endereço de c, pois não são do mesmo tipo.
}
