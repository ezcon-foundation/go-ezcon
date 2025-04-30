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

package config

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli/v2"
)

type Config struct {
	NodeID     string   `toml:"node_id"`
	PrivKey    string   `toml:"priv_key"`
	UNL        []string `toml:"unl"`
	LedgerPath string   `toml:"ledger_path"`
	RPCPort    string   `toml:"rpc_port"`
}

func LoadConfig(ctx *cli.Context) (*Config, error) {
	cfg := &Config{
		NodeID:     "node1",
		PrivKey:    "privkey1",
		UNL:        []string{"node2:8081", "node3:8082"},
		LedgerPath: "./ledger.json",
		RPCPort:    "8080",
	}

	// Load từ file TOML nếu có
	if ctx.IsSet("config") {
		file := ctx.String("config")
		f, err := os.Open(file)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		_, err = toml.DecodeFile(file, cfg)
		if err != nil {
			return nil, err
		}
	}

	// Áp dụng flags từ CLI
	if ctx.IsSet("nodeid") {
		cfg.NodeID = ctx.String("nodeid")
	}
	if ctx.IsSet("privkey") {
		cfg.PrivKey = ctx.String("privkey")
	}
	if ctx.IsSet("unl") {
		cfg.UNL = ctx.StringSlice("unl")
	}
	if ctx.IsSet("ledgerpath") {
		cfg.LedgerPath = ctx.String("ledgerpath")
	}
	if ctx.IsSet("rpcport") {
		cfg.RPCPort = ctx.String("rpcport")
	}

	return cfg, nil
}
