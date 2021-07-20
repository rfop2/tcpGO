package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func factorial(x int) (result int) {
	if x == 0 {
	  result = 1;
	} else {
	  result = x * factorial(x - 1);
	}
	return;
}

func conexaoServidor(conexao net.Conn){

	fmt.Println("Conexão feita com o cliente")
	
	for {
	
		mensagem, erro3 := bufio.NewReader(conexao).ReadString('\n')
		if erro3 != nil {
			fmt.Println(erro3, "ss")
			return
		}

		// se chegou no final encerra a conexao
		temp := strings.TrimSpace(string(mensagem))
		if(temp == "para"){
			fmt.Println("Conexão encerrada com o cliente")
			break
		}
		
		fmt.Print("Número recebido: ", string(mensagem))
		
		numero := 0
		msg := string(mensagem)
		fmt.Sscan(msg, &numero)
		
		resultado := factorial(numero)

		novamensagem := strconv.Itoa(resultado)
		
		conexao.Write([]byte(novamensagem + "\n"))
	}
}

func main() {
	ln, erro1 := net.Listen("tcp", ":3030")
	if erro1 != nil {
		fmt.Println(erro1)
		os.Exit(3)
	}

	defer ln.Close()
	fmt.Println("Aguardando conexões com clientes...")
	for{
		// aceitando conexões
		conexao, erro2 := ln.Accept()
		if erro2 != nil {
			fmt.Println(erro2)
			return
		}

		go conexaoServidor(conexao)
	}
	
}