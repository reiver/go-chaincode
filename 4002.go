package chaincode

import (
	"github.com/reiver/go-chainid"
)

const FantomTestnet = "FTM"

func init() {
	const value string =                   FantomTestnet
	var   key   string = createkey(chainid.FantomTestnet)

	chainCodes[key] = value
}
