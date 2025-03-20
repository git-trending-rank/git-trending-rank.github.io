package scheduler

import (
	"fmt"
	"github-trending-scraper/internal/config"
	"github-trending-scraper/internal/formatter"
	"github-trending-scraper/internal/lib"
	"github-trending-scraper/internal/scraper"
	"log"
	"os"
)

func Schedule(cfg config.Config) {

	repos, err := scraper.ScrapeGithubTrending(cfg.Frequency)
	if err != nil {
		log.Println("Error scraping:", err)
		return
	}

	markdown, err := formatter.FormatToMarkdown(repos, cfg.Frequency)
	if err != nil {
		log.Println("Error formatting:", err)
		return
	}

	filePath := fmt.Sprintf("%s/trending-%s-%s.md", cfg.OutputPath, cfg.Frequency, lib.PrintDateInfo(cfg.Frequency))
	err = os.WriteFile(filePath, []byte(markdown), 0644)
	if err != nil {
		log.Println("Error writing file:", err)
	}

}
