# 🦊 Mobilfox – Take-home task

**Goal:** build a small, CLI backend in Go that manages product stock in PostgreSQL, runs via Docker Compose, and is safe under concurrency.

**Timebox:** ~6–8 hours.  

⚠️ *No frontend work is expected.*

---

## 💾 How to submit

1. **Fork** this repository to your own GitHub account (we can’t accept direct pushes to the original repo).
2. Implement the task in your fork.
3. Make sure everything runs via Docker Compose as described below.
4. Send us the link to your fork (and optionally a short note about tradeoffs you made).

---

## ⚙️ Domain

The database contains a small list of our products (SKU, name, current stock).  
Stock changes happen via CLI commands (increase/decrease).

You will find a data/products.csv file in this repository for sample data.  
*You do not need to parse the CSV — use it only as a source for testing*

---

## *️⃣ Requirements

### Tech constraints

- Golang
- PostgreSQL
- Docker + Docker Compose
- CLI only (no HTTP/gRPC API required)

### Commands to implement

Your program must implement the following commands:

#### 1) list

Print all products in a readable format:
- SKU
- name
- current stock

Example:

    docker compose run --rm app list

#### 2) increase

Increase stock for a SKU.

Required flags:
- --id (idempotency key, required)
- --sku (required)
- --qty (required, must be > 0)

Optional:
- --reason

Example:

    docker compose run --rm app increase --id req-001 --sku M-PSH-IP15PL-FG-000 --qty 10 --reason production

#### 3) decrease

Decrease stock for a SKU.

Required flags:
- --id (idempotency key, required)
- --sku (required)
- --qty (required, must be > 0)

Optional:
- --reason

Example:

    docker compose run --rm app decrease --id req-002 --sku M-PSH-IP15PL-FG-000 --qty 2 --reason order

Rules:
- Stock must never become negative.
- If the decrease would result in negative stock, reject the command with a non-zero exit code and a clear error message.

#### 4) report

Print a simple report:
- total SKU count
- total stock units
- top N products by stock
- low stock list (<= threshold)

Flags:
- --top (default 5)
- --low-stock (default 10)

Example:

    docker compose run --rm app report --low-stock 10 --top 5

---

## 🧩 Data & database requirements

### Postgres schema

The schema is up to you and your implementation, however based on the sample data you might end up with a table like this:

`products`
- sku (primary key)
- name
- stock (must be >= 0)
- timestamps are welcome :)

### Migrations

- Use SQL migrations committed in the repo.
- Your app should run migrations automatically on startup (or provide a command that runs them).
- Use data/products.csv as the source to create the initial INSERT seed in migrations or you can import it on you own way (with some postgres client directly)
- No need to parse CSV in Go

---

## ✅ Correctness requirements

### 1) Idempotency (--id)

Every increase/decrease must be idempotent based on the provided --id.

- If the same --id is used again with the same command, it must not apply a second time.
- The command should respond as “skipped/already applied” (exact wording is up to you).


### 2) Concurrency safety

Your solution must be safe when two CLI commands run at the same time (for example, two decreases on the same SKU).

- Must not lose updates.
- Must not allow negative stock due to race conditions.

Recommended approach:
- Use a DB transaction per movement
- Apply stock update atomically in SQL (e.g. UPDATE ... WHERE stock + delta >= 0 RETURNING stock)

---

## 🐋 Runtime requirements

The project must run using Docker Compose.

We should be able to run this flow:

    docker compose up -d db
    docker compose run --rm app list
    docker compose run --rm app increase --id req-001 --sku <SKU> --qty 10
    docker compose run --rm app decrease --id req-002 --sku <SKU> --qty 2
    docker compose run --rm app report --low-stock 10 --top 5

### Configuration

Provide:
- .env.example
- App should read DB settings from environment variables

---

## 📋 Deliverables (what we expect in your fork)

- Working Go code
- docker-compose.yml + Dockerfile
- SQL migrations
- A short README.md that explains:
  - how to run the project
  - example commands
  - any assumptions or tradeoffs you made

---

## 👓 Evaluation criteria (how we review)

We’ll look at:

- Correctness (no negative stock, idempotency works)
- Concurrency safety (no lost updates)
- SQL & schema quality (constraints, reasonable structure)
- Go code quality (readability, error handling, structure)
- Dev usability (Docker Compose runs smoothly, README is accurate)

*Bonus (optional):*
- *a few tests (go test ./...)*
- *structured logging*
- *small DX niceties (help text, consistent output, exit codes)*

---

## 🗒️ Notes

- Keep it simple and readable — this is intentionally a small task.
- If you’re unsure about a requirement, document your assumption in the README. 