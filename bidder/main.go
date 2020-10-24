package bidder

import (
	"auctions/common"
	"fmt"
	"log"
	"os"
)

// POST request
var endpointToNotifyMeAboutAuction = "/notification/auction"
var auctioneerURL = "localhost:8080/bidder/register"

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("insufficient number of arguments, usage: [./bidder ip port bidderid]")
	}

	// let us assume that the program is passed correct params, in correct sequence
	ip, port, bidderID := os.Args[1], os.Args[2], os.Args[3]

	if ip == "" {
		log.Fatalf("empty ip is not allowed in program arguments")
	}
	if port == "" {
		log.Fatalf("empty port is not allowed in program arguments")
	}
	if bidderID == "" {
		log.Fatalf("empty bidderid is not allowed in program arguments")
	}

	start(ip, port, bidderID)
	// defer stop()

	// TODO: make this main routine, wait on signal handler
	// handleSystemSignals()
}

func start(ip, port, bidderID string) {
	details := common.Bidder{
		ID:             bidderID,
		URLToAskForBid: fmt.Sprintf("http://%v:%v/%v", ip, port, endpointToNotifyMeAboutAuction),
	}
	log.Println("registering myself with the auctioneer ...")
	if err := registerWithAuctioneer(auctioneerURL, details); err != nil {
		log.Fatal("failed to register myself with the auctioneer, more: %v", err)
	}
	log.Println("successfully registered with the auctioneer")

	// go func() {
	log.Println("starting auction listner ...")
	if err := listenAuctions(ip, port, bidderID); err != nil {
		log.Fatal("failed to start auction listener, more: %v", err)
	}
	log.Println("listening for new auction events")
	// }()
}

func stop() {
}

func handleSystemSignals() {
}
