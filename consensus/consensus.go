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
}

func NewConsensus(unl []string, nodeID string, privKey []byte, tpcPort string) *Consensus {

	// create tpc server
	server, err := network.NewTCPServer(tpcPort)
	if err != nil {
		return nil
	}

	// create tcp client
	client := network.NewTCPClient(2 * time.Second)

	// init consensus instance
	c := &Consensus{
		UNL:       unl,
		NodeID:    nodeID,
		PrivKey:   privKey,
		Threshold: 0.8,
		MaxRounds: 5,
		server:    server,
		client:    client,
	}

	// start tpc server
	c.server.Start()

	return c
}

// RunConsensus runs the consensus algorithm
/*
 * Step 1: Propose transactions
 * Step 2: Sign the transactions
 * Step 3: Broadcast the transactions to UNL
 * Step 4: Collect votes from UNL
 * Step 5: Check if the votes reach the threshold
 * Step 6: If yes, commit the transactions to the ledger
 * Step 7: If no, repeat from step 1
 * Step 8: If the maximum rounds are reached, return an error
 * Step 9: If the transactions are committed, return the transactions
 */
func (c *Consensus) RunConsensus() ([]transaction.Transaction, error) {

	return nil, nil
}

func (c *Consensus) Run(ctx context.Context) {
	ticker := time.NewTicker(3 * time.Second) // Ticker 3 seconds
	defer ticker.Stop()
	defer c.server.Stop()

	hasProposal := false
	var proposedTxs []*transaction.Transaction

	for {
		select {

		case msg := <-c.server.Receive(): // Condition 1: Nhận được đề xuất động thuận từ bất kỳ validator trong mạng

			// Lặp qua các node liên kết có trong URL
			// Xác định signature được gửi tới
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
						_ = tx
					}

					hasProposal = true
					proposedTxs = txs
					break
				}

			}

			// Kiểm tra điều kiện đồng thuận
			if hasProposal && !c.isConsensing {
				go c.startConsensus(proposedTxs)
				hasProposal = false
				proposedTxs = nil
			}

		case <-ticker.C: // Condition 2: Ticker 3 second

			// Lấy danh sách các giao dịch đề xuất của validator
			proposalTxs := c.getProposalTransaction()

			if hasProposal && !c.isConsensing {
				go c.startConsensus(proposalTxs)
				hasProposal = false
				proposedTxs = nil
			} else if !c.isConsensing && len(proposalTxs) > 0 {
				go c.startConsensus(proposalTxs)
			}
		}
	}
}

func (c *Consensus) startConsensus([]*transaction.Transaction) {

}

func (c *Consensus) getProposalTransaction() []*transaction.Transaction {

	// todo: choose some transaction
	return c.Transactions
}
