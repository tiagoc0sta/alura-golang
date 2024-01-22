package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntroducao()
	exibeMenu()
	comando := leComando() 

	/*if comando == 1 {                 // if no golang sempre retorna booleano (true or false)
		fmt.Println("Monitorando..")
	} else if comando == 2 {
		fmt.Println("Exibindo logs..")
	} else if comando == 0 {
		fmt.Println("Saindo do Programa")
	} else {
		fmt.Println("Não conheço este comando")
	}*/

	switch comando {
		case 1:			
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs..")
		case 0:
			fmt.Println("Saindo do Programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1) // indica que ocorreu algum erro no programa
	}
}


func exibeIntroducao(){
	nome := "Tiago"
	versao := 1.1
	fmt.Println("Olá Sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu(){
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair d Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O endereço da minha variavel comando é", &comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}


func iniciarMonitoramento() {
	fmt.Println("1- Iniciar Monitoramento")
	site := "https://www.alura.com.br"
	resp, _ := http.Get(site)
	fmt.Println(resp)	 
}