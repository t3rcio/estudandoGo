package main

import "fmt"

func main() {
	pontosPlanningPoker := []int{1, 2, 3, 5, 8, 13, 21, 40}

	// usando for
	// for i := 0; i < len(pontosPlanningPoker); i++ {
	// 	fmt.Println(pontosPlanningPoker[i])
	// }

	// usando range
	for i, ponto := range pontosPlanningPoker {
		fmt.Println("O ponto na posição", i, " tem o valor", ponto)
	}
}
