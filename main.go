package main

import (
	"flag"
)

func main() {
	run := flag.String("run", "render", "Run process : crawl|render")
	flag.Parse()

	parseConfig()

	switch *run {
	case "crawl":
		crawl()
	case "render":
		render()
	case "server":
		newServer()
	}
}
