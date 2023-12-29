package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const VERSION = 1.102
const SITE_OK = 1
const SITE_ERROR = 2
const DELAY = 5
const MAX_TENTATIVAS = 5
const LOG_FILE = "log.txt"

func salva_log() {
	logFile, err := os.Open(LOG_FILE)
	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo de logs ", LOG_FILE, ": ", err)
		os.Exit(1)
	}
}

func intro() {
	fmt.Println("Bem vindo ao monitor (versao: ", VERSION, ")")
}

func menu() {
	fmt.Println("Opções")
	fmt.Println("1: iniciar monitoramento")
	fmt.Println("2: exibir logs")
	fmt.Println("0: sair do monitor")
}

func getOption() int {
	var option int = 0
	fmt.Scan(&option)
	return option
}

func iniciarMonitoramento(site string) int {
	/*
		Monitora o site efetuando 5 tenativas
	*/
	for i := 0; i < MAX_TENTATIVAS; i++ {
		fmt.Println("Monitorando: ", site, " tentiva: ", i)
		response, _ := http.Get(site)
		if response != nil && response.StatusCode == 200 {
			return SITE_OK
		}
		_time := DELAY * time.Second
		time.Sleep(_time)
	}
	return SITE_ERROR
}

func main() {

	intro()
	var site = ""
	for {
		menu()
		var option int = getOption()
		fmt.Println("Opção digitada: ", option)
		switch option {
		case 1:
			fmt.Println("Monitoramento")
			fmt.Println("Que site deseja monitorar?")
			fmt.Scan(&site)
			resultado := iniciarMonitoramento(site)
			if resultado == SITE_OK {
				fmt.Println("Site OK")
			} else {
				fmt.Println("Não foi possível conectar ao site: ", site)
			}
		case 2:
			fmt.Println("Logs")
		case 0:
			fmt.Println("Saindo")
			os.Exit(0)
		default:
			fmt.Println("Opção não existe")
			os.Exit(0)
		}
	}
}
