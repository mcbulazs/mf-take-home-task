# 🦊 Mobilfox – Take-home task

See the original [`task`](/task/README.md)


## Packages

Outside of standard packages the following (and their dependencies) were used:
- [lib/pq](https://github.com/lib/pq)
- [spf13/cobra](https://github.com/spf13/cobra)
- [DATA-DOG/go-sqlmock](https://github.com/DATA-DOG/go-sqlmock)
- [stretchr/testify](https://github.com/stretchr/testify)

---

## Assumptions

- If a stock change fails (e.g. SKU not found or stock would become negative), the idempotency key is **not stored**, meaning the operation can be retried.
- Schema migration and seeding run automatically on application startup.


---

## Environment

The only required file is a .env file.
An example is provided in .env.example.

---

## Help

The tool has an extensive help which you can excess by the command **`help`** or with the flag **`--help`** (**`-h`**).

---

## Running

As the original task required the cli can be used like the following:
```sh
    docker compose up -d db
    docker compose run --rm app list
    docker compose run --rm app increase --id req-001 --sku <SKU> --qty 10
    docker compose run --rm app decrease --id req-002 --sku <SKU> --qty 2
    docker compose run --rm app report --low-stock 10 --top 5
```
## Tests

The repository tests use [`go-sqlmock`](https://github.com/DATA-DOG/go-sqlmock)
```sh
    go test ./test/* -v
```
