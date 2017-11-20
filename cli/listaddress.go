package cli

import (
	"fmt"
	"github.com/sasaxie/go-blockchain/core"
	"log"
)

func (cli *CLI) listAddresses() {
	wallets, err := core.NewWallets()
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}
