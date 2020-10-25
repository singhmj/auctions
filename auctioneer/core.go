package auctioneer

import (
	"auctions/common"
	"fmt"
	"log"

	"context"
)

// startAuction :
func startAuction(ctx context.Context, auction *common.Auction) error {
	// 1. fetch all the bidders from cache
	bidCriteria := ""
	bidders := getInterestedBidders(bidCriteria)
	if len(bidders) == 0 {
		return fmt.Errorf("no bidder available for the auction")
	}

	// 2. let us create the book for bids, aka auction book
	auctionBook := openAuctionBook(auction.ID, len(bidders))
	go auctionBook.Run(ctx)

	// 3. notify all of the bidders about this new auction in different routines
	for _, bidder := range bidders {
		go func(ctx context.Context) {
			err := notifyBidderAboutAuction(ctx, auction, bidder)
			if err != nil {
				// TODO: hmm, what should be done over here? should drop this bidder?
				log.Printf("failed to notify bidder: %v, about auction: %v, err: %v")
				return
			}
		}(ctx)
	}

	return nil
}

func getInterestedBidders(criteria interface{}) []*common.Bidder {
	// TODO : fetch them from cache
	return []*common.Bidder{
		&common.Bidder{ID: "localhost", URLToNotifyAboutAuction: "http://localhost:9999/bid"},
		&common.Bidder{ID: "localhost", URLToNotifyAboutAuction: "http://localhost:9998/bid"},
		&common.Bidder{ID: "localhost", URLToNotifyAboutAuction: "http://localhost:9997/bid"},
	}
}

func parseBid() (*common.Bid, error) {
	return nil, nil
}
