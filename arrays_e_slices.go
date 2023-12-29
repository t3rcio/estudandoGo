/*
  - Diferenças entre Array e Slice
  - O slice é uma abstração que funciona acima do array.
  - Ao criarmos um array, o item é criado na memória com as suas posições e com o tipo especificado.
  - Quando atribuímos valores aos seus índices, eles vão sendo preenchidos,
    e os índices que não utilizarmos vão continuar em branco, esperando serem preenchidos.
  - Já no caso do slice, ele cria um array de acordo com os elementos que passamos para ele.
*/
package main

import "fmt"

func main() {
	var arrayString [4]string
	var meuSlice []string

	arrayString[0] = "Elemento 0"
	arrayString[1] = "Elemento 1"
	arrayString[2] = "Elemento 2"
	arrayString[3] = "Elemento 3"

	fmt.Println("Array criado: ", arrayString)

	//❗- Um erro é lançado aqui: o slice foi inicializado com tamanho 0
	// meuSlice[0] = "Elemento 0"

	// ✅ - Adicionando elemento ao slice
	// Para adicionar um elemento usa-se a funcao append(slice, elemento)
	// A funcao append retorna um novo slice com o elemento adiconado
	meuSlice = append(meuSlice, "Elemento 0")

	fmt.Println("Slice criado: ", meuSlice, "|Length: ", len(meuSlice))

	// 💟 - Outra forma de inicializar um slice é
	meuOutroSlice := []string{"Elemento 0", "Elemento 1", "Elemento 2"}
	fmt.Println("meuOutroSlice", meuOutroSlice)

}
