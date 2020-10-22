package auctioneer

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Start : Starts the auctioneer
func Start(config Config) error {
	// 1. restore the already registered bidders
	// restoreBidders()

	// 2. start the http server and listen to requests
	err := startHTTPServer(config.IP, config.Port)
	if err != nil {
		return fmt.Errorf("failed to start http server: %v", err)
	}
}

func startHTTPServer(ip string, port int) error {
	// 1. create the mux
	mux := http.NewServeMux()

	// 2. register all the routes
	registerRoutes(mux)

	// 3. bind the server on the port and start serving requests
	log.Printf("starting the auctioneer http server at: %v: %v", ip, port)
	http.ListenAndServe(ip+":"+strconv.Itoa(port), mux)
	return nil
}

func registerRoutes(mux *http.ServeMux) {
	// TODO: decide the api path to reflect REST practices
	mux.HandleFunc("/register", registerBidder)
	mux.HandleFunc("/ask", processAuctionRequest)
}
