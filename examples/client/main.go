package main

import (
	"fmt"

	"github.com/ilyabukanov123/go-unix-sockets/pkg/client"
)

func main() {
	cnt, err := client.NewClient("/tmp/at.sock")
	if err != nil {
		panic(err)
	}

	response, err := cnt.Request([]byte("Я клиент и я отправляю сообщение на сервер"))
	if err != nil {
		panic(err)
	}

	fmt.Println("Ответ от сервера: ", string(response))
	cnt.Close()
}
