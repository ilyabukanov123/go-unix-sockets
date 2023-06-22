package server

import (
	"context"
	"fmt"
	"net"
	"os"
)

// Сервер позволяет контролировать входящие сообщения от клиентов
type server struct {
	socketsPath string                                 // Адрес socket
	listener    net.Listener                           // Позволяет слушать входящие соединения
	handler     func(request []byte) (response []byte) // Кастомная функция по обработке сообщения
}

// NewServer создает новый сервер и возвращает его
func NewServer(socketsPath string, handler func(request []byte) (response []byte)) *server {
	var server = server{socketsPath: socketsPath, handler: handler}
	return &server
}

// ListenAndServe удаляет сокет если такой уже существует и создает новый сокет
func (s *server) ListenAndServe(ctx context.Context) error {
	if err := os.RemoveAll(s.socketsPath); err != nil {
		return fmt.Errorf("произошла ошибка при удалении сокета:%s", err)
	}
	sockets := s.socketsPath
	ln, err := net.Listen("unix", sockets)
	s.listener = ln
	if err != nil {
		return fmt.Errorf("произошла ошибка при создании сокета: %s", err)
	}
	return s.serve(ctx)
}

// serve читает входящих соединений и распределение их по горутинам
func (s *server) serve(ctx context.Context) error {
	defer s.listener.Close()
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			connect, err := s.listener.Accept()
			if err != nil {
				return fmt.Errorf("произошла ошибка при создании сокета для обмена данными с клиентом:%s", err)
			}

			go serve(connect, s, ctx)
		}
	}
}

// serve читает информацию из соединения и пишет в это соединение
func serve(connect net.Conn, srv *server, ctx context.Context) error {
	defer connect.Close()
	buf := make([]byte, 1024)
	length, err := connect.Read(buf)
	if err != nil {
		return fmt.Errorf("произошла ошибка при чтении сообщения из подключения:%s", err)
	}

	response := srv.handler(buf[:length])
	_, err = connect.Write(response)
	if err != nil {
		return fmt.Errorf("произошла ошибка при записи сообщения в подключение:%s", err)
	}
	return nil
}
