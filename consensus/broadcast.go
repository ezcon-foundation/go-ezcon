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
	"encoding/json"
	"github.com/ezcon-foundation/go-ezcon/core/transaction"
	"github.com/ezcon-foundation/go-ezcon/network"
	"log"
	"sync"
)

// Broadcast gửi candidate set đến UNL
func (c *Consensus) Broadcast(txs []transaction.Transaction, sig []byte) error {

	// json marshal all transaction
	data, err := json.Marshal(txs)
	if err != nil {
		return err
	}

	// create message
	msg := network.Message{Txs: data, Sig: sig}

	var wg sync.WaitGroup
	errChan := make(chan error, len(c.UNL))

	// forloop all UNL and send via tcp
	for _, addr := range c.UNL {
		wg.Add(1)
		go func(addr string) {
			defer wg.Done()
			if err := c.client.Send(addr, msg); err != nil {
				errChan <- err
			}
		}(addr)
	}

	wg.Wait()
	close(errChan)
	for err := range errChan {
		log.Printf("Broadcast error: %v", err)
	}

	return nil
}
