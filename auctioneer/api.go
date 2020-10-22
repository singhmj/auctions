package auctioneer

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// registerBidder : registers the bidder
var registerBidder = func(w http.ResponseWriter, r *http.Request) {
	// TODO:
}

// processAuctionRequest :
var processAuctionRequest = func(w http.ResponseWriter, r *http.Request) {
	// parse the bid request
	body, err := readHTTPBody(r.Body)
	if err != nil {
		// TODO: return with a suitable http code
		return
	}
	auctionRequest, err := parseAuctionRequest(body)
	if err != nil {
		// write suitable http error code
		return
	}

	// err := validateAuctionRequest()

	// start the auction
	auction := &Auction{ID: auctionRequest.ID}
	bids, err := startAuction(r.Context(), auction)
	if err != nil {
		// write suitable http error code
		// and a message too
		return
	}

	// find the best bid
	bestBid, err := findBestBid(bids)

	// return reponse with best bidder
	jsonResp, err := json.Marshal(bestBid)
	if err != nil {
		// return with apt http status code
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func parseAuctionRequest(data []byte) (auction AuctionRequest, err error) {
	err = json.Unmarshal(data, &auction)
	return auction, err
}

// TODO: Move this function into common
func readHTTPBody(httpBody io.ReadCloser) ([]byte, error) {
	data, err := ioutil.ReadAll(httpBody)
	return data, err
}
