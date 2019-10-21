package playstation

import (
	"github.com/gocolly/colly"
)

var (
	WeeklyDealsUrls = map[string]string{
		"it": "it-IT/grid/STORE-MSF75508-DOTW1",
		"en": "en-US/grid/STORE-MSF77008-WEEKLYDEALS",
		"at": "de-AT/grid/STORE-MSF75508-DOTW2/1",
		"de": "de-DE/grid/STORE-MSF75508-DOTW2/1",
	}

	SearchUrls = map[string]string{
		"it": "it-it/search/",
		"en": "en-US/search/",
		"at": "de-at/search/",
		"de": "de-de/search/",
	}

	AddonsUrls = map[string]string{
		"en": "en-US/grid/STORE-MSF77008-NEWPS4ADDONSCATE",
		"it": "it-it/grid/STORE-MSF75508-ADDONSSEEALL",
		"at": "de-at/grid/STORE-MSF75508-ADDONSSEEALL",
		"de": "de-de/grid/STORE-MSF75508-ADDONSSEEALL",
	}

	AllGamesUrls = map[string]string{
		"en": "en-US/grid/STORE-MSF77008-PS4ALLGAMESCATEG",
		"it": "it-it/grid/STORE-MSF75508-PS4CAT",
		"at": "de-at/grid/STORE-MSF75508-PS4CAT",
		"de": "de-de/grid/STORE-MSF75508-PS4CAT",
	}
)

type Game struct {
	Title   string
	Cover   string
	Details string
	Price   string
	Url     string
}

func Visit(rootUrl string) chan Game {
	ch := make(chan Game)

	go func() {
		// [\p{Sc}]

		collector := colly.NewCollector(
			colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:62.0) Gecko/20100101 Firefox/62.0"),
			colly.AllowedDomains("store.playstation.com"),
		)

		collector.OnHTML("div[class*='grid-cell--game']", func(e *colly.HTMLElement) {
			cover := e.ChildAttr("div.grid-cell__thumbnail img[src^='https://store.playstation.com/store/api/chihiro']", "src")

			title := e.ChildText("div.grid-cell__body > a")
			details := e.ChildText("div.grid-cell__bottom > grid-cell__details-container grid-cell__left-detail--detail-2")
			price := e.ChildText("div.grid-cell__bottom > div.grid-cell__footer h3.price-display__price")
			url := e.ChildAttr("div.grid-cell__body a", "href")

			if price != "" {
				ch <- Game{Title: title, Cover: cover, Details: details, Price: price, Url: url}
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
