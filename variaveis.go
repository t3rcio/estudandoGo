package main

import "fmt"

func main() {
	fmt.Println(" ----- Diferentes modos de declarar uma variavel --- ")
	var nome string = "Usuario da Silva Sauro"
	fmt.Println("Declaracao longa: 'var <nome> <tipo> = <valor>' ")
	fmt.Println("Declaracao curta: '<nome> := <valor>' ")
	fmt.Println(nome)
}
