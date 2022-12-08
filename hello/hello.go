package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = monitoramentos + 2

func main() {
	showIntroduction()

	for {

		showMenu()

		commandInt := readCommand()

		switch commandInt {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Exibindo Logs")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
	}

}

func showIntroduction() {
	namePeople := "Felipe"
	versionGoLang := runtime.Version()

	fmt.Println("Olá sr.", namePeople)
	fmt.Println("Esse programa está na versão ", versionGoLang)
}

func showMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")
}

func readCommand() int {
	var commandReaded int
	fmt.Scan(&commandReaded)

	fmt.Println("O comando escolhido foi", commandReaded)
	fmt.Println(" ")

	return commandReaded
}

func startMonitoring() {
	fmt.Println("Monitorando...")

	// sites := []string{"https://www.alura.com.br", "https://www.zapsign.com.br", "https://random-status-code.herokuapp.com/"}

	sites := readSitesFile()

	for i := 0; i < monitoramentos; i++ {
		for idx, site := range sites {
			fmt.Println("Testando site", idx, ":", site)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println(" ")

	}

	fmt.Println(" ")
}

func testSite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso")
		registerLogs(site, true)
	} else {
		fmt.Println("Site:", site, "não foi carregado e está com problemas. Status code:", response.StatusCode)
		registerLogs(site, false)
	}
}

func readSitesFile() []string {
	var sites []string
	file, err := os.Open("sites.txt") // Abre um arquivo

	// file, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}
	reader := bufio.NewReader(file) // Cria um leitor de texto

	for {
		line, err := reader.ReadString('\n') // Le o texto
		line = strings.TrimSpace(line)
		fmt.Println(line)

		sites = append(sites, line)

		if err == io.EOF { // Saia do loop ao encontrar o erro EOF (End of Line)
			break
		}

	}

	file.Close()

	return sites
}

func registerLogs(site string, status bool) {
	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755) // Ler documentação pra entrar melhor a OpenFile

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}
