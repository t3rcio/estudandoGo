package main

import (
	"fmt"
	"os"
)

func main() {
	termo := leTermo()
	fmt.Println("O termo digitado foi: ", termo)
	os.Exit(0)
}

func leTermo() string {
	var termo string
	fmt.Println("Digite o termo")
	fmt.Scan(&termo)
	return termo
}
