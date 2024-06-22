# go-chaincode

Package **chaincode** provides default **chain-codes** (i.e., (so called) **currency-codes**) for different blockchain-networks, for the Go programming language.

(For example, for Ethereum mainnet (i.e., chain-id=1) the default **chain-code** is `"ETH"` (and not `"wei"`).)

The data for the **chain currency-code**s comes from https://chainid.network/

This package provides those **chain-codes**s as **constants**, which are easily usable by programs written in the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-chaincode

[![GoDoc](https://godoc.org/github.com/reiver/go-chaincode?status.svg)](https://godoc.org/github.com/reiver/go-chaincode)

## Import

To import package **chaincode** use `import` code like the follownig:
```
import "github.com/reiver/go-chaincode"
```

## Installation

To install package **chaincode** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-chaincode
```

## Author

Package **chaincode** was written by [Charles Iliya Krempeaux](http://reiver.link)
