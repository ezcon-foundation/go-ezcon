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
	"fmt"
	"github.com/ezcon-foundation/go-ezcon/config"
	"github.com/ezcon-foundation/go-ezcon/node"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
	"os"
)

var (
	ConfigFlag = &cli.StringFlag{
		Name:  "config",
		Usage: "TOML configuration file",
	}
)

func main() {

	app := &cli.App{
		Flags: []cli.Flag{
			ConfigFlag,
		},
		Name: "ezcon",
		Action: func(c *cli.Context) error {

			// load config
			cfg, err := config.LoadConfig(c)
			if err != nil {
				return fmt.Errorf("load config failed: %v", err)
			}

			node, err := node.NewNode(cfg)
			if err != nil {
				panic(err)
			}
			http.Handle("/rpc", node.RPCServer)

			log.Printf("Starting RPC server on :%s", cfg.RPCPort)

			return http.ListenAndServe(fmt.Sprintf(":%v", cfg.RPCPort), nil)
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
