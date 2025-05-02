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
	"context"
	"encoding/json"
	"github.com/ezcon-foundation/go-ezcon/core/transaction"
	"github.com/ezcon-foundation/go-ezcon/crypto"
	"github.com/ezcon-foundation/go-ezcon/network"
	"log"
	"sync"
	"time"
)

type Consensus struct {
	Transactions []*transaction.Transaction
	UNL          []string
	NodeID       string
	PrivKey      []byte
	Threshold    float64 // 0.8
	MaxRounds    int     // 5

	// tcp server
	server *network.TCPServer
	client *network.TCPClient

	isConsensing bool
	mutex        sync.Mutex

	proposalChan <-chan network.Message // Kênh cho đề xuất
	voteChan     <-chan network.Message // Kênh cho phiếu bầu
}

func NewConsensus(unl []string, nodeID string, privKey []byte, tpcPort string) *Consensus {

	// create tpc server
	server, err := network.NewTCPServer(tpcPort)
	if err != nil {
		return nil
	}

	// create tcp client
	client := network.NewTCPClient(2 * time.Second)

	// khởi tạo channel, giới hạn 1000 giao dịch
	proposalChan := make(chan network.Message, 100)
	voteChan := make(chan network.Message, 100)

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

	// start tpc server
	c.server.Start(c.IsConsensing, proposalChan, voteChan)

	return c
}

// RunConsensus runs the consensus algorithm
func (c *Consensus) RunConsensus() ([]transaction.Transaction, error) {

	return nil, nil
}

func (c *Consensus) handleVote(msg network.Message) {

}

func (c *Consensus) handleProposal(msg network.Message) {
	hasProposal := false
	var proposedTxs []*transaction.Transaction

	// Lặp qua các node có trong UNL, xác định giao dịch được gửi đến
	for _, node := range c.UNL {
		pubKey, err := crypto.PubKeyFromNode(node)
		if err != nil {
			continue
		}

		if crypto.Verify(msg.Txs, msg.Sig, pubKey) {

			// Cần phải phân biệt message nhận được thuộc loại message nào?
			var txs []*transaction.Transaction
			if err := json.Unmarshal(msg.Txs, &txs); err != nil {
				log.Printf("Invalid proposal: %v", err)
				continue
			}

			// Kiểm tra các giao dịch có hợp lệ không, nếu hợp lệ thì đưa vào danh sách những giao dịch hợp lệ
			// của node, lưu ý cần sắp xếp các giao dịch theo thứ tự sequence của account
			for _, tx := range txs {

				// todo: kiểm tra tính hợp lệ của tx
				proposedTxs = append(proposedTxs, tx)
			}

			hasProposal = true
			break
		}
	}

	// Nếu giao dịch gửi đến không thuộc bất kỳ một node nào đã biết, thì không xử lý
	if !hasProposal {
		return
	}

	// Kiểm tra điều kiện đồng thuận
	if !c.isConsensing {

		// khởi động trạng thái đồng thuận của node
		c.isConsensing = true

		go c.startConsensus(proposedTxs)

		hasProposal = false
		proposedTxs = nil
	}
}

func (c *Consensus) Run(ctx context.Context) {
	ticker := time.NewTicker(3 * time.Second) // Ticker 3 seconds
	defer ticker.Stop()
	defer c.server.Stop()

	for {
		select {

		case msg := <-c.proposalChan: // Condition 1: Nhận được đề xuất động thuận từ bất kỳ validator trong mạng
			c.mutex.Lock()
			c.handleProposal(msg)
			c.mutex.Unlock()
		case msg := <-c.voteChan:
			c.mutex.Lock()
			c.handleVote(msg)
			c.mutex.Unlock()
		case <-ticker.C: // Condition 2: Ticker 3 second

		}
	}
}

func (c *Consensus) startConsensus([]*transaction.Transaction) {

}

func (c *Consensus) getProposalTransaction() []*transaction.Transaction {

	// todo: choose some transaction
	return c.Transactions
}

// IsConsensing trả về trạng thái đồng thuận
func (c *Consensus) IsConsensing() bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.isConsensing
}
