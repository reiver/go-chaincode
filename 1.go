package chaincode

import (
	"github.com/reiver/go-chainid"
)

const EthereumMainnet = "ETH"

func init() {
	const value string =                   EthereumMainnet
	var   key   string = createkey(chainid.EthereumMainnet)

	chainCodes[key] = value
}
