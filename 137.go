package chaincode

import (
	"github.com/reiver/go-chainid"
)

const PolygonMainnet = "MATIC"

func init() {
	const value string =                   PolygonMainnet
	var   key   string = createkey(chainid.PolygonMainnet)

	chainCodes[key] = value
}
