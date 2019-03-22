package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	ps "github.com/lucasepe/playstation"
	"github.com/lucasepe/playstation/utils"
)

var (
	version string
)

/**
 * NIX: CGO_ENABLED=0 GOOS=linux go build -o ps4-games.exe -a -ldflags '-extldflags "-static" -X main.version=0.5.0'
 *
 * WIN (powershell):
 * $env:CGO_ENABLED="0"
 * go build -o ps4-games -ldflags "-X main.version=0.5.0"
 */
func main() {

	flag.Usage = func() {
		fmt.Printf("PS Store CLI (v.%s)\n", version)
		fmt.Printf("Usage: %s [OPTIONS] [Game Title]\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}

	optLang := flag.String("lang", "it", "language (it, en)")
	optAddons := flag.Bool("addons", false, "show also extra contents")
	optFree := flag.Bool("free", false, "show only free titles")
	optWeeklyDeals := flag.Bool("weekly-deals", false, "show only weekly deals titles")

	flag.Parse()

	searchFor := len(flag.Args()) > 0 && strings.TrimSpace(flag.Args()[0]) != ""

	var uriPath string
	if *optAddons {
		uriPath = ps.AddonsUrls[*optLang]
	} else if searchFor {
		uriPath = ps.SearchUrls[*optLang]
	} else if *optWeeklyDeals {
		uriPath = ps.WeeklyDealsUrls[*optLang]
	} else {
		uriPath = ps.AllGamesUrls[*optLang]
	}

	u, err := url.Parse(fmt.Sprintf("https://store.playstation.com/%s", uriPath))
	if err != nil {
		log.Fatal(err)
	}

	if searchFor {
		what := strings.TrimSpace(flag.Args()[0])
		u, err = url.Parse(fmt.Sprintf("https://store.playstation.com/%s/%s", uriPath, what))
		if err != nil {
			log.Fatal(err)
		}

	} else {
		q := u.Query()
		q.Add("direction", "asc")
		q.Add("sort", "price")
		q.Add("platform", "ps4")

		if !*optAddons && !*optWeeklyDeals {
			q.Add("gameContentType", "games,bundles")
		}

		if !*optWeeklyDeals {
			if *optFree {
				q.Add("price", "0-0")
			} else if *optLang == "it" {
				q.Add("price", "1000-1999,2000-2999,3000-3999")
			}
		}

		u.RawQuery = q.Encode()
	}

	doSearch(u)
}

func doSearch(u *url.URL) {
	for g := range ps.Visit(u.String()) {
		fmt.Printf("%s %s\n", utils.RightPad(g.Title, 68, "."), utils.RightPad(g.Price, 7, " "))
	}
}
