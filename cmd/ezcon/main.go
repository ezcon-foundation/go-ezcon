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

package main

import (
	"github.com/ezcon-foundation/go-ezcon/node"
	"log"
	"net/http"
)

func main() {

	node, err := node.NewNode()
	if err != nil {
		panic(err)
	}

	http.Handle("/rpc", node.RPCServer)
	log.Printf("Starting RPC server on :%s", "3000")
	if err := http.ListenAndServe("0.0.0.0:3000", nil); err != nil {
		log.Fatalf("RPC server failed: %v", err)
	}

}
