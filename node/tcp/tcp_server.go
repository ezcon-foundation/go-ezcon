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

package tcp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

// TCPServer quản lý server nhận candidate set
type TCPServer struct {
	listener     net.Listener
	proposalChan chan<- Message
	voteChan     chan<- Message
	isConsensing func() bool
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

		s.handleConnection(conn)
	}
}

// Stop đóng server
func (s *TCPServer) Stop() {
	s.listener.Close()
}

// handleConnection xử lý kết nối
func (s *TCPServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	decoder := json.NewDecoder(reader)

	var msg Message
	err := decoder.Decode(&msg)
	if err != nil {
		fmt.Printf("Lỗi giải mã JSON hoặc client ngắt kết nối: %v\n", err)
		return
	}

	log.Printf("Receive msg: %+v", msg)

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
