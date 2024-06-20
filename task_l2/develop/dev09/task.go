package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"golang.org/x/net/html"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Config struct {
	URL       string
	Recursive bool
	Depth     int
	OutputDir string
}

func parseFlags() (Config, error) {
	config := Config{}
	flag.StringVar(&config.URL, "url", "", "URL to download")
	flag.BoolVar(&config.Recursive, "r", false, "Recursive download")
	flag.IntVar(&config.Depth, "l", 1, "Download depth")
	flag.StringVar(&config.OutputDir, "o", ".", "Output directory")
	flag.Parse()

	if config.URL == "" {
		return config, fmt.Errorf("invalid url")
	}
	return config, nil
}

func downloadPage(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func savePage(outputDir string, urlToSave string, body []byte) error {
	u, err := url.Parse(urlToSave)
	if err != nil {
		return err
	}

	saveDir := path.Join(outputDir, u.Host, u.Path)
	if strings.HasSuffix(urlToSave, "/") {
		saveDir = path.Join(saveDir, "index.html")
	} else if !strings.Contains(path.Base(saveDir), ".") {
		saveDir += ".html"
	}

	err = os.MkdirAll(path.Dir(saveDir), 0755)
	if err != nil {
		return err
	}

	return os.WriteFile(saveDir, body, 0644)
}

func extractLinks(baseURL string, body []byte) ([]string, error) {
	var links []string
	tokenizer := html.NewTokenizer(strings.NewReader(string(body)))
	for {
		tt := tokenizer.Next()
		switch tt {
		case html.ErrorToken:
			return links, nil
		case html.StartTagToken, html.SelfClosingTagToken:
			t := tokenizer.Token()
			switch t.Data {
			case "a", "link", "script", "img":
				for _, attr := range t.Attr {
					if (t.Data == "a" && attr.Key == "href") ||
						(t.Data == "link" && attr.Key == "href" && attr.Val != "" && (strings.HasSuffix(attr.Val, ".css") || strings.HasPrefix(attr.Val, "http"))) ||
						(t.Data == "script" && attr.Key == "src" && attr.Val != "" && (strings.HasSuffix(attr.Val, ".js") || strings.HasPrefix(attr.Val, "http"))) ||
						(t.Data == "img" && attr.Key == "src") {
						link := attr.Val
						absoluteURL := resolveURL(baseURL, link)
						links = append(links, absoluteURL)
					}
				}
			}
		}
	}
}

func resolveURL(baseURL string, href string) string {
	base, err := url.Parse(baseURL)
	if err != nil {
		return href
	}
	ref, err := url.Parse(href)
	if err != nil {
		return href
	}
	return base.ResolveReference(ref).String()
}

func shouldVisit(url string, visited map[string]bool) bool {
	if _, ok := visited[url]; ok {
		return false
	}
	return true
}

func downloadSite(config Config, currentDepth int, visited map[string]bool) error {
	if currentDepth > config.Depth {
		return nil
	}

	body, err := downloadPage(config.URL)
	if err != nil {
		return err
	}

	err = savePage(config.OutputDir, config.URL, body)
	if err != nil {
		return err
	}

	visited[config.URL] = true

	if config.Recursive && currentDepth < config.Depth {
		links, err := extractLinks(config.URL, body)
		if err != nil {
			return err
		}

		for _, link := range links {
			if shouldVisit(link, visited) {
				newConfig := config
				newConfig.URL = link
				err := downloadSite(newConfig, currentDepth+1, visited)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func main() {
	config, err := parseFlags()
	if err != nil {
		log.Fatal(err)
	}
	visited := make(map[string]bool)
	downloadSite(config, 0, visited)
}
