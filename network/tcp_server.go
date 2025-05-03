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
)

// Message định nghĩa dữ liệu gửi/nhận qua TCP
type Message struct {

	// Danh sách các giao dịch
	Txs []byte `json:"txs"`

	// Chữ ký của node
	Sig []byte `json:"sig"`
}

// TCPServer quản lý server nhận candidate set
type TCPServer struct {
	listener     net.Listener
	proposalChan chan<- Message
	voteChan     chan<- Message
	isConsensing func() bool
	wg           sync.WaitGroup
}

// NewTCPServer khởi tạo server
func NewTCPServer(port string) (*TCPServer, error) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return nil, fmt.Errorf("failed to start TCP server on port %s: %v", port, err)
	}
	return &TCPServer{
		listener: listener,
	}, nil
}

// Start chạy server TCP
func (s *TCPServer) Start(isConsensing func() bool, proposalChan, voteChan chan<- Message) {

	log.Printf("Start Consensus TCP")

	s.isConsensing = isConsensing
	s.proposalChan = proposalChan
	s.voteChan = voteChan

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
}

// Stop đóng server
func (s *TCPServer) Stop() {
	s.listener.Close()
	s.wg.Wait()
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

	// Phân loại message dựa trên trạng thái đồng thuận
	if s.isConsensing != nil && s.isConsensing() {
		select {
		case s.voteChan <- msg:
		default:
			log.Printf("Vote channel full, dropping message")
		}
	} else {
		select {
		case s.proposalChan <- msg:
		default:
			log.Printf("Proposal channel full, dropping message")
		}
	}
}
