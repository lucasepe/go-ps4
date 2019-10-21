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
	browser "github.com/pkg/browser"
)

var (
	version string
)

/**
 * NIX: CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static" -X main.version=0.5.1'
 *
 * WIN (powershell):
 * $env:CGO_ENABLED="0"
 * go build -ldflags "-X main.version=0.5.1"
 */

func main() {
	flag.Usage = func() {
		fmt.Printf("PS Store CLI (v.%s)\n", version)
		fmt.Printf("Usage: %s [OPTIONS] [Game Title]\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}

	optRegion := flag.String("region", "it", "region (it, en, at, de)") // should probably renamed to optCountry
	optAddons := flag.Bool("addons", false, "show also extra contents")
	optFree := flag.Bool("free", false, "show only free titles")
	optWeeklyDeals := flag.Bool("weekly-deals", false, "show only weekly deals titles")

	urls := []string{}

	flag.Parse()

	searchFor := len(flag.Args()) > 0 && strings.TrimSpace(flag.Args()[0]) != ""

	var uriPath string
	if *optAddons {
		uriPath = ps.AddonsUrls[*optRegion]
	} else if searchFor {
		uriPath = ps.SearchUrls[*optRegion]
	} else if *optWeeklyDeals {
		uriPath = ps.WeeklyDealsUrls[*optRegion]
	} else {
		uriPath = ps.AllGamesUrls[*optRegion]
	}

	u, err := url.Parse(fmt.Sprintf("https://store.playstation.com/%s",
		uriPath))
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
			} else if *optRegion == "it" {
				q.Add("price", "1000-1999,2000-2999,3000-3999")
			}
		}

		u.RawQuery = q.Encode()
	}

	urls = doSearch(u)
	prompt(urls)
}

func doSearch(u *url.URL) []string {
	urls := []string{}
	idx := 0
	for g := range ps.Visit(u.String()) {
		fmt.Printf("%s %s %s\n", padNumberWithSpaces(idx), utils.RightPad(g.Title, 66, "."), utils.RightPad(g.Price, 7, " "))
		urls = append(urls, g.Url)
		idx++
	}
	return urls
}

// https://stackoverflow.com/a/48089636/1548552
func padNumberWithSpaces(value int) string {
	return fmt.Sprintf("%3v", value)
}

func prompt(urls []string) {
	var idx int
	fmt.Println("")
	fmt.Println("Which game should be opened in the browser? Enter number")
	fmt.Scan(&idx)
	browser.OpenURL("https://store.playstation.com" + urls[idx])
}
