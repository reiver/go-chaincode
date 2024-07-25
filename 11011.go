package chaincode

import (
	"github.com/reiver/go-chainid"
)

const ShapeSepolia = "ETH"

func init() {
	const value string =                   ShapeSepolia
	var   key   string = createkey(chainid.ShapeSepolia)

	chainCodes[key] = value
}
