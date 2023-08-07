package main

import (
	"fmt"
	"log"
	"net/rpc"
	"rpcsrv/server"
)

const addr = ":8888"
const network = "tcp4"

func main() {
	// Создаем клиента службы RPC.
	client, err := rpc.DialHTTP(network, addr)
	if err != nil {
		log.Fatal(err)
	}

	var args = server.Points{A: server.Point{X: 7, Y: 5},
		B: server.Point{X: 4, Y: 1}}

	var resp float64

	// Удаленный вызов процедуры Server.Dist. Должен быть ответ.
	err = client.Call("Server.Dist", args, &resp)
	if err != nil {
		fmt.Println("ошибка:", err)
	}
	fmt.Println("distance:", resp)

}
