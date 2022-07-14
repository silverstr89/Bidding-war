package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Bid struct {
	Bidder   common.Address
	BidValue big.Int
}

type LogBid struct {
	Bidder    common.Address
	Amount    *big.Int
	Timestamp *big.Int
}

type LogGameResult struct {
	From   common.Address
	Amount *big.Int
}
