package chaincode

import (
	"github.com/reiver/go-chainid"
)

const ArbitrumNova = "ETH"

func init() {
	const value string =                   ArbitrumNova
	var   key   string = createkey(chainid.ArbitrumNova)

	chainCodes[key] = value
}
