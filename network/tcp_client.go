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
	"net"
	"time"
)

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
