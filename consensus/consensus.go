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

package consensus

import (
	"github.com/ezcon-foundation/go-ezcon/core/transaction"
	"github.com/ezcon-foundation/go-ezcon/node/tcp"
	"sync"
	"time"
)

type Consensus struct {
	Transactions        []*transaction.Transaction
	proposalTransaction []*transaction.Transaction

	UNL          []string
	UNLPublicKey []string
	NodeID       string
	PrivKey      []byte
	Threshold    float64 // 0.8

	// tcp server
	server *tcp.TCPServer
	client *tcp.TCPClient

	isConsensing bool
	mutex        sync.Mutex

	proposalChan <-chan tcp.Message // Kênh cho đề xuất
	voteChan     <-chan tcp.Message // Kênh cho phiếu bầu
}

func NewConsensus(unl, unlPublicKey []string, nodeID string, privKey []byte, tpcPort string) *Consensus {

	// create tcp server
	server, err := tcp.NewTCPServer(tpcPort)
	if err != nil {
		return nil
	}

	// create tcp client
	client := tcp.NewTCPClient(2 * time.Second)

	// khởi tạo channel, giới hạn 100 giao dịch
	proposalChan := make(chan tcp.Message, 100)
	voteChan := make(chan tcp.Message, 100)

	// init consensus instance
	c := &Consensus{
		UNL:          unl,
		UNLPublicKey: unlPublicKey,
		NodeID:       nodeID,
		PrivKey:      privKey,
		Threshold:    0.8,
		server:       server,
		client:       client,
		proposalChan: proposalChan,
		voteChan:     voteChan,
		isConsensing: false,
	}

	// start tcp server
	go c.server.Start(c.IsConsensing, proposalChan, voteChan)

	return c
}

func (c *Consensus) getProposalTransaction() []*transaction.Transaction {

	// todo: choose available transaction for proposal
	return c.Transactions
}

func (c *Consensus) saveProposalTransaction(txs []*transaction.Transaction) {

	for _, tx := range txs {
		c.proposalTransaction = append(c.proposalTransaction, tx)
	}
}

// IsConsensing trả về trạng thái đồng thuận
func (c *Consensus) IsConsensing() bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.isConsensing
}
