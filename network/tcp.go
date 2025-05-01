/*
 * Copyright (c) 2025 EZCON Foundation.
 *
 * The go-ezcon library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The go-ezcon library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with the go-ezcon library. If not, see <http://www.gnu.org/licenses/>.
 */

package network

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

// Message định nghĩa dữ liệu gửi/nhận qua TCP
type Message struct {
	Txs []byte `json:"txs"`
	Sig []byte `json:"sig"`
}

// TCPServer quản lý server nhận candidate set
type TCPServer struct {
	listener net.Listener
	msgChan  chan Message
	wg       sync.WaitGroup
}

// NewTCPServer khởi tạo server
func NewTCPServer(port string) (*TCPServer, error) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return nil, fmt.Errorf("failed to start TCP server on port %s: %v", port, err)
	}
	return &TCPServer{
		listener: listener,
		msgChan:  make(chan Message, 100),
	}, nil
}

// Start chạy server
func (s *TCPServer) Start() {
	go func() {
		defer s.listener.Close()
		for {
			conn, err := s.listener.Accept()
			if err != nil {
				log.Printf("TCP accept error: %v", err)
				return
			}
			s.wg.Add(1)
			go s.handleConnection(conn)
		}
	}()
}

// Stop đóng server
func (s *TCPServer) Stop() {
	s.listener.Close()
	s.wg.Wait()
	close(s.msgChan)
}

// Receive trả về channel để nhận message
func (s *TCPServer) Receive() <-chan Message {
	return s.msgChan
}

// handleConnection xử lý kết nối
func (s *TCPServer) handleConnection(conn net.Conn) {
	defer conn.Close()
	defer s.wg.Done()

	data := make([]byte, 4096)
	n, err := conn.Read(data)
	if err != nil {
		log.Printf("TCP read error: %v", err)
		return
	}

	var msg Message
	if err := json.Unmarshal(data[:n], &msg); err != nil {
		log.Printf("TCP unmarshal error: %v", err)
		return
	}

	select {
	case s.msgChan <- msg:
	default:
		log.Printf("Message channel full, dropping message")
	}
}

// TCPClient gửi candidate set đến node
type TCPClient struct {
	timeout time.Duration
}

// NewTCPClient khởi tạo client
func NewTCPClient(timeout time.Duration) *TCPClient {
	return &TCPClient{timeout: timeout}
}

// Send gửi message đến addr
func (c *TCPClient) Send(addr string, msg Message) error {
	conn, err := net.DialTimeout("tcp", addr, c.timeout)
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %v", addr, err)
	}
	defer conn.Close()

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = conn.Write(data)
	if err != nil {
		return fmt.Errorf("failed to send to %s: %v", addr, err)
	}
	return nil
}
