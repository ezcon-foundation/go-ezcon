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
	"log"
	"time"
)

func (c *Consensus) RunEngine() {
	ticker := time.NewTicker(3 * time.Second) // Ticker 3 seconds
	defer ticker.Stop()
	defer c.server.Stop()

	for {
		select {
		case msg := <-c.proposalChan:
			c.mutex.Lock()
			c.handleProposal(msg)
			c.mutex.Unlock()
		case msg := <-c.voteChan:
			c.mutex.Lock()
			c.handleVote(msg)
			c.mutex.Unlock()
		case <-ticker.C:

			// nếu node đang trong qua trình đồng thuận, thì không thực hiện proposal mới
			if c.isConsensing {
				continue
			}

			c.mutex.Lock()

			log.Println("New Proposal...")
			// Lấy đề xuất các giao dịch
			proposedTxs := c.getProposalTransaction()

			// Lưu vào danh sách các giao dịch đang đề xuất
			c.saveProposalTransaction(proposedTxs)

			log.Println(c.UNL)

			// todo: Cần ký proposal transaction

			// Chuyển tiếp các giao dịch đề xuất cho các node trong danh sách UNL
			err := c.Broadcast(proposedTxs, []byte{})
			if err != nil {
				log.Printf("broadcast error", err)
				c.mutex.Unlock()
				continue
			}

			c.isConsensing = true

			c.mutex.Unlock()
		}
	}
}
