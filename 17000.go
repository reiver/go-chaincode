package chaincode

import (
	"github.com/reiver/go-chainid"
)

const Holesky = "ETH"

func init() {
	const value string =                   Holesky
	var   key   string = createkey(chainid.Holesky)

	chainCodes[key] = value
}
