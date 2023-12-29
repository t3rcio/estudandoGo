package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const ARQUIVO = "./example.txt"

func convertByteToString(sliceBytes []byte) string {
	/*
	* Converte um slice de bytes para uma string
	 */
	var result []string
	for _, value := range sliceBytes {
		result = append(result, string(value))
	}
	return strings.Join(result, "")
}

func main() {

	// ```os.Open(caminho_do_arquivo)```` retorna dois valores
	// 1 - o conteudo do arquivo
	// 2 - erros que ocorreram ao tentar abrir o arquivo
	arquivo, err := os.Open(ARQUIVO)
	fmt.Println("Arquivo: ", arquivo)
	fmt.Println("Erros: ", err)

	var pronto string

	fmt.Println("Crie um arquivo example.txt no diretorio atual. Digite  \"pronto\" quando terminar")
	for {
		fmt.Scan(&pronto)
		if pronto != "" && pronto == "pronto" {
			break
		}
		fmt.Println("Digite \"pronto\" para informar que criou o arquivo")
	}

	conteudo, err := ioutil.ReadFile(ARQUIVO)
	if err != nil {
		fmt.Println("Arquivo não existe. Crie o arqiuvo e tente novamente")
	}

	fmt.Println("Conteudo do arquivo: ")
	fmt.Println(conteudo)
	conteudoStr := convertByteToString(conteudo)
	fmt.Println("Conteudo em string:", conteudoStr)
}
