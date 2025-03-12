# Go-LeetCode Server Guidelines

## Build and Run Commands
- Build: `go build ./...`
- Run: `go run main.go`
- Test all: `go test ./...`
- Test package: `go test ./api/handlers`
- Test single: `go test -v ./api/handlers -run TestRegisterHandler`
- Format: `go fmt ./...`

## Code Style
- **Naming**: PascalCase for exported items, camelCase for internal
- **Imports**: Grouped by standard lib, external libs, internal packages
- **Error Handling**: Use `fmt.Errorf` with context, check `sql.ErrNoRows` for not-found
- **API Responses**: Use `response.JSON()` and `response.Error()` with standard status codes
- **SQL Queries**: Format with backticks and proper indentation
- **Tests**: Use table-driven tests with clear input/expected pairs

## Database
- Use parameterized queries with `$1, $2` placeholders
- Always check rows.Err() after iterating through sql.Rows
- Close database resources in defer statements

## Project Structure
- `api/`: HTTP handlers, middleware and routes
- `models/`: Data models and database operations
- `pkg/`: Shared utilities and helper packages
- `internal/`: Implementation details hidden from external packages