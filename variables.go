package main

import "fmt"

func main() {
	nome := "Tiago"
	versao := 1.1
	

	fmt.Println("Olá Sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair d Programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O endereço da minha variavel comando é", &comando)
	fmt.Println("O comando escolhido foi", comando)

	if comando == 1 {
		fmt.Println("Monitorando..")
	} else if comando == 2 {
		fmt.Println("Exibindo logs..")
	} else if comando == 0 {
		fmt.Println("Saindo do Programa")
	} else {
		fmt.Println("Não conheço este comando")
	}
}