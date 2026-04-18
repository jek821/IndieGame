Curl Request Format: curl http://localhost:8090/createUser -H "Content-Type: application/json" --data '{"username":"jacob1"}'

## New Feature Checklist

1. **`transport.go`** — Add `RequestCode` constant
2. **`transport.go`** — Add `ResponseCode` constant(s)
3. **`transport.go`** — Define request struct
4. **`transport.go`** — Add case to `unpackRequest`
5. **`store.go`** — Update `migrate()` *(if the feature needs a new table or column)*
6. **`store.go`** — Add method to `Store` interface
7. **`store.go`** — Implement on `SqliteStore`
8. **`handlers.go`** — Write `handle<Feature>`
9. **`routes.go`** — Register the route
