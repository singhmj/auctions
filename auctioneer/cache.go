package auctioneer

import (
	"sync"
)

var (
	singletonCache *cache.Cache
	once           sync.Once
)

// GetCache :
func GetCache() *cache.Cache {
	if singletonCache == nil {
		once.Do(func() {
			singletonCache = cache.New()
		})
	}

	return singletonCache
}

func saveAuctionBook(book *auctionBook) {
	GetCache().Set("auction-book-" + book.AuctionID)
}

// getAuctionBook :
func getAuctionBook(auctionID string) (*auctionBook, bool) {
	data, doesExist := GetCache().Get("auction-book-" + auctionID)
	if !doesExist {
		return nil, doesExist
	}

	return data.(*auctionBook), doesExist
}

// createAuctionBook :
func createAuctionBook(auction *Auction) error {
	GetCache().Set("auction-book"+auction.ID, auction)
}

func getBidders() {
	// GetCache().
}

func addBidToAuction(auctionID string, bid *Bid) {
	// TODO: this is not thread safe, so review this later
	// getAuctionBook(auctionID)
}

func removeAuction() {

}

func addBid() {

}

func addBidder() {

}

func removeBid() {

}

func removeBidder() {

}
