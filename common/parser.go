package common

import "encoding/json"

// ParseAuctionNotification :
func ParseAuctionNotification(data []byte) (*Auction, error) {
	auction := Auction{}
	err := json.Unmarshal(data, &auction)
	return auction, err
}
