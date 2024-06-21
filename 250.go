package chaincode

import (
	"github.com/reiver/go-chainid"
)

const FantomOpera = "FTM"

func init() {
	const value string =                   FantomOpera
	var   key   string = createkey(chainid.FantomOpera)

	chainCodes[key] = value
}
