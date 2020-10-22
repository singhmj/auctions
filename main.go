package main

// rough code
func main() {
	// StartAuctioneer()
	// StartBidders() // random number of bidders
}

func StartAuctioneer() {
	// restoreBidders()
	// startHTTPServer()
}

// func RegisterBidder()
func ProcessAuctionRequest() {
	// notify all available bidders (separate go routines) [put max 500ms of wait, otherwise move to deciding the best bidder]
	// find the best bidder amongst them
	// return reponse with best bidder
}

func StartBidder(delayInReplyingToBids int, portToListenBids int, urlToAuction int) {
	// start server to listen for any possible bids
	// register with the auctioner
}

type Auction struct {
	ID        string
	Bids      []*Bid
	WinnerBid *Bid
}

type Bid struct {
	ID        string
	AuctionID int
	BidderID  string
	Amount    float32
}
