# 🛒 Mercado Livre Discount Crawler (Go)

A simple Go-based command-line crawler for **Mercado Livre** that helps you search for discounts on products like "smart TV", "notebook", or anything else you're interested in. This project is intended as a **learning tool** for understanding the basics of Golang, including:

- HTTP requests
- HTML parsing
- CLI argument handling
- File I/O
- Basic error handling

> ✅ Right now, it works as a CLI utility. You run it with a product name and get up to 3 top results, their prices, discounts, and links.

---

## 🚀 Getting Started

### 📦 Requirements

Make sure you have Go installed. If not, [download and install it here](https://go.dev/doc/install).

> This project uses only Go’s standard library — no third-party packages!

Check your Go installation:

```bash
go version
````

---

### 🔧 Building & Running

#### Run directly (no build)

```bash
go run . smart tv
```

#### Build and run as an executable

```bash
go build -o discount-crawler-ml
./discount-crawler-ml smart tv
```

> You can replace `smart tv` with any product you'd like to search for.

---

## 🧠 What It Does

1. Constructs a search URL using your product name.
2. Sends an HTTP GET request to Mercado Livre.
3. Parses the HTML to extract:

   * Product title
   * Price
   * Discount (if available)
   * Product link
4. Prints the top 4 results to the terminal.
5. Saves the full HTML response to a local file for inspection/debugging.

---

## 📚 Example Output

```text
-----
Título: Smart TV 55" 4K Samsung
Preço: R$ 2.299,00
Desconto: 15%
Link: https://...
-----
Título: Smart TV LG 50"
Preço: R$ 2.099,00
Sem desconto
Link: https://...
```

---

## 🛠 Planned Features

This project is still evolving! Here's what’s coming up:

* [ ] 📨 **Turn it into a web server** that runs daily
* [ ] 📧 **Send daily emails** with best discounts
* [ ] 📊 **Store prices in a database** over time
* [ ] 🕵️ **Analyze discount patterns**
* [ ] 📁 Output results to **CSV or JSON**
