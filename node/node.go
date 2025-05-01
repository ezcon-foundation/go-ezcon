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
	"github.com/ezcon-foundation/go-ezcon/config"
	"github.com/ezcon-foundation/go-ezcon/consensus"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
)

type Node struct {
	Consensus *consensus.Consensus
	RPCServer *rpc.Server
}

func NewNode(cfg *config.Config) (*Node, error) {

	// create rpc server
	s := rpc.NewServer()

	// regis codec for rpc server
	s.RegisterCodec(json2.NewCodec(), "application/json")

	// init node parameter
	node := &Node{
		RPCServer: s,
		Consensus: consensus.NewConsensus(
			cfg.UNL,
			cfg.NodeID,
			cfg.PrivKey,
			cfg.ConsensusPort,
		),
	}

	// regis server under name 'ezcon'
	err := s.RegisterService(node, "ezcon")
	if err != nil {
		return nil, err
	}

	return node, nil
}
