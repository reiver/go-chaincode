package chaincode

import (
	"github.com/reiver/go-chainid"
)

const ZoraSepoliaTestnet = "ETH"

func init() {
	const value string =                   ZoraSepoliaTestnet
	var   key   string = createkey(chainid.ZoraSepoliaTestnet)

	chainCodes[key] = value
}
