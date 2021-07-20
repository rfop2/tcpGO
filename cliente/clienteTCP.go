package main

import "net"
import "fmt"
import "bufio"
import "os"

func conexaoCliente(){
	// conectando na porta 3030 via protocolo tcp/ip na máquina local
	conexao, erro1 := net.Dial("tcp", "127.0.0.1:3030")
	if erro1 != nil {
		fmt.Println(erro1)
		os.Exit(3)
	}

	defer conexao.Close()	

	requisicoes := 5

	for i:=0; i<=requisicoes; i++ {
		fmt.Println("Cliente Conectado ao Servidor")
		
		// lendo entrada 
		texto:="5"

		//checa se chegou no final
		if( i == requisicoes){
			texto = "para"
		}
		
		// escrevendo a mensagem na conexão (socket)
		fmt.Fprintf(conexao, texto+"\n")

		// ouvindo a resposta do servidor (eco)
		mensagem, err3 := bufio.NewReader(conexao).ReadString('\n')
		if err3 != nil {
			fmt.Println(err3)
			os.Exit(3)
		}
		// escrevendo a resposta do servidor no terminal
		fmt.Print("Resposta do servidor: " + mensagem)
	}
}

func main() {
	for i:=0; i<1; i++{
		go conexaoCliente()
	}
	fmt.Scanln()
}