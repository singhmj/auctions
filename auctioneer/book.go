package auctioneer

import (
	"auctions/common"
	"context"
	"log"
	"math"
	"time"
)

// auctionBook :
type auctionBook struct {
	AuctionID         string
	bidsChannel       chan *common.Bid
	totalReceivedBids int
	deadline          time.Duration
	isOpen            bool
	openedAt          time.Time
	closedAt          time.Time
}

// openAuctionBook : please use this function to create a new auction book
func openAuctionBook(auctionID string, size int) *auctionBook {
	book := newBook(auctionID, size)
	saveAuctionBook(book)
	return book
}

// newBook :
func newBook(auctionID string, size int) *auctionBook {
	return &auctionBook{
		AuctionID:   auctionID,
		bidsChannel: make(chan *common.Bid, size),
	}
}

func (book *auctionBook) IsBookOpen() bool {
	return book.isOpen
}

func (book *auctionBook) Start(ctx context.Context) {
	auctionCountdown := time.NewTimer(book.deadline * time.Microsecond)
	for {
		select {
		// case bid := <-book.bidsChannel:
		// 	bids = append(bids, bid)
		case auctionClosureTime := <-auctionCountdown.C:
			close(book.bidsChannel)
			log.Printf("auction: %v has been closed at: %v, received : %v bids so far", book.AuctionID, auctionClosureTime, len(bids))
			break
		case <-ctx.Done():
			// someone who started this auction now wants to end it, so let us do it on his/her command
			close(book.bidsChannel)
			break
		}
	}

	auctionCountdown.Stop()
}

func (book *auctionBook) FindBestBid() *common.Bid {
	var bestBid *common.Bid
	maxBidAmount := -(math.MaxFloat32 - 1) // rather use shift operation

	for bid := range book.bidsChannel {
		if bid.Amount > maxBidAmount {
			bestBid = bid
			maxBidAmount = bid.Amount
		}
		book.totalReceivedBids++

		// REVIEW:
		// NOTE:
		// what if two bids are at the same price???
		// seems like we need to store the time when the bid actually was placed
		// and then we can decide the winner on the basis of reply time
	}

	return bestBid
}

// Stop :
func (book *auctionBook) Stop() {

}

func (book *auctionBook) getStats() common.AuctionStats {
	return common.AuctionStats{
		AuctionID:        book.AuctionID,
		TotalBids:        book.totalReceivedBids,
		IsAuctionOpen:    book.IsBookOpen(),
		WinnerBid:        book.FindBestBid(),
		AuctionStartedAt: book.openedAt,
		AuctionClosedAt:  book.closedAt,
	}
}
