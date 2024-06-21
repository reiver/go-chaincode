package chaincode

import (
	"github.com/reiver/go-chainid"
)

const Zora = "ETH"

func init() {
	const value string =                   Zora
	var   key   string = createkey(chainid.Zora)

	chainCodes[key] = value
}
