package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ilyabukanov123/go-unix-sockets/pkg/server"
)

func main() {
	srv := server.NewServer("/tmp/at.sock", func(request []byte) []byte {
		fmt.Println("Сообщение от клиента:", string(request))
		return []byte("Я сервер и я принял сообщение")
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err := srv.ListenAndServe(ctx)
	if err != nil {
		fmt.Println("Произошла ошибка при чтении данных из соединения")
	}
}
