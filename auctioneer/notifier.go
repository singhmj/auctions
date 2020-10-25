package auctioneer

import (
	"auctions/common"
	"context"
	"fmt"
	"log"
	"net/http"
)

func notifyBidderAboutAuction(ctx context.Context, auction *common.Auction, bidder *common.Bidder) error {
	// REVIEW: should this be a GET request?
	req, err := http.NewRequest("POST", bidder.URLToNotifyAboutAuction+"/"+auction.ID, nil)
	if err != nil {
		return err
	}

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return err
	}

	log.Printf("bidder replied with: %v\n", resp)
	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("bidder failed to reply properly, err: %v", err)
	}

	return nil
}
