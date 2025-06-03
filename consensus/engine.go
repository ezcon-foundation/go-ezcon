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
	"github.com/ezcon-foundation/go-ezcon/crypto"
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

			go func() {
				c.mutex.Lock()
				defer c.mutex.Unlock()

				// xử lý msg trong giai đoạn trạng thái engine chưa bắt đầu quá trình đồng thuận
				c.handleProposal(msg)
			}()

		case msg := <-c.voteChan:

			go func() {
				c.mutex.Lock()
				defer c.mutex.Unlock()

				// xử lý msg trong giai đoạn đang đồng thuận
				c.handleVote(msg)
			}()
		case <-ticker.C: // Đây là trường hợp node tự đề xuất trong 3 giây

			go func() {
				c.mutex.Lock()
				defer c.mutex.Unlock()

				// Nếu trạng thái engine đang đồng thuận, thì không gửi consensus
				if c.isConsensing {
					return
				}

				log.Println("Start send proposal transaction ...")

				// Lấy đề xuất các giao dịch
				proposedTxs := c.getProposalTransaction()

				// Lưu vào danh sách các giao dịch đang đề xuất
				c.saveProposalTransaction(proposedTxs)

				// marshal data
				data, err := json.Marshal(proposedTxs)
				if err != nil {
					log.Println("can not marshal proposal txs", err)
					return
				}

				// ký proposal transaction
				signature, err := crypto.Sign(data, c.PrivKey)
				if err != nil {
					log.Println("can not sign data", err)
					return
				}

				// Chuyển tiếp các giao dịch đề xuất cho các node trong danh sách UNL
				err = c.Broadcast(proposedTxs, signature)
				if err != nil {
					return
				}

				c.isConsensing = true

				log.Println("Start consensus...")
			}()

		}
	}
}
