package main

import (
	"fmt"
	"os"
	"time"
)

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}
func contato() {
	contato := "hmm acho que consegui te convencer!! \n Aqui esta o telefone de Nicolas para contato (54)996457873\n Se preferir aqui esta suas redes sociais! \n Instagram: @nickcristofolli\n Linkedin: 'https://www.linkedin.com/in/nicolascristofolli/'"

	for _, char := range contato {
		fmt.Printf("%c", char)
		time.Sleep(100 * time.Millisecond) // Atraso de 100 milissegundos entre cada caractere
	}
}

func main() {
	olivia()
	mensagem()

	for {
		comando := leComando()

		switch comando {
		case 1:
			contato()
			os.Exit(1)
		case 2:
			fmt.Println("Saindo!!")
			os.Exit(2)
		default:
			fmt.Printf("Acho que nao entendi sua resposta, tente novamente. AUAU\n")
		}
	}
}
