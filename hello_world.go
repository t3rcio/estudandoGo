// Executaveis comecam com um "package main"
package main

import "fmt"

func main() {
	// Todo package main tem uma funcao main que recebe nenhum parametro e retorna nada
	var texto string = "Programa Go na versao: "
	var nome string = "Hello World"
	var versao float32 = 1.1
	fmt.Println(nome, texto, versao)
}

// 🖥 - para compilar
// go build hello.go

// 🖥 - para executar em memoria (sem gerar o executavel)
// go run hello.go
