package main

import (
	"flag"
)

func main() {
	run := flag.String("run", "render", "Run process : crawl|render")
	flag.Parse() 

	switch *run {
	case "crawl":
		crawl()
	case "render":
		render()
	}
}