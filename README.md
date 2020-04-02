### last.friends

###### installation 

`$ go get -u go.senan.xyz/last-dot-friends/cmd/friends`

###### usage 

    $ friends -h
    * usage:
    friends [options] [username]...
    
    * api keys:
    https://www.last.fm/api/account/create
    
    * options:
      -api-key string
        	last.fm api key (required)
      -api-secret string
        	last.fm api secret (required)
      -interval int
        	time (in seconds) to wait between fetches (default 20)

###### api keys

get them [here](https://www.last.fm/api/account/create)

###### example with desktop notifications

```bash
friends \
    -interval 30 \
    -api-key "my-api-key" \
    -api-secret "my-api-secret" \
    alexkraak \
    bovineknight \
    devoxel \
    izaak \
    mortalslayer \
    sentriz \
| while IFS=$'\t' read -r username artist track; do
    notify-send --icon ~/downloads/lastfm.png "$username" "$artist — $track"
done
```
