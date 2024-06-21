package chaincode

import (
	"github.com/reiver/go-chainid"
)

const OPMainnet = "ETH"

func init() {
	const value string =                   OPMainnet
	var   key   string = createkey(chainid.OPMainnet)

	chainCodes[key] = value
}
