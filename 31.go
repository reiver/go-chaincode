package chaincode

import (
	"github.com/reiver/go-chainid"
)

const RootstockTestnet = "tRBTC"

func init() {
	const value string =                   RootstockTestnet
	var   key   string = createkey(chainid.RootstockTestnet)

	chainCodes[key] = value
}
