package watch

import (
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/shkh/lastfm-go/lastfm"
)

type Event struct {
	Username string
	Artist   string
	Track    string
}

type Watcher struct {
	api        *lastfm.Api
	interval   time.Duration
	lastEvents map[string]string
}

func NewWatcher(names []string, key string, secret string, interval time.Duration) *Watcher {
	w := &Watcher{
		api:        lastfm.New(key, secret),
		interval:   interval,
		lastEvents: make(map[string]string, len(names)),
	}
	for _, name := range names {
		w.lastEvents[name] = ""
	}
	return w
}

func (w *Watcher) fetchUser(name string) (Event, error) {
	result, err := w.api.User.GetRecentTracks(lastfm.P{
		"user":  name,
		"limit": 1,
	})
	if err != nil {
		return Event{}, errors.Wrap(err, "fetching from last.fm")
	}
	if len(result.Tracks) == 0 {
		return Event{}, fmt.Errorf("user %q has no recent tracks", name)
	}
	track := result.Tracks[0]
	if track.NowPlaying != "true" {
		return Event{}, fmt.Errorf("user %q is not currently listening", name)
	}
	return Event{
		Artist: track.Artist.Name,
		Track:  track.Name,
	}, nil
}

func (w *Watcher) Watch(ch chan<- Event) {
	for {
		for name := range w.lastEvents {
			event, err := w.fetchUser(name)
			if err != nil {
				log.Printf("error fetching user: %v\n", err)
				continue
			}
			eventID := event.Artist + event.Track
			if eventID == w.lastEvents[name] {
				continue
			}
			event.Username = name
			ch <- event
			w.lastEvents[name] = eventID
		}
		time.Sleep(w.interval)
	}
}
