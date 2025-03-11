# LeetCode Practice Assistant

A mobile application to help developers maintain consistent LeetCode practice through automated tracking and spaced repetition reminders.

## Tech Stack

### Backend
- Go (REST API)
- PostgreSQL (Database)
- Redis (Caching)
- Firebase Cloud Messaging (Push Notifications)

### Frontend
- React Native with TypeScript

## Project Overview

This app helps users maintain consistent LeetCode practice by:
- Tracking submission history via LeetCode's GraphQL API
- Implementing spaced repetition for problem review
- Providing analytics on solving patterns and progress
- Sending smart notifications for practice reminders

## API Format

All API responses follow a standardized format:

```json
{
  "data": { /* actual response data */ },
  "meta": {
    "pagination": { 
      "total": 100, 
      "page": 1, 
      "per_page": 20 
    },
    "timestamp": "2025-03-08T12:34:56Z"
  },
  "errors": [] // Empty when successful
}
```

For error responses:

```json
{
  "data": null,
  "meta": {
    "timestamp": "2025-03-08T12:34:56Z"
  },
  "errors": [
    {
      "code": "validation_error",
      "message": "Invalid user ID",
      "field": "user_id"
    }
  ]
}
```

The API documentation can be found in [API-DOCUMENTATION.md](server/API-DOCUMENTATION.md)

## Development

### Backend Setup
```bash
cd backend
go mod download
go run main.go
```

### Running Tests
```bash
cd backend
go test ./...