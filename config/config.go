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
	"encoding/hex"
	"errors"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli/v2"
)

type Config struct {
	NodeID        string   `toml:"node_id"`
	PrivKey       []byte   `toml:"private_key"`
	UNL           []string `toml:"unl"`
	LedgerPath    string   `toml:"ledger_path"`
	RPCPort       string   `toml:"rpc_port"`
	ConsensusPort string   `toml:"consensus_port"`
}

func LoadConfig(ctx *cli.Context) (*Config, error) {
	cfg := &Config{}

	// Load từ file TOML nếu có
	if ctx.IsSet("config") {
		file := ctx.String("config")
		f, err := os.Open(file)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		var tomlCfg struct {
			NodeID        string   `toml:"node_id"`
			PrivKey       string   `toml:"private_key"`
			UNL           []string `toml:"unl"`
			LedgerPath    string   `toml:"ledger_path"`
			RPCPort       string   `toml:"rpc_port"`
			ConsensusPort string   `toml:"consensus_port"`
		}

		_, err = toml.DecodeFile(file, &tomlCfg)
		if err != nil {
			return nil, err
		}

		cfg.NodeID = tomlCfg.NodeID
		if tomlCfg.PrivKey != "" {
			privKeyBytes, err := hex.DecodeString(tomlCfg.PrivKey)
			if err != nil || len(privKeyBytes) != 32 {
				return nil, errors.New("invalid hex private key in TOML")
			}
			cfg.PrivKey = privKeyBytes
		}

		cfg.UNL = tomlCfg.UNL
		cfg.LedgerPath = tomlCfg.LedgerPath
		cfg.RPCPort = tomlCfg.RPCPort
		cfg.ConsensusPort = tomlCfg.ConsensusPort
	}

	// Áp dụng flags từ CLI
	if ctx.IsSet("node_id") {
		cfg.NodeID = ctx.String("node_id")
	}
	if ctx.IsSet("private_key") {
		privKeyBytes, err := hex.DecodeString(ctx.String("private_key"))
		if err != nil || len(privKeyBytes) != 32 {
			return nil, errors.New("invalid hex private key in CLI")
		}
		cfg.PrivKey = privKeyBytes
	}
	if ctx.IsSet("unl") {
		cfg.UNL = ctx.StringSlice("unl")
	}
	if ctx.IsSet("ledger_path") {
		cfg.LedgerPath = ctx.String("ledger_path")
	}
	if ctx.IsSet("rpc_port") {
		cfg.RPCPort = ctx.String("rpc_port")
	}
	if ctx.IsSet("consensus_port") {
		cfg.ConsensusPort = ctx.String("consensus_port")
	}

	return cfg, nil
}
