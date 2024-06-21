package chaincode

import (
	"github.com/reiver/go-chainid"
)

const AvalancheCChain = "AVAX"

func init() {
	const value string =                   AvalancheCChain
	var   key   string = createkey(chainid.AvalancheCChain)

	chainCodes[key] = value
}
