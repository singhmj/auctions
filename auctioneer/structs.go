package auctioneer

import (
	"auctions/common"
	"time"
)

// TODO: Move some of them into common & also review them

// Auction :
type Auction struct {
	ID        string
	Timeout   time.Duration
	Bids      []*common.Bid
	WinnerBid *common.Bid
}

// AuctionRequest :
type AuctionRequest struct {
	ID string `json:"auction_id"`
}

// AuctionResponse :
type AuctionResponse struct {
	ID      string      `json:"auction_id"`
	BestBid *common.Bid `json:"bid"`
}
