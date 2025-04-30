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
)

type Consensus struct {
	UNL       []string
	NodeID    string
	PrivKey   []byte
	Threshold float64 // 0.8
	MaxRounds int     // 5
}

func NewConsensus(unl []string, nodeID string, privKey []byte) *Consensus {
	return &Consensus{
		UNL:       unl,
		NodeID:    nodeID,
		PrivKey:   privKey,
		Threshold: 0.8,
		MaxRounds: 5,
	}
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

	//TODO: Lưu trữ transaction ở đâu?
	//
	//currentTxs := c.Ledger.OpenLedger.ProposedTxs
	//if len(currentTxs) == 0 {
	//	return nil, nil
	//}
	//votes := make(map[string]int)
	//
	//for round := 1; round <= c.MaxRounds; round++ {
	//
	//	data, err := json.Marshal(currentTxs)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	sig, err := crypto.Sign(data, c.PrivKey)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	if err := c.Broadcast(currentTxs, sig); err != nil {
	//		log.Printf("Round %d: Broadcast failed: %v", round, err)
	//	}
	//
	//	for _, node := range c.UNL {
	//		receivedTxs, receivedSig, err := c.ReceiveFromNode(node)
	//		if err != nil {
	//			continue
	//		}
	//		pubKey, err := crypto.PubKeyFromNode(node)
	//		if err != nil {
	//			continue
	//		}
	//		if !crypto.Verify(receivedTxs, receivedSig, pubKey) {
	//			continue
	//		}
	//		var txs []types.Transaction
	//		if err := json.Unmarshal(receivedTxs, &txs); err != nil {
	//			continue
	//		}
	//
	//		for _, tx := range txs {
	//			data, _ := tx.Serialize()
	//			txID := crypto.SHA256(data)
	//			votes[string(txID)]++
	//		}
	//	}
	//
	//	newTxs := []types.Transaction{}
	//	threshold := c.Threshold * float64(round) / float64(c.MaxRounds)
	//	if threshold < 0.5 {
	//		threshold = 0.5
	//	}
	//	for _, tx := range currentTxs {
	//		data, _ := tx.Serialize()
	//		txID := crypto.SHA256(data)
	//		voteCount := votes[string(txID)]
	//		if float64(voteCount)/float64(len(c.UNL)) >= threshold {
	//			newTxs = append(newTxs, tx)
	//		}
	//	}
	//	currentTxs = newTxs
	//
	//	if len(currentTxs) > 0 {
	//		data, _ := json.Marshal(currentTxs)
	//		txSetID := crypto.SHA256(data)
	//		if float64(votes[string(txSetID)]) >= c.Threshold*float64(len(c.UNL)) {
	//			break
	//		}
	//	}
	//	time.Sleep(1 * time.Second)
	//}
	//
	//data, _ := json.Marshal(currentTxs)
	//txSetID := crypto.SHA256(data)
	//if float64(votes[string(txSetID)]) < c.Threshold*float64(len(c.UNL)) {
	//	return nil, errors.New("consensus not reached")
	//}

	return nil, nil
}
