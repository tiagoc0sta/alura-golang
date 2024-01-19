package main

import "fmt"

func main() {

	exibeIntroducao()
	exibeMenu()
	comando := leComando() 

	/*if comando == 1 {
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
			fmt.Println("1- Iniciar Monitoramento")
		case 2:
			fmt.Println("Exibindo logs..")
		case 0:
			fmt.Println("Saindo do Programa")
		default:
			fmt.Println("Não conheço este comando")
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