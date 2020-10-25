package auctioneer

import (
	"auctions/common"
)

// TODO: Move some of them into common & also review them

// AuctionRequest :
type AuctionRequest struct {
	ID string `json:"auction_id"`
}

// AuctionResponse :
type AuctionResponse struct {
	ID      string      `json:"auction_id"`
	BestBid *common.Bid `json:"bid"`
}
