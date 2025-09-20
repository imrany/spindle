package scrape

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/imrany/spindle/internal/localization"
	"golang.org/x/net/html"
)


type PageInfo struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Links       []string `json:"links"`
	Favicon     string   `json:"favicon"`
	Images      []string `json:"images"`
	Video       string   `json:"video"`
}

func ExtractInfo(urlStr string, lang string) (PageInfo, error) {
	var pageInfo PageInfo
	pageInfo.Links = make([]string, 0)
	pageInfo.Images = make([]string, 0)

	client := &http.Client{}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return pageInfo, fmt.Errorf("error creating request: %w", err)
	}

	// Pick correct Accept-Language
	langFormat := "en-US,en;q=0.9" // default fallback
	if val, ok := localization.LangMap[strings.ToLower(lang)]; ok {
		langFormat = val
	}
	req.Header.Set("Accept-Language", langFormat)

	// User-Agent
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; SpindleBot/1.0; +https://spindle.villebiz.com)")

	resp, err := client.Do(req)
	if err != nil {
		return pageInfo, fmt.Errorf("error fetching URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return pageInfo, fmt.Errorf("HTTP status code: %d", resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return pageInfo, fmt.Errorf("error parsing HTML: %w", err)
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "title":
				if n.FirstChild != nil {
					pageInfo.Title = n.FirstChild.Data
				}
			case "meta":
				var name, content string
				for _, a := range n.Attr {
					if a.Key == "name" {
						name = a.Val
					}
					if a.Key == "content" {
						content = a.Val
					}
				}
				if strings.ToLower(name) == "description" {
					pageInfo.Description = content
				}
			case "a":
				for _, a := range n.Attr {
					if a.Key == "href" {
						absoluteURL := resolveURL(urlStr, a.Val)
						if absoluteURL != "" {
							pageInfo.Links = append(pageInfo.Links, absoluteURL)
						}
						break
					}
				}
			case "link":
				var rel, href string
				for _, a := range n.Attr {
					if a.Key == "rel" {
						rel = a.Val
					}
					if a.Key == "href" {
						href = a.Val
					}
				}
				if strings.ToLower(rel) == "icon" || strings.ToLower(rel) == "shortcut icon" {
					pageInfo.Favicon = resolveURL(urlStr, href)
				}
			case "img":
				for _, a := range n.Attr {
					if a.Key == "src" {
						absoluteURL := resolveURL(urlStr, a.Val)
						if absoluteURL != "" {
							pageInfo.Images = append(pageInfo.Images, absoluteURL)
						}
						break
					}
				}
			case "video":
				for _, a := range n.Attr {
					if a.Key == "src" {
						pageInfo.Video = resolveURL(urlStr, a.Val)
						break
					}
				}
				if pageInfo.Video == "" {
					for c := n.FirstChild; c != nil; c = c.NextSibling {
						if c.Type == html.ElementNode && c.Data == "source" {
							for _, a := range c.Attr {
								if a.Key == "src" {
									pageInfo.Video = resolveURL(urlStr, a.Val)
									break
								}
							}
							if pageInfo.Video != "" {
								break
							}
						}
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return pageInfo, nil
}

// resolveURL stays unchanged
func resolveURL(baseURL string, relativeURL string) string {
	base, err := url.Parse(baseURL)
	if err != nil {
		return ""
	}
	rel, err := url.Parse(relativeURL)
	if err != nil {
		return ""
	}
	return base.ResolveReference(rel).String()
}

func ScrapeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	urlParam := r.URL.Query().Get("url")
	langParam := r.URL.Query().Get("lang")

	if urlParam == "" {
		http.Error(w, "URL parameter is required", http.StatusBadRequest)
		return
	}

	pageInfo, err := ExtractInfo(urlParam, langParam)
	if err != nil {
		log.Printf("Error extracting info: %v", err)
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(pageInfo)
}