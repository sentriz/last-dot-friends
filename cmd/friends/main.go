package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"senan.xyz/g/last-dot-friends/watch"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "* usage:\n")
		fmt.Fprintf(os.Stderr, "%s [options] [username]...\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "* api keys:\n")
		fmt.Fprintf(os.Stderr, "https://www.last.fm/api/account/create\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "* options:\n")
		flag.PrintDefaults()
	}
	interval := flag.Int("interval", 20, "time (in seconds) to wait between fetches")
	apiKey := flag.String("api-key", "", "last.fm api key (required)")
	apiSecret := flag.String("api-secret", "", "last.fm api secret (required)")
	flag.Parse()
	if *apiSecret == "" || *apiKey == "" {
		fmt.Fprintln(os.Stderr, "please provide both a last.fm api key and secret. see -h")
		os.Exit(1)
	}
	usernames := flag.Args()
	if len(usernames) == 0 {
		fmt.Fprintln(os.Stderr, "please provide usernames. see -h")
		os.Exit(1)
	}
	w := watch.NewWatcher(
		usernames,
		*apiKey,
		*apiSecret,
		time.Duration(*interval)*time.Second,
	)
	events := make(chan watch.Event)
	go w.Watch(events)
	for {
		event := <-events
		fmt.Printf("%s\t%s\t%s\n",
			event.Username,
			event.Artist,
			event.Track,
		)
	}
}
