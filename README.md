<p align="center">
    <img src="https://github.com/user-attachments/assets/fffae98a-a175-495b-8066-cce4e536929b" align="center" width="30%">
</p>
<p align="center"><h1 align="center">Crawli</h1></p>
<p align="center">
    <em><code>❯ SEO Analytics tool for analyzing a website's internal linking profile</code></em>
</p>
<p align="center">
    <img src="https://img.shields.io/github/license/abhinaenae/crawli?style=default&logo=opensourceinitiative&logoColor=white&color=28a745" alt="license">
    <img src="https://img.shields.io/github/last-commit/abhinaenae/crawli?style=default&logo=git&logoColor=white&color=28a745" alt="last-commit">
    <img src="https://img.shields.io/github/languages/top/abhinaenae/crawli?style=default&color=28a745" alt="repo-top-language">
    <img src="https://img.shields.io/github/languages/count/abhinaenae/crawli?style=default&color=28a745" alt="repo-language-count">
</p>

---

## 🚀 Overview

Crawli is a fast, concurrent CLI web crawler designed to analyze a website’s internal link structure. It recursively follows links, tracking unique URLs and their occurrences, with user-defined options for concurrency and recursion depth to optimize performance and control crawl scope.

---

## ✨ Features
- Multithreaded crawling: Uses concurrency for efficient link discovery.
- Customizable depth & concurrency: Control the number of pages to crawl and the number of concurrent requests.
- Formatted reports: Generates a neatly structured table summarizing the crawled URLs.
---

## 🛠 Getting Started

### 📋 Prerequisites

Before getting started with Crawli, ensure your runtime environment meets the following requirements:

- **Programming Language:** Go 1.23
- **Package Manager:** Go modules

### 📥 Installation

**Build from source:**

1. Clone the crawli repository:
```sh
❯ git clone https://github.com/abhinaenae/crawli
```

2. Navigate to the project directory:
```sh
❯ cd crawli
```

3. Install the project dependencies:

**Using `go modules`**
```sh
❯ go build -o crawli ./cmd/main.go
```

### Usage

Run crawli using the following command:
```sh
❯ ./crawli <baseURL> <maxConcurrency> <maxPages>
```

### Example:

![image](https://github.com/user-attachments/assets/aafc45bb-b8ae-45cd-8919-149f3506bb1f)

---
## 🤝 Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to improve the project.

## 📜 License

This project is licensed under the MIT License.
