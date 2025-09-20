package scrape

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)


type PageInfo struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Links       []string `json:"links"`
	Favicon     string   `json:"favicon"`
	Images      []string `json:"images"`
	Video       string   `json:"video"` //URL of the first found video source, for simplicity

}

func ExtractInfo(urlStr string) (PageInfo, error) {
	var pageInfo PageInfo
	pageInfo.Links = make([]string, 0)
	pageInfo.Images = make([]string, 0)

	resp, err := http.Get(urlStr)
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
			if n.Data == "title" {
				if n.FirstChild != nil {
					pageInfo.Title = n.FirstChild.Data
				}
			} else if n.Data == "meta" {
				name := ""
				content := ""
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
			} else if n.Data == "a" {
				for _, a := range n.Attr {
					if a.Key == "href" {
						absoluteURL := resolveURL(urlStr, a.Val)
						if absoluteURL != "" {
							pageInfo.Links = append(pageInfo.Links, absoluteURL)
						}
						break
					}
				}
			} else if n.Data == "link" {
				rel := ""
				href := ""

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
			} else if n.Data == "img" {
				for _, a := range n.Attr {
					if a.Key == "src" {
						absoluteURL := resolveURL(urlStr, a.Val)
						if absoluteURL != "" {
							pageInfo.Images = append(pageInfo.Images, absoluteURL)
						}
						break
					}
				}
			} else if n.Data == "video" { //Simplified video extraction
				for _, a := range n.Attr {
					if a.Key == "src" {
						pageInfo.Video = resolveURL(urlStr, a.Val)
						break //Just taking the first video src, if present as an attribute
					}
				}

				if pageInfo.Video == "" { //Check for source tags within the video tag
					for c := n.FirstChild; c != nil; c = c.NextSibling {
						if c.Type == html.ElementNode && c.Data == "source" {
							for _, a := range c.Attr {
								if a.Key == "src" {
									pageInfo.Video = resolveURL(urlStr, a.Val)
									break // First video source is enough
								}
							}
							if pageInfo.Video != "" {
								break // stop iterating source tags
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

// resolveURL: Same as before
func resolveURL(baseURL string, relativeURL string) string {
	base, err := url.Parse(baseURL)
	if err != nil {
		return ""
	}

	rel, err := url.Parse(relativeURL)
	if err != nil {
		return ""
	}

	absoluteURL := base.ResolveReference(rel).String()
	return absoluteURL
}

func ScrapeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	urlParam := r.URL.Query().Get("url")
	if urlParam == "" {
		http.Error(w, "URL parameter is required", http.StatusBadRequest)
		return
	}

	pageInfo, err := ExtractInfo(urlParam)
	if err != nil {
		log.Printf("Error extracting info: %v", err)
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(pageInfo)
}