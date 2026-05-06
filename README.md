# 🔐 CyberSec Portfolio

A **cyberpunk-aesthetic personal portfolio website** for a Cybersecurity Specialist. Built with pure Go (no external dependencies) using embedded static files.

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![Security](https://img.shields.io/badge/Focus-Cybersecurity-00ff88?style=flat)
![License](https://img.shields.io/badge/License-MIT-green?style=flat)

## ✨ Features

- ⚡ **Zero external dependencies** — pure Go stdlib
- 🎨 **Cyberpunk terminal aesthetic** — matrix rain, scan lines, glitch effects
- 📱 **Fully responsive** — works on all screen sizes
- 🔒 **Single binary deployment** — static files embedded via `//go:embed`
- 🌐 **GitHub Pages compatible** — also serves as static HTML

## 🚀 Quick Start

### Run locally
```bash
git clone https://github.com/yourusername/cybersec-portfolio
cd cybersec-portfolio
go run main.go
# Open http://localhost:8080
```

### Build binary
```bash
go build -o portfolio main.go
./portfolio
```

### Deploy as static site (GitHub Pages)
Just upload `static/index.html` — it works as a standalone HTML file too.

## 🛠️ Customization

Edit `static/index.html` to update:
- **Your name** — find `CYBER<span class="accent">SEC</span>` in the hero section
- **Contact links** — update GitHub, LinkedIn, email in `#contact` section
- **Certifications** — update the `#certifications` section
- **Skills** — modify cards in `#skills` section

## 📦 Project Structure

```
.
├── main.go          # Go web server (embeds static/)
├── go.mod           # Go module definition
├── static/
│   └── index.html   # Full portfolio (self-contained)
└── README.md
```

## 🎯 Tech Stack

- **Backend**: Go 1.21+ (net/http, embed)
- **Frontend**: Vanilla HTML/CSS/JS
- **Fonts**: Orbitron, Share Tech Mono, Rajdhani (Google Fonts)
- **Design**: Custom cyberpunk/terminal aesthetic

## 📄 License

MIT — feel free to use and modify.
