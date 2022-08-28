package main

import (
	"fmt"
)

// Utilizado para declarar variáveis com escopo de pacote(package scope), diferente da declaração de variável curta que é necessário declarar um valor, ao usar a keyword var podemos inicializar a variável sem valor;
// use var para package-level scope.

var x = 10 // declaração e inicialização

// declaração, por padrão, ele receberá o valor zero(valores por padrão para variáveis antes delas serem inicializadas), mais informações onsultar explanations/valor-zero.md
var a int
var b string
var c float64
var d bool

func main() {

	fmt.Println(x)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
