package chaincode

import (
	"github.com/reiver/go-chainid"
)

const Blast = "ETH"

func init() {
	const value string =                   Blast
	var   key   string = createkey(chainid.Blast)

	chainCodes[key] = value
}
