package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 2
const delay = 5

func main() {

	exibeIntroducao()

	for {
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
				imprimeLogs()
			case 0:
				fmt.Println("Saindo do Programa")
				os.Exit(0)
			default:
				fmt.Println("Não conheço este comando")
				os.Exit(-1) // indica que ocorreu algum erro no programa
		}

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
	fmt.Println("")

	return comandoLido
}


func iniciarMonitoramento() {
	fmt.Println("Monitorando...")	
	//sites:= []string{"https://httpbin.org/status/40", "https://www.alura.com.br", "https://www.caelum.com.br"}
	
	sites := lerSitesDoArquivo()

	//example using for loop
	/*for i:=0; i< len(sites); i++ {
		fmt.Println(sites[i])		
	}*/

	for i:=0; i< monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	
	fmt.Println("")
	
	
}

func testaSite(site string) {
	if site == "" {
			fmt.Println("Site vazio!")
			return
	}

	resp, err := http.Get(site)

	if err != nil {
			fmt.Println("Ocorreu um erro: ", err)
			return
	}

	if resp.StatusCode == 200 {
			fmt.Println("Site:", site, "foi carregado com sucesso!")
			registraLog(site, true)
	} else {
			fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
			registraLog(site, false)
	}
}


func lerSitesDoArquivo()[]string {

	var sites []string

	//arquivo, err := os.ReadFile("sites.txt") //ler todos os arquivos de uma so vez 
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	leitor := bufio.NewReader(arquivo)
	
	for{
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		//fmt.Println(linha)
		
		sites = append(sites, linha)
		
		if err == io.EOF {
			break
		}	

	}	

	arquivo.Close()

	return sites
}


func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	
	arquivo.Close()
}

func imprimeLogs(){
	arquivo, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}