package common

import "time"

// Bidder :
type Bidder struct {
	ID                      string `json:"id"`
	URLToNotifyAboutAuction string `json:"url_to_notify_about_auction"`
}

// Bid :
type Bid struct {
	ID        string
	AuctionID int
	BidderID  string
	Amount    float32
	CreatedAt time.Time
}

// Auction :
type Auction struct {
	ID        string
	Timeout   time.Duration
	Bids      []*Bid
	WinnerBid *Bid
}

// AuctionStats :
type AuctionStats struct {
	AuctionID        string
	IsAuctionOpen    bool
	TotalBids        int
	WinnerBid        *Bid
	AuctionStartedAt time.Time
	AuctionClosedAt  time.Time
}
