package main

import (
	"flag"
	"fmt"
)

// Custom function for logging
func msg(v ...interface{}) {
	if settings.Debug {
		fmt.Println(v)
	}
}

func main() {
	// Get command line flags
	crawlFlag := flag.Bool("crawl", false, "Crawl images from gopherize.me before rendering")
	serverFlag := flag.Bool("server", false, "Start in server mode")
	flag.Parse()

	// Parse yaml configuration
	err := parseSettings()
	if err != nil {
		fmt.Println("Can't parse settings.yml file")
		return
	}

	// Refresh images if needed
	if *crawlFlag {
		crawl()
	}

	// Start rendering
	if *serverFlag {
		newServer()
	} else {
		render()
	}
}
