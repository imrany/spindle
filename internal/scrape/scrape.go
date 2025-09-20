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
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Links        []string `json:"links"`
	Favicon      string   `json:"favicon"`
	Images       []string `json:"images"`
	PreviewImage string   `json:"preview_image"`
	Video        string   `json:"video"`
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

	log.Printf("[INFO] Fetching URL: %s (lang=%s)", urlStr, langFormat)

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
				var name, property, content string
				for _, a := range n.Attr {
					switch strings.ToLower(a.Key) {
					case "name":
						name = a.Val
					case "property":
						property = a.Val
					case "content":
						content = a.Val
					}
				}

				// Description
				if strings.EqualFold(name, "description") {
					pageInfo.Description = content
				}

				// OpenGraph / Twitter images
				if strings.EqualFold(property, "og:image") ||
					strings.EqualFold(property, "og:image:url") ||
					strings.EqualFold(name, "twitter:image") ||
					strings.EqualFold(name, "twitter:image:src") {
					abs := resolveURL(urlStr, content)
					if abs != "" {
						pageInfo.Images = append(pageInfo.Images, abs)
						if pageInfo.PreviewImage == "" {
							pageInfo.PreviewImage = abs
						}
					}
				}

				// OpenGraph video
				if strings.EqualFold(property, "og:video") {
					pageInfo.Video = resolveURL(urlStr, content)
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
							if pageInfo.PreviewImage == "" {
								pageInfo.PreviewImage = absoluteURL
							}
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

	log.Printf("[INFO] Extracted: title=%q, desc=%q, images=%d, video=%q",
		pageInfo.Title, pageInfo.Description, len(pageInfo.Images), pageInfo.Video)

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
		log.Printf("[ERROR] Extracting info: %v", err)
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(pageInfo); err != nil {
		log.Printf("[ERROR] Encoding response: %v", err)
	}
}
