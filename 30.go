package chaincode

import (
	"github.com/reiver/go-chainid"
)

const RootstockMainnet = "RBTC"

func init() {
	const value string =                   RootstockMainnet
	var   key   string = createkey(chainid.RootstockMainnet)

	chainCodes[key] = value
}
