package scraper

import (
	"fmt"
	"github-trending-scraper/internal/lib"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type RepoInfo struct {
	Name        string
	URL         string
	Description string
	Language    string
	Stars       int
	Forks       int
	StarsToday  int
}

func ScrapeGithubTrending(frequency string) ([]RepoInfo, error) {
	url := fmt.Sprintf("https://github.com/trending?since=%s", frequency)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var repos []RepoInfo
	doc.Find("article.Box-row").Each(func(i int, s *goquery.Selection) {
		repo := RepoInfo{}

		repo.Name = s.Find("h2 a").Text()
		repo.URL, _ = s.Find("h2 a").Attr("href")
		repo.URL = "https://github.com" + repo.URL
		repo.Description = s.Find("p").Text()

		languageSpan := s.Find("span[itemprop='programmingLanguage']")
		repo.Language = languageSpan.Text()

		starsText := s.Find("a[href$=stargazers]").Text()
		starsText = lib.NoWhitespace(strings.ReplaceAll(starsText, ",", ""))
		repo.Stars, _ = strconv.Atoi(starsText)

		forksText := s.Find("a[href$=forks]").Text()
		forksText = lib.NoWhitespace(strings.ReplaceAll(forksText, ",", ""))
		repo.Forks, _ = strconv.Atoi(forksText)

		starsTodayText := s.Find("span.d-inline-block.float-sm-right").Text()
		starsTodayText = strings.ReplaceAll(starsTodayText, ",", "")
		starsTodayText = strings.Fields(starsTodayText)[0] // "1,234 stars today" -> "1234"
		repo.StarsToday, _ = strconv.Atoi(starsTodayText)

		repos = append(repos, repo)
	})

	return repos, nil
}
