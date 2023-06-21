package client

import (
	"fmt"
	"net"
)

// Клиент устанавливает соединение с сервером
type client struct {
	conn net.Conn // Позволяет подключиться к серверу
}

// NewClient создает нового клиента по работе с сервером
func NewClient(socketsPath string) (*client, error) {
	conn, err := net.Dial("unix", socketsPath)
	if err != nil {
		return nil, fmt.Errorf("произошла ошибка при подключении к сокету:%s", err)
	}
	return &client{conn: conn}, nil
}

// Request пишет сообщение в соединение и читает сообщение от сервера
func (c *client) Request(request []byte) ([]byte, error) {
	_, err := c.conn.Write(request)
	if err != nil {
		return nil, fmt.Errorf("произошла ошибка при записи сообщения в соединение:%s", err)
	}

	buf := make([]byte, 1024)
	length, err := c.conn.Read(buf)
	if err != nil {
		return nil, fmt.Errorf("произошла ошибка при чтении сообщения из соединения:%s", err)
	}

	return buf[:length], nil
}

// Close закрывает соединение с сервером
func (c *client) Close() error {
	return c.conn.Close()
}
