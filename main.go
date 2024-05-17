package main

import (
	"log"
	"net"
	"os"

	"sf035/handler"
	"sf035/proverb"
)

const defaultAddress = "0.0.0.0:64777"
const network = "tcp4"

func main() {
	// Получение адреса и порта из аргументов или подстановка значений
	// по умолчанию
	var address string
	if len(os.Args) > 1 {
		address = os.Args[1]
	} else {
		address = defaultAddress
	}

	// Создание сборника пословиц
	proverbReader, err := os.Open("proverb.txt")
	if err != nil {
		log.Fatal(err)
	}
	proverb, err := proverb.NewProverb(proverbReader)

	// Запуск TCP-сервера
	listener, err := net.Listen(network, address)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handler.Handle(conn, proverb)
	}
}
