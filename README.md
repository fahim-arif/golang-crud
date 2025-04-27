## 🚀 Project overview
This repo is a playground for building **CRUD APIs in Go**.  
It uses **Gin** for routing, **GORM** for data access, and **PostgreSQL 17.4** running in Docker.  
A multi-stage Dockerfile keeps the image tiny, while **Docker Compose** spins the whole stack up with a single command and shares a `.env` file so your Go code and the database always agree on credentials.

## 📦 Tech stack

| Layer | Tool | Notes |
|-------|------|-------|
| HTTP router | **Gin v1.10** | Reads `PORT` if you call `router.Run()` with no args. |
| ORM | **GORM v1.26** + `gorm.io/driver/postgres` | DSN pattern: `host=… user=… password=… dbname=… port=… sslmode=… TimeZone=…`. |
| DB | **PostgreSQL 17.4** official image | Needs `POSTGRES_PASSWORD` (or `trust` mode) on first boot. |
| Dev env | **Docker Compose v2** | `env_file`, health‑checks, and `depends_on: condition: service_healthy`. |
| Hot reload | **CompileDaemon** | Rebuilds & restarts on file change. |
| Config | **godotenv** | Loads `.env` into `os.Getenv` in `init()`. |

## 🛠 Prerequisites

* Go 1.24+  
* Docker & Docker Compose (v2)  
* Make (optional, for shortcuts)

## ⚡️ Quick‑start

```bash
# 1. Clone and cd
git clone https://github.com/your-user/crud-basic.git
cd crud-basic

# 2. Copy sample environment
cp .env.example .env        # tweak if you like

# 3. One‑command boot
docker compose up --build   # builds API image + starts db

# 4. Hit the health‑check
curl http://localhost:3000/ping   # {"message":"pong"}


# PostgreSQL
DB_USER=gorm
DB_PASSWORD=gorm
DB_NAME=gorm
DB_HOST=db
DB_PORT=5432
DB_SSLMODE=disable
DB_TIMEZONE=Asia/Shanghai

environment:
  POSTGRES_USER:     ${DB_USER}
  POSTGRES_PASSWORD: ${DB_PASSWORD}
  POSTGRES_DB:       ${DB_NAME}


crud-basic/
├── cmd/
├── initializers/
├── models/
├── handlers/
├── Dockerfile
├── compose.yaml
├── .env.example
└── README.md

CompileDaemon -build "go build -o server ." -command "./server"


```

## 🧪 Testing
Unit tests: go test ./...

## Lint: golangci-lint run

API smoke test: curl -X GET :3000/ping

## 🚚 Deployment notes
Build once, run anywhere: the final Alpine image is <20 MB.

Set GIN_MODE=release in prod.

If using an external Postgres, change DB_HOST and remove the db: service.

## 🙏 Credits
Inspired by examples in the Gin, GORM, Docker and PostgreSQL communities.
