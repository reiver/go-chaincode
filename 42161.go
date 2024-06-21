package chaincode

import (
	"github.com/reiver/go-chainid"
)

const ArbitrumOne = "ETH"

func init() {
	const value string =                   ArbitrumOne
	var   key   string = createkey(chainid.ArbitrumOne)

	chainCodes[key] = value
}
