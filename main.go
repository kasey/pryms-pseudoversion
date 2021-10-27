package main

import (
	"fmt"
	"github.com/prysmaticlabs/prysm/config/params"
)

func main() {
	cfg := params.MainnetConfig()
	fmt.Printf("MinDepositAmount: %d\n", cfg.MinDepositAmount)
}
