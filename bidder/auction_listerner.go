package bidder

import (
	"auctions/common"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// listenAuctions :
func listenAuctions(ip, port, id string) error {
	// 1. create the mux
	mux := http.NewServeMux()

	// 2. register all the routes
	mux.HandleFunc("/ask", handleNewAuctionNotification)

	// 3. bind the server on the port and start serving requests
	log.Printf("listening for new auctions at: %v: %v", ip, port)
	err := http.ListenAndServe(ip+":"+port, mux)
	return err
}

var handleNewAuctionNotification = func(w http.ResponseWriter, r *http.Request) {
	// read the request body
	data, err := common.ReadHTTPBody(r.Body)
	if err != nil {
		log.Println("failed to read the auction notification request, more: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// parse auction notification
	auction, err := common.ParseAuctionNotification(data)
	if err != nil {
		log.Println("failed to parse the auction notification from the payload, more: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)

	// create a bid request
	bid := common.Bid{
		AuctionID: auction.ID,
	}

	// send the request on /bid http endpoint of auctioneer, after x ms of delay
	err = placeBid(bid)
	if err != nil {
		log.Println("failed to place bid to the auction system, more: %v", err)
	}
}

// registerWithAuctioneer :
func registerWithAuctioneer(auctioneerURL string, details common.Bidder) error {
	detailsInJSON, err := json.Marshal(&details)
	if err != nil {
		return fmt.Errorf("failed to convert bidder details to json, err: %v", err)
	}

	body := bytes.NewBuffer(detailsInJSON)
	req, err := http.NewRequest(http.MethodPost, auctioneerURL, body)
	if err != nil {
		return fmt.Errorf("failed to create http request, err: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// read body
	if resp.StatusCode != http.StatusOK {
		// TODO: extract more information from resp
		return fmt.Errorf("failed to register myself with the auctioneer, received %v from auctioneer", resp.StatusCode)
	}

	// it worked!
	return nil
}
