package chaincode

import (
	"github.com/reiver/go-chainid"
)

const Sepolia = "ETH"

func init() {
	const value string =                   Sepolia
	var   key   string = createkey(chainid.Sepolia)

	chainCodes[key] = value
}
