# Chirpy

A lightweight HTTP web server written in Go.

## 🚀 What This Project Does

**Chirpy** is a minimal HTTP server built using Go’s standard library. It provides a clean foundation for building APIs or web services without relying on external frameworks. The project focuses on clarity, simplicity, and extensibility.

Chirpy is intentionally small: it shows how routing, request handling, and server configuration work under the hood in Go.

## 🤔 Why You Should Care

- 🧠 **Learn Go web fundamentals** without abstraction-heavy frameworks
- ⚡ **Fast and lightweight** thanks to Go’s `net/http`
- 🧩 **Easy to extend** into a real API or service
- 🛠️ **Great starter repo** for personal projects, interviews, or coursework

If you want to understand *how Go web servers actually work*, Chirpy is a good place to start.

## 🛠️ Installation & Setup

### Prerequisites

- Go **1.18+** installed

### Clone the Repository

```bash
git clone https://github.com/jaharbaugh/chirpy.git
cd chirpy
```

### Run the Server

```bash
go run main.go
```

Or build a binary:

```bash
go build -o chirpy
./chirpy
```

By default, the server starts locally and listens for HTTP requests.

## 🌐 Example Endpoints

> (Exact behavior may vary depending on current implementation.)

### Health Check

```http
GET /health
```

**Response**
```json
{"status": "ok"}
```

### Example API Route

```http
GET /api/chirp
```

**Response**
```json
{
  "message": "chirp chirp"
}
```

You can test endpoints with `curl`:

```bash
curl http://localhost:8080/health
```

## ⚙️ Configuration

Chirpy can be configured using environment variables.

| Variable | Description | Default |
|--------|------------|---------|
| `PORT` | Port the server listens on | `8080` |

Example:

```bash
PORT=3000 go run main.go
```

## 📁 Project Structure

```
chirpy/
├── main.go        # Application entry point
├── handlers.go    # HTTP handlers
├── routes.go      # Route definitions
├── assets/        # Static files (if enabled)
└── README.md
```

## 🧱 Extending Chirpy

Here are a few easy ways to build on this project:

### ➕ Add a New Route

```go
http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, world!"))
})
```

### 🧩 Add Middleware

Wrap handlers to add logging, auth, or request timing.

### 🗄️ Add Persistence

Connect Chirpy to:
- SQLite or Postgres
- In-memory storage
- Redis

### 🧪 Add Tests

Use Go’s built-in testing tools:

```bash
go test ./...
```

## 📌 Goals of This Project

- Keep dependencies minimal
- Favor readability over cleverness
- Provide a solid base for experimentation

## 📜 License

MIT License

---

*Chirp proudly. Build simply.* 🐦

