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
	"github.com/ezcon-foundation/go-ezcon/crypto"
	"github.com/ezcon-foundation/go-ezcon/network"
	"log"
)

// handleProposal sẽ tập trung vào việc xử lý các giao dịch đề xuất trong trạng thái nghỉ của validator
func (c *Consensus) handleProposal(msg network.Message) {
	hasProposal := false
	var proposedTxs []*transaction.Transaction

	// Lặp qua các node có trong UNL, xác định giao dịch được gửi đến
	for _, node := range c.UNL {
		pubKey, err := crypto.PubKeyFromNode(node)
		if err != nil {
			continue
		}

		// xác thực các giao dịch có phải đến từ các node đã biết hay không?
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

	// khởi động trạng thái đồng thuận của node
	c.isConsensing = true

	go c.startConsensus(proposedTxs)

	hasProposal = false
}
