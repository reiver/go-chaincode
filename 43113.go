package chaincode

import (
	"github.com/reiver/go-chainid"
)

const AvalancheFujiTestnet = "AVAX"

func init() {
	const value string =                   AvalancheFujiTestnet
	var   key   string = createkey(chainid.AvalancheFujiTestnet)

	chainCodes[key] = value
}
