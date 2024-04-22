package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/shkh/lastfm-go/lastfm"
)

func main() {
	interval := flag.Int("interval", 20, "time (in seconds) to wait between fetches")
	apiKey := flag.String("api-key", "", "last.fm api key (required)")
	apiSecret := flag.String("api-secret", "", "last.fm api secret (required)")
	flag.Parse()

	if *apiSecret == "" || *apiKey == "" {
		log.Fatal("please provide both a last.fm api key and secret. see -h")
	}

	usernames := flag.Args()
	if len(usernames) == 0 {
		log.Fatal("please provide usernames. see -h")
	}

	api := lastfm.New(*apiKey, *apiSecret)

	lastEvents := map[string]string{}
	for {
		for _, username := range usernames {
			time.Sleep(time.Duration(float32(*interval)/float32(len(lastEvents))) * time.Second)

			username, artist, track, err := fetchUser(api, username)
			if err != nil {
				log.Printf("error fetching user: %v", err)
				continue
			}
			if artist == "" {
				continue
			}

			k := artist + track
			if lastEvents[username] == k {
				continue
			}
			lastEvents[username] = k

			fmt.Printf("%s\t%s\t%s\n", username, artist, track)
		}
	}
}

func fetchUser(api *lastfm.Api, name string) (string, string, string, error) {
	r, err := api.User.GetRecentTracks(lastfm.P{"user": name, "limit": 1})
	if err != nil {
		return "", "", "", err
	}
	if len(r.Tracks) == 0 || r.Tracks[0].NowPlaying != "true" {
		return "", "", "", nil
	}
	return r.User, r.Tracks[0].Artist.Name, r.Tracks[0].Name, nil
}
