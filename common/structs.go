package common

type Bidder struct {
	ID             string `json:"id"`
	URLToAskForBid string `json:"url_to_ask_for_bid"`
}

type Bid struct {
	ID        string
	AuctionID int
	BidderID  string
	Amount    float32
}
