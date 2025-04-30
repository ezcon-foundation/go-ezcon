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

package node

import (
	"github.com/ezcon-foundation/go-ezcon/consensus"
	"github.com/ezcon-foundation/go-ezcon/core/ledger"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
	"log"
	"net/http"
	"time"
)

type Node struct {
	Ledger        *ledger.Ledger
	Consensus     *consensus.Consensus
	RPCServer     *rpc.Server
	IsInConsensus bool
}

type SubmitTxRequest struct {
	RawTx map[string]interface{} `json:"tx"`
}

type SubmitTxResponse struct {
	Status      string `json:"status"`
	TxID        string `json:"tx_id"`
	LedgerIndex uint64 `json:"ledger_index"`
}

func NewNode() (*Node, error) {
	s := rpc.NewServer()
	s.RegisterCodec(json2.NewCodec(), "application/json")
	node := &Node{
		RPCServer: s,
	}

	node.Consensus = &consensus.Consensus{
		Ledger:    node.Ledger,
		NodeID:    "node_id",
		PrivKey:   []byte("private_key"),
		Threshold: 0.8,
		MaxRounds: 5,
	}

	err := s.RegisterService(node, "ezcon")
	if err != nil {
		return nil, err
	}

	return node, nil
}

func (n *Node) RunValidator() {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if n.IsInConsensus {
				log.Println("Node is in consensus, skipping...")
				continue
			}

			log.Println("Running consensus...")

		}
	}
}

func (n *Node) TrustSet(r *http.Request, args *SubmitTxRequest, reply *SubmitTxResponse) error {

	log.Println("TrustSet called with args:", args)

	return nil
}
