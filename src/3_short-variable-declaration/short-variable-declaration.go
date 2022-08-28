package main

import (
	"fmt"
)

func main() {

	// Utilizado para criar variáveis com um nome próprio e um valor inicial. O objetivo principal desse tipo de usar este operador é declarar e inicializar variávelis locais dentro das funções e estreitar o escopo das variáveis(codeblocks).
	x := 10
	// x := 20        // Não é possível atribuir valores para as variáveis dessa forma
	x, y := 30, 40 // Porém é possível atribuir dessa forma, ou seja, é necessário ao mesmo tempo que atribui um valor, declarar uma nova variável.

	fmt.Println(x, y)
}
