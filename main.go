package main

import (
	"coin_tutorial/blockchain"
	"coin_tutorial/cli"
	"coin_tutorial/db"
)

func main() {
	defer db.Close()
	blockchain.Blockchain()
	cli.Start()
}
