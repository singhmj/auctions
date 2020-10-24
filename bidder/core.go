package bidder

import "auctions/common"

// NewBid :
func NewBid(auctionID string, criteria string) *common.Bid {
	return &common.Bid{}
}

func placeBid(bid *common.Bid) error {
	// calculate whatever metrics you wanna do, before placing the bidr
	return nil
}
