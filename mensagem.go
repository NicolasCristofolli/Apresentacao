package main

import (
	"fmt"
	"time"
)

func mensagem() {

	apresentacao := "Oie, Eu sou a Olivia, assistente virtual do Nicolas. Fui criada para apresentar o meu programador e, te convencer à dar uma chance para ele (o garoto é gente boa):"

	for _, char := range apresentacao {
		fmt.Printf("%c", char)
		time.Sleep(100 * time.Millisecond) // Atraso de 100 milissegundos entre cada caractere
	}

	fmt.Println() // Adiciona uma nova linha após a animação

	mensagem := "Nicolas é apaixonado por games e tecnologia e isso fez ele querer cursar Sistemas de informação na UNISINOS e me criar também!! \n Ele atualmente trabalha como técnico de informática, mas mesmo assim já fez diversos cursos de programação e agora está em busca de seu primeiro estágio na área de programação! \n Seu nível de inglês é intermediario (B1), mas está estudando todos os dias para 'improve' cada vez mais kkkk \n É colaborativo, comunicativo, ah, quase me esqueci, suas linguagens favoritas atualmente sao: Go (na qual fui criada), JAVA e C# \n E ai, te convenci? AU AU \n Tecle 01 para mandar uma mensagem ao Nicolas \n Tecle 02 Para sair"

	for _, char := range mensagem {
		fmt.Printf("%c", char)
		time.Sleep(100 * time.Millisecond) // Atraso de 100 milissegundos entre cada caractere
	}
	fmt.Println() // Adiciona uma nova linha após a animação
}
