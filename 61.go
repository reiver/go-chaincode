package chaincode

import (
	"github.com/reiver/go-chainid"
)

const EthereumClassic = "ETC"

func init() {
	const value string =                   EthereumClassic
	var   key   string = createkey(chainid.EthereumClassic)

	chainCodes[key] = value
}
