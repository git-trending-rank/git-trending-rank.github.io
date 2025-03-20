package formatter

import (
	"bytes"
	"github-trending-scraper/internal/lib"
	"text/template"
	"time"

	"github-trending-scraper/internal/scraper"
)

const markdownTemplate = `---
title: GitHub 趋势 {{.FrequencyDate}}
date: {{.Date}}
categories:
- {{.Frequency}}
---
<link rel="stylesheet" href="/public/css/trending.css">
{{"{{"}}< raw >{{"}}"}}
	<main class="container">
        <div class="repo-list" id="repoList">

	{{range $index, $repo := .Repos}}
			<div class="repo-card">
				<p><a href="{{$repo.URL}}" target="_blank">{{$repo.Name}}</a></p>
				<p>{{$repo.Description}}</p>
				<div class="repo-stats">
					<div>
						<span>🔠 {{$repo.Language}}</span>
						<span>⭐ {{$repo.Stars}}</span>
						<span>🔱 {{$repo.Forks}}</span>
					</div>
				<div class="stars-today">⭐ {{$repo.StarsToday}} stars {{$.RankCycle}}</div>
				</div>
			</div>
	{{end}}

		</div>
    </main>
{{"{{"}}< /raw >{{"}}"}}
`

func FormatToMarkdown(repos []scraper.RepoInfo, frequency string) (string, error) {
	funcMap := template.FuncMap{
		"add1": func(i int) int {
			return i + 1
		},
	}

	tmpl, err := template.New("markdown").Funcs(funcMap).Parse(markdownTemplate)
	if err != nil {
		return "", err
	}

	data := struct {
		Repos         []scraper.RepoInfo
		Frequency     string
		FrequencyDate string
		RankCycle     string
		Date          string
	}{
		Repos:         repos,
		Frequency:     frequency,
		FrequencyDate: lib.PrintDateInfo(frequency),
		Date:          time.Now().Format(time.RFC3339),
		RankCycle:     lib.PrintDateRankCycle(frequency),
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
