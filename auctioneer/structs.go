package auctioneer

import "time"

type Auction struct {
	ID        string
	Timeout   time.Duration
	Bids      []*Bid
	WinnerBid *Bid
}

type Bid struct {
	ID        string
	AuctionID int
	BidderID  string
	Amount    float32
}

type BidderServer struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

type Bidder struct {
	ID             string `json:"id"`
	URLToAskForBid string `json:"url_to_ask_for_bid"`
}

// AuctionRequest :
type AuctionRequest struct {
	ID string `json:"auction_id"`
}

// AuctionResponse :
type AuctionResponse struct {
	ID      string `json:"auction_id"`
	BestBid Bid    `json:"bid"`
}
