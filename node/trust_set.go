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
	"log"
	"net/http"
)

type TrustSetRequest struct {
	RawTx map[string]interface{} `json:"tx"`
}

type TrustSetResponse struct {
	Status      string `json:"status"`
	TxID        string `json:"tx_id"`
	LedgerIndex uint64 `json:"ledger_index"`
}

func (n *Node) TrustSet(r *http.Request, args *TrustSetRequest, reply *TrustSetResponse) error {

	log.Println("TrustSet called with args:", args)

	return nil
}
