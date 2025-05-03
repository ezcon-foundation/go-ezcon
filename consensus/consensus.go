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
	tpc2 "github.com/ezcon-foundation/go-ezcon/node/tcp"
	"sync"
	"time"
)

type Consensus struct {
	Transactions        []*transaction.Transaction
	proposalTransaction []*transaction.Transaction

	UNL       []string
	NodeID    string
	PrivKey   []byte
	Threshold float64 // 0.8
	MaxRounds int     // 5

	// tcp server
	server *tpc2.TCPServer
	client *tpc2.TCPClient

	isConsensing bool
	mutex        sync.Mutex

	proposalChan <-chan tpc2.Message // Kênh cho đề xuất
	voteChan     <-chan tpc2.Message // Kênh cho phiếu bầu
}

func NewConsensus(unl []string, nodeID string, privKey []byte, tpcPort string) *Consensus {

	// create tcp server
	server, err := tpc2.NewTCPServer(tpcPort)
	if err != nil {
		return nil
	}

	// create tcp client
	client := tpc2.NewTCPClient(2 * time.Second)

	// khởi tạo channel, giới hạn 100 giao dịch
	proposalChan := make(chan tpc2.Message, 100)
	voteChan := make(chan tpc2.Message, 100)

	// init consensus instance
	c := &Consensus{
		UNL:          unl,
		NodeID:       nodeID,
		PrivKey:      privKey,
		Threshold:    0.8,
		MaxRounds:    5,
		server:       server,
		client:       client,
		proposalChan: proposalChan,
		voteChan:     voteChan,
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
