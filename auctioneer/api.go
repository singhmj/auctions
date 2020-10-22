package auctioneer

import (
	"encoding/json"
	"net/http"
)

// registerBidder : registers the bidder
var registerBidder = func(w http.ResponseWriter, r *http.Request) {

}

// processAuctionRequest :
var processAuctionRequest = func(w http.ResponseWriter, r *http.Request) {
	// parse the bid request
	auctionRequest, err := parseAuctionRequest(r.Body)
	if err != nil {
		// write suitable http error code
		return
	}

	// err := validateAuctionRequest()

	// notify all available bidders
	//   - : separate go routines)
	//   - : [put max 200ms of wait, otherwise move to deciding the best bidder]
	bids, err = startAuction(Auction{ID: auctionRequest.ID})
	if err != nil {
		// write suitable http error code
		// and a message too
		return
	}

	// find the best bid
	bestBid, err := findBestBid(bids)

	// return reponse with best bidder
}

func parseAuctionRequest(data []byte) (auction AuctionRequest, err error) {
	err = json.Unmarshal(data, &auction)
	return auction, err
}
