package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"rpcsrv/server"
)

const addr = ":8888"
const network = "tcp4"

func main() {
	// Создаем указатель на переменную типа Server.
	srv := new(server.Server)
	// Регистрируем методы типа Server в службе RPC.
	rpc.Register(srv)
	// Регистрируем HTTP-обработчик для службы RPC.
	rpc.HandleHTTP()
	// Создаём сетевую службу.
	listener, err := net.Listen(network, addr)
	if err != nil {
		log.Fatal(err)
	}
	// Запускаем HTTP-сервер поверх созданной сетевой службы.
	http.Serve(listener, nil)
}
