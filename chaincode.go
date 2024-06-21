package chaincode

var chainCodes map[string]string = map[string]string{}

func ChainCode(chainid uint64) string {

	var key string = createkey(chainid)

	code, found := chainCodes[key]
	if !found {
		return ""
	}

	return code
}
