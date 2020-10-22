package auctioneer

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"context"
)

// startAuction : it modifies the passed auction
// TODO: find a better design
func startAuction(auction *Auction) ([]*Bid, error) {
	// 1. fetch all the bidders from cache
	bidCriteria := ""
	bidders := getInterestedBidders(bidCriteria)
	if len(bidders) == 0 {
		return nil, fmt.Errorf("no bidder available for the auction")
	}

	// 2. notify all of them in different routines
	// let us create the book of bids, aka auction book
	ctxToControlBidders := context.WithDeadline()
	bidBook := make(chan *Bid, len(bidders))
	for _, bidder := range bidders {
		go func() {
			bid, err := notifyBidderAboutAuction(bidder)
			if err != nil {
				// TODO: hmm, what should be done over here?
				log.Printf("failed to notify bidder: %v, about auction: %v, err: %v")
			}
			bidBook <- bid
		}()
	}

	// 3. keep processing bids till auction.Timeout
	bids := make([]*Bid, 0)
	auctionCountdown := time.NewTimer(auction.Timeout * time.MilliSecond)
	for {
		select {
		case bid := <-bidBook:
			bids = append(bids, bid)
		case auctionClosureTime := <-auctionCountdown.C:
			close(bidBook)
			ctxToControlBidders.Cancel()
			log.Printf("auction: %v has been closed at: %v, received : %v bids so far", auction.ID, auctionClosureTime, len(bids))
			break
		}
	}
	auctionCountdown.Stop()
	return bids, nil
}

func notifyBidderAboutAuction(bidder *Bidder) (*Bid, error) {
	client := &http.Client{}
	_, err := client.Get(bidder.URLToAskForBid)
	if err != nil {
		return nil, err
	}

	// TODO:
	// err := parseBid(resp.Body)

	return nil, nil
}

func getInterestedBidders(criteria interface{}) []*Bidder {
	// TODO : fetch them from cache
	return []*Bidder{
		&Bidder{ID: "localhost", URLToAskForBid: "http://localhost:9999/bid"},
		&Bidder{ID: "localhost", URLToAskForBid: "http://localhost:9998/bid"},
		&Bidder{ID: "localhost", URLToAskForBid: "http://localhost:9997/bid"},
	}
}

func parseBid() (*Bid, error) {
	return nil, nil
}
