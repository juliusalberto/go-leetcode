# LeetCode Practice Server

A Go-based server for tracking LeetCode problem submissions and managing spaced repetition review schedules.

## Running with Docker

### Build the Docker image locally

```bash
# From the server directory
docker build -t leetcode-server .
```

### Run the Docker container

```bash
# Run with environment variables
docker run -p 8080:8080 \
  -e DB_HOST=your-db-host \
  -e DB_PORT=5432 \
  -e DB_USER=your-db-user \
  -e DB_PASSWORD=your-db-password \
  -e DB_NAME=your-db-name \
  leetcode-server
```

### Using the GitHub Container Registry image

After a successful build via GitHub Actions, you can also pull and run the published image:

```bash
# Pull the image
docker pull ghcr.io/juliusalberto/go-leetcode:latest

# Run the container
docker run -p 8080:8080 \
  -e DB_HOST=your-db-host \
  -e DB_PORT=5432 \
  -e DB_USER=your-db-user \
  -e DB_PASSWORD=your-db-password \
  -e DB_NAME=your-db-name \
  ghcr.io/juliusalberto/go-leetcode:latest
```

## Local Development

### Prerequisites

- Go 1.24+
- PostgreSQL

### Setup

1. Clone the repository
2. Create a `.env` file with your database credentials:
   ```
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your-db-user
   DB_PASSWORD=your-db-password
   DB_NAME=your-db-name
   ```
3. Run the migrations: `migrate -database ${POSTGRESQL_URL} -path migrations up`
4. Seed the test data: `psql -h localhost -U your-db-user -d your-db-name -f migrations/seed_test_problems.sql`
5. Build: `go build ./...`
6. Run: `go run main.go`
7. Test: `go test ./...`

## CI/CD Pipeline

This project uses GitHub Actions for continuous integration and deployment:

1. **Automated Testing**: On every push and pull request, all tests are automatically run against a PostgreSQL database.
2. **Docker Image Building**: When code is pushed to the main branch and passes all tests, a Docker image is built and pushed to GitHub Container Registry (ghcr.io).

## API Documentation

See [API-DOCUMENTATION.md](API-DOCUMENTATION.md) for details on the available endpoints and request formats.