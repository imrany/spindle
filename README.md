# 🕸️ Spindle

**Spindle** is an open-source, lightweight **web crawler and scraper**.
It can discover links on the web (_crawler_) and extract structured data from webpages (_scraper_).

## ✨ Purpose

- **Crawler** → Navigates pages and discovers new URLs.
- **Scraper** → Extracts specific information from a given page (title, description, favicon, links, images, and videos).

Together: Spindle explores **what to scrape** and **extracts the data you care about**.

## ⚙️ How It Works

1. Takes an input URL (from CLI or API).
2. Fetches the HTML.
3. Extracts structured data:

   - Title
   - Description
   - Links
   - Favicon
   - Images
   - Videos (if available)

4. In crawler mode, follows links to discover additional pages.

> **GitHub Container Registry**: available at `ghcr.io/imrany/spindle`

### Run on docker

```bash
## Pull the docker image
docker pull ghcr.io/imrany/spindle

## Runs the image, creating a spindle container
docker run -d --name spindle --restart unless-stopped -p 5020:5020 -v ~/.spindle:/var/opt/spindle ghcr.io/imrany/spindle server
```

## 📦 Build

```bash
# Clone the repository
git clone https://github.com/imrany/spindle.git
cd spindle

# Install Go dependencies
go mod download

# Build
go build main.go
```

## 🚀 Usage

### 🔹 Scrape URL in CLI

```bash
go run main.go https://www.youtube.com/watch?v=pum3k4yECT4
```

**Response (truncated for readability):**

```bash
Title: America Is In Trouble.. Candace Owens Might Be Cooked & Zuck Got Massively Embarrassed! - YouTube
Description: THIS WEEK ON NEWSDADDYYYY!!! 🥤🍿**JIMMY KIMMEL — ABC PULLS THE PLUG**Jimmy Kimmel’s late-night show was pulled from the schedule after his comments about Ch...
Favicon: https://www.youtube.com/s/desktop/2ea5cbbe/img/favicon_144x144.png
Video:
Links: [https://www.youtube.com/ https://www.youtube.com/ https://www.youtube.com/about/ https://www.youtube.com/about/press/ https://www.youtube.com/about/copyright/ https://www.youtube.com/t/contact_us/ https://www.youtube.com/creators/ https://www.youtube.com/ads/ https://developers.google.com/youtube https://www.youtube.com/t/terms https://www.youtube.com/t/privacy https://www.youtube.com/about/policies/ https://www.youtube.com/howyoutubeworks?utm_campaign=ytgen&utm_source=ythp&utm_medium=LeftNav&utm_content=txt&u=https%3A%2F%2Fwww.youtube.com%2Fhowyoutubeworks%3Futm_source%3Dythp%26utm_medium%3DLeftNav%26utm_campaign%3Dytgen https://www.youtube.com/new]
Images: []
```

### 🔹 Run in Server Mode

Start the server (defaults: `0.0.0.0:5020`):

```bash
go run main.go server --addr=0.0.0.0 --port=5020
```

Test with `curl` or browser:

```bash
curl "http://localhost:5020/scrape?url=https://www.youtube.com/watch?v=pum3k4yECT4"
```

**JSON Response:**

```json
{
  "title": "America Is In Trouble.. Candace Owens Might Be Cooked \u0026 Zuck Got Massively Embarrassed! - YouTube",
  "description": "THIS WEEK ON NEWSDADDYYYY!!! 🥤🍿**JIMMY KIMMEL — ABC PULLS THE PLUG**Jimmy Kimmel’s late-night show was pulled from the schedule after his comments about Ch...",
  "links": [
    "https://www.youtube.com/",
    "https://www.youtube.com/",
    "https://www.youtube.com/about/",
    "https://www.youtube.com/about/press/",
    "https://www.youtube.com/about/copyright/",
    "https://www.youtube.com/t/contact_us/",
    "https://www.youtube.com/creators/",
    "https://www.youtube.com/ads/",
    "https://developers.google.com/youtube",
    "https://www.youtube.com/t/terms",
    "https://www.youtube.com/t/privacy",
    "https://www.youtube.com/about/policies/",
    "https://www.youtube.com/howyoutubeworks?utm_campaign=ytgen\u0026utm_source=ythp\u0026utm_medium=LeftNav\u0026utm_content=txt\u0026u=https%3A%2F%2Fwww.youtube.com%2Fhowyoutubeworks%3Futm_source%3Dythp%26utm_medium%3DLeftNav%26utm_campaign%3Dytgen",
    "https://www.youtube.com/new"
  ],
  "favicon": "https://www.youtube.com/s/desktop/2ea5cbbe/img/favicon_144x144.png",
  "images": [],
  "video": ""
}
```

## 🔍 Features

- CLI and API modes.
- Extracts metadata (title, description, favicon, images, videos).
- Lightweight crawler for link discovery.
- JSON API for integration into other services.

## 📖 Example Use Cases

- Preview cards for links in chat apps.
- SEO or content analysis.
- Building your own search index.
- Research & data mining.

## 🗺️ Roadmap

- [ ] Respect `robots.txt` for crawler.
- [ ] Add caching & rate limiting.
- [ ] Support deeper recursive crawling.
- [ ] Extract Open Graph / Twitter Card metadata.

## Contributing

Spindle is an open-source project that welcomes contributions from developers, designers, and users worldwide. We maintain a collaborative and inclusive development environment that values quality, innovation, and community feedback.

### Ways to Contribute

- **Code Contributions**: Bug fixes, feature implementations, and performance improvements
- **Documentation**: API documentation, user guides, and technical specifications
- **Testing**: Quality assurance, test case development, and bug reporting
- **Localization**: Translation support for multiple languages and regions
- **Community Support**: Helping users on GitHub discussions, and forums

## License

Spindle is released under the MIT License, providing maximum flexibility for both personal and commercial use. This license allows for:

- **Commercial Use**: Deploy Spindle in commercial environments without licensing fees
- **Modification**: Adapt and customize the codebase for specific requirements
- **Distribution**: Share modified versions while maintaining license attribution
- **Private Use**: Use Spindle internally without disclosure requirements

See the [LICENSE](./LICENSE) file for complete licensing terms.
