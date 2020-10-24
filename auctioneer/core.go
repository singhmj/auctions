package auctioneer

import (
	"auctions/common"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	"context"
)

// startAuction :
// TODO: find a better design
func startAuction(ctx context.Context, auction *Auction) ([]*common.Bid, error) {
	// 1. fetch all the bidders from cache
	bidCriteria := ""
	bidders := getInterestedBidders(bidCriteria)
	if len(bidders) == 0 {
		return nil, fmt.Errorf("no bidder available for the auction")
	}

	// 2. notify all of them in different routines
	// let us create the book for bids, aka auction book
	ctxToControlBidders, cancelPendingBiddingRequests := context.WithCancel(context.Background())
	bidBook := make(chan *common.Bid, len(bidders))
	for _, bidder := range bidders {
		go func(ctx context.Context) {
			bid, err := engageBidderInAuction(ctx, auction, bidder)
			if err != nil {
				// TODO: hmm, what should be done over here?
				log.Printf("failed to notify bidder: %v, about auction: %v, err: %v")
				return
			}
			bidBook <- bid
		}(ctxToControlBidders)
	}

	// 3. keep processing bids till auction.Timeout
	bids := make([]*common.Bid, 0)
	auctionCountdown := time.NewTimer(auction.Timeout * time.Microsecond)
	for {
		select {
		case bid := <-bidBook:
			bids = append(bids, bid)
		case auctionClosureTime := <-auctionCountdown.C:
			cancelPendingBiddingRequests()
			close(bidBook)
			log.Printf("auction: %v has been closed at: %v, received : %v bids so far", auction.ID, auctionClosureTime, len(bids))
			break
		case <-ctx.Done():
			cancelPendingBiddingRequests()
			close(bidBook)
			// someone who started this auction now wants to end it, so let us do it on his/her command
			break
		}
	}

	auctionCountdown.Stop()
	return bids, nil
}

func engageBidderInAuction(ctx context.Context, auction *Auction, bidder *common.Bidder) (*common.Bid, error) {
	// REVIEW: should this be a GET request?
	req, err := http.NewRequest("GET", bidder.URLToAskForBid+"/"+auction.ID, nil)
	if err != nil {
		return nil, err
	}

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}

	// TODO:
	// err := parseBid(resp.Body)
	log.Printf("bidder replied with: %v\n", resp)

	return nil, nil
}

func getInterestedBidders(criteria interface{}) []*common.Bidder {
	// TODO : fetch them from cache
	return []*common.Bidder{
		&common.Bidder{ID: "localhost", URLToAskForBid: "http://localhost:9999/bid"},
		&common.Bidder{ID: "localhost", URLToAskForBid: "http://localhost:9998/bid"},
		&common.Bidder{ID: "localhost", URLToAskForBid: "http://localhost:9997/bid"},
	}
}

func parseBid() (*common.Bid, error) {
	return nil, nil
}

func findBestBid(bids []*common.Bid) (*common.Bid, error) {
	bestBid := &common.Bid{
		// TODO: use bitwise shift operation
		Amount: -(math.MaxFloat32 - 1),
	}
	for _, bid := range bids {
		if bid.Amount > bestBid.Amount {
			bestBid = bid
		}

		// REVIEW:
		// NOTE:
		// what if two bids are at the same price???
		// seems like we need to store the time when the bid actually was placed
		// and the winner then can be decided on the basis of reply time
	}

	return bestBid, nil
}
