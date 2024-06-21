package chaincode

import (
	"github.com/reiver/go-chainid"
)

const BlastSepoliaTestnet = "ETH"

func init() {
	const value string =                   BlastSepoliaTestnet
	var   key   string = createkey(chainid.BlastSepoliaTestnet)

	chainCodes[key] = value
}
