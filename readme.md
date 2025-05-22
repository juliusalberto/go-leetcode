# LeetCode Practice Assistant (SPACECODE)

[![Go](https://img.shields.io/badge/Go-1.24+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![React Native](https://img.shields.io/badge/React_Native-Expo-61DAFB?style=flat&logo=react)](https://reactnative.dev/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-latest-336791?style=flat&logo=postgresql)](https://www.postgresql.org/)
[![Docker](https://img.shields.io/badge/Docker-latest-2496ED?style=flat&logo=docker)](https://www.docker.com/)
[![Supabase](https://img.shields.io/badge/Supabase-Database-3ECF8E?style=flat&logo=supabase)](https://supabase.com/)
[![Fly.io](https://img.shields.io/badge/Fly.io-Deployment-9A50FF?style=flat&logo=fly-dot-io)](https://fly.io/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

> A mobile application to help developers maintain consistent LeetCode practice through automated tracking and spaced repetition reminders.

## 📋 Table of Contents

- [Features](#-features)
- [Tech Stack](#-tech-stack)
- [API Format](#-api-format)
- [Getting Started](#-getting-started)
  - [Prerequisites](#prerequisites)
  - [Backend Setup](#backend-setup)
  - [Frontend Setup](#frontend-setup)
  - [Docker Deployment](#docker-deployment)
  - [Fly.io Deployment](#flyio-deployment)
- [Testing](#-testing)
- [Project Structure](#-project-structure)
- [Database](#-database)
- [Contributing](#-contributing)
- [License](#-license)

## ✨ Features

- **Automated Tracking**: Seamlessly integrates with LeetCode's GraphQL API to track your submission history
- **Spaced Repetition**: Intelligently schedules problem reviews based on solving performance
- **Problem Database**: Complete catalog of LeetCode problems with filtering by difficulty, topics, and more

## 🛠 Tech Stack

### Backend
- **Go**: Fast and efficient REST API server
- **PostgreSQL**: Robust data storage solution
- **Docker**: Containerization for easy deployment
- **Supabase**: Database hosting and migrations
- **Fly.io**: Cloud hosting platform

### Frontend
- **React Native**: Cross-platform mobile application with TypeScript
- **Expo**: Simplified React Native development
- **Redux**: State management
- **Axios**: API communication
- **NativeWind**: Tailwind CSS for React Native
- **Jest**: Testing framework

## 📡 API Format

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

Detailed API documentation can be found in [API-DOCUMENTATION.md](server/API-DOCUMENTATION.md)

## 🚀 Getting Started

### Prerequisites

- Go 1.24+
- Node.js 14+
- PostgreSQL
- Redis
- Python (for data scripts)

### Backend Setup

```bash
# Clone the repository
git clone https://github.com/juliusalberto/leetcode-practice-assistant.git
cd leetcode-practice-assistant

# Install dependencies
cd server
go mod download

# Configure environment variables
cp .env.example .env
# Edit .env with your configuration

# Run migrations (if needed)
migrate -database ${POSTGRESQL_URL} -path migrations up

# Run the server
go run main.go
```

### Frontend Setup

```bash
# Navigate to frontend directory
cd client

# Install dependencies
npm install

# Start the Expo development server
npx expo start

# Run on iOS
npm run ios

# Run on Android
npm run android
```

### Docker Deployment

```bash
# Build the Docker image
cd server
docker build -t leetcode-server .

# Run with environment variables
docker run -p 8080:8080 \
  -e DB_HOST=your-db-host \
  -e DB_PORT=5432 \
  -e DB_USER=your-db-user \
  -e DB_PASSWORD=your-db-password \
  -e DB_NAME=your-db-name \
  leetcode-server

# Alternatively, use the pre-built image
docker pull ghcr.io/juliusalberto/go-leetcode:latest
```

### Fly.io Deployment

The application is configured for deployment on Fly.io:

```bash
# Deploy using the fly CLI
fly deploy
```

## 🧪 Testing

```bash
# Backend tests
cd server
go test ./...

# Frontend tests
cd client
npm test
```

## 📂 Project Structure

```
leetcode-practice-assistant/
├── server/                # Go server code
│   ├── api/               # API handlers and routes
│   ├── models/            # Database models
│   ├── internal/          # Internal packages
│   ├── pkg/               # Shared packages
│   ├── migrations/        # Database migrations
│   └── Dockerfile         # Docker configuration
├── client/                # React Native Expo app
│   ├── app/               # Main application screens
│   ├── components/        # Reusable components
│   ├── contexts/          # React contexts
│   ├── services/          # API services
│   ├── utils/             # Helper functions
│   └── assets/            # Static assets
├── supabase/              # Supabase configuration
│   └── migrations/        # Supabase migrations
└── .github/               # GitHub Actions workflows
```

## 🗄 Database

The application uses PostgreSQL for data storage with the following key tables:

- **users**: User accounts and preferences
- **problems**: LeetCode problem catalog
- **submissions**: User problem submissions
- **reviews**: Spaced repetition review schedules

Supabase is used for database hosting and management in production.

## 👥 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
