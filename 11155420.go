package chaincode

import (
	"github.com/reiver/go-chainid"
)

const OPSepoliaTestnet = "ETH"

func init() {
	const value string =                   OPSepoliaTestnet
	var   key   string = createkey(chainid.OPSepoliaTestnet)

	chainCodes[key] = value
}
