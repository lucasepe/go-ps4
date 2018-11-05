package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/gocolly/colly"
	"github.com/lucasepe/go-ps4/utils"
)

type Game struct {
	Title   string
	Cover   string
	Details string
	Price   string
}

var (
	version string
	date    string
)

func visit(rootUrl string) chan Game {
	ch := make(chan Game)

	go func() {
		// [\p{Sc}]

		collector := colly.NewCollector(
			colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:62.0) Gecko/20100101 Firefox/62.0"),
			colly.AllowedDomains("store.playstation.com"),
			//colly.Async(true),
			//colly.CacheDir("./playstation-store-cache/"),
		)

		collector.OnHTML("div[class*='grid-cell--game']", func(e *colly.HTMLElement) {
			cover := e.ChildAttr("div.grid-cell__thumbnail img[src^='https://store.playstation.com/store/api/chihiro']", "src")

			title := e.ChildText("div.grid-cell__body > a")
			details := e.ChildText("div.grid-cell__bottom > grid-cell__details-container grid-cell__left-detail--detail-2")
			price := e.ChildText("div.grid-cell__bottom > div.grid-cell__footer h3.price-display__price")

			if price != "" {
				ch <- Game{Title: title, Cover: cover, Details: details, Price: price}
			}
		})

		collector.OnHTML("div.grid-footer-controls a.paginator-control__next", func(e *colly.HTMLElement) {
			link := e.Request.AbsoluteURL(e.Attr("href"))
			e.Request.Visit(link)
		})

		collector.OnRequest(func(r *colly.Request) {
			//log.Println("visiting", r.URL.String())
		})

		collector.Visit(rootUrl)

		collector.Wait()

		close(ch)
	}()

	return ch
}

func showVersion() {
	fmt.Printf("PS4 Games CLI (build:%s@%s)\n", version, date)
}

/**
 * NIX: go build -ldflags "-X main.version=0.0.1 -X main.date=%date:~10,4%-%date:~4,2%-%date:~7,2%T%time:~0,2%:%time:~3,2%:%time:~6,2%"
 *
 * WIN: go build -ldflags "-X main.version=0.0.1 -X main.date=%date:~-4%-%date:~3,2%-%date:~6,2%T%time:~0,2%:%time:~3,2%:%time:~6,2%"
 */
func main() {

	weeklyDealsUrls := map[string]string{
		"it": "it-IT/grid/STORE-MSF75508-DOTW1",
		"en": "en-US/grid/STORE-MSF77008-WEEKLYDEALS",
	}

	searchUrls := map[string]string{
		"it": "it-it/search/",
		"en": "en-US/search/",
	}

	addonsUrls := map[string]string{
		"en": "en-US/grid/STORE-MSF77008-NEWPS4ADDONSCATE",
		"it": "it-it/grid/STORE-MSF75508-ADDONSSEEALL",
	}

	allGamesUrls := map[string]string{
		"en": "en-US/grid/STORE-MSF77008-PS4ALLGAMESCATEG",
		"it": "it-it/grid/STORE-MSF75508-PS4CAT",
	}

	optVersion := flag.Bool("version", false, "show app version")
	optLang := flag.String("lang", "it", "language (it, en)")
	optAddons := flag.Bool("addons", false, "show also extra contents")
	optFree := flag.Bool("free", false, "show only free titles")
	optWeeklyDeals := flag.Bool("weekly-deals", false, "show only weekly deals titles")
	optSearch := flag.String("search", "", "search for specified title")

	flag.Parse()

	if *optVersion {
		showVersion()
		os.Exit(0)
	}

	var uriPath string
	if *optAddons {
		uriPath = addonsUrls[*optLang]
	} else if *optSearch != "" {
		uriPath = searchUrls[*optLang]
	} else if *optWeeklyDeals {
		uriPath = weeklyDealsUrls[*optLang]
	} else {
		uriPath = allGamesUrls[*optLang]
	}

	u, err := url.Parse(fmt.Sprintf("https://store.playstation.com/%s", uriPath))
	if err != nil {
		log.Fatal(err)
	}

	if *optSearch == "" {
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

	} else {
		u, err = url.Parse(fmt.Sprintf("https://store.playstation.com/%s/%s", uriPath, *optSearch))
		if err != nil {
			log.Fatal(err)
		}
	}

	for g := range visit(u.String()) {
		fmt.Printf("%s %s\n", utils.RightPad(g.Title, 68, "."), utils.RightPad(g.Price, 7, " "))
	}
}
