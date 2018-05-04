package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	exibeIntroducao()
	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			inicarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func exibeIntroducao() {
	nome := "Rodolfo"
	versao := 1.1
	fmt.Println("Olá sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func leComando() int {
	var comandoLido int
	fmt.Scanf("%d", &comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")

	return comandoLido
}

func inicarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{"https://github.com", "https://minecraft.net/pt-br/", "https://code.visualstudio.com/", "https://www.mozilla.org/pt-BR/firefox/new/", "http://random-status-code.herokuapp.com/"}

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			monitoraSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func monitoraSite(site string) {
	res, _ := http.Get(site)

	if res.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status code:", res.StatusCode)
	}
}
