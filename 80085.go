package chaincode

import (
	"github.com/reiver/go-chainid"
)

const BerachainArtio = "BERA"

func init() {
	const value string =                   BerachainArtio
	var   key   string = createkey(chainid.BerachainArtio)

	chainCodes[key] = value
}
