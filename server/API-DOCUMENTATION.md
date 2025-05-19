# SPACECODE API Documentation

This document provides comprehensive information about the SPACECODE (LeetCode Practice Assistant) API endpoints, their functionality, request/response formats, and usage guidelines.

## 📋 Table of Contents
- [Overview](#overview)
- [Base URL](#base-url)
- [Authentication](#authentication)
- [API Response Format](#api-response-format)
- [Rate Limiting](#rate-limiting)
- [Using the Postman Collection](#using-the-postman-collection)
- [Endpoints](#endpoints)
  - [Health](#health)
  - [Users](#users)
  - [Problems](#problems)
  - [Submissions](#submissions)
  - [Reviews](#reviews)
- [Error Codes](#error-codes)

## Overview

The SPACECODE API enables developers to:
- Access LeetCode problem data
- Track user submissions and progress
- Manage spaced repetition review schedules
- Retrieve user-specific analytics

## Base URL

```
Production: https://api.spacecode.fly.dev
Development: http://localhost:8080
```

## Authentication

The API currently uses simple API key authentication. Include your API key in the request header:

```
X-API-Key: your_api_key_here
```

## API Response Format

All API responses follow a standardized format:

```json
{
  "data": { /* your actual response data */ },
  "meta": {
    "pagination": { "total": 100, "page": 1, "per_page": 20 },
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

## Rate Limiting

API requests are limited to 100 requests per minute per API key. The following headers are included in responses:

```
X-RateLimit-Limit: 100
X-RateLimit-Remaining: 99
X-RateLimit-Reset: 1616173799
```

## Using the Postman Collection

A Postman collection file named `LeetCode Practice API.postman_collection.json` is provided to help you explore and test the API.

1. Import the collection into Postman
2. Set the `base_url` environment variable to `http://localhost:8080` (or your server URL)
3. Use the predefined requests to interact with the API

## Endpoints

### Health

#### GET /health
Check if the API server is running properly.

**Response:**
```json
{
  "data": {
    "status": "ok",
    "version": "1.0.0"
  },
  "meta": {
    "timestamp": "2025-03-11T12:34:56Z"
  },
  "errors": []
}
```

### Users

#### POST /api/users/register
Register a new user.

**Request:**
```json
{
  "username": "testuser",
  "leetcode_username": "leetcode_testuser"
}
```

**Response:**
```json
{
  "data": {
    "id": 1,
    "username": "testuser",
    "leetcode_username": "leetcode_testuser",
    "created_at": "2025-03-11T12:34:56Z"
  },
  "meta": {
    "timestamp": "2025-03-11T12:34:56Z"
  },
  "errors": []
}
```

#### GET /api/users/{id}
Retrieve user information by ID.

**Response:**
```json
{
  "data": {
    "id": 1,
    "username": "testuser",
    "leetcode_username": "leetcode_testuser",
    "created_at": "2025-03-11T12:34:56Z",
    "stats": {
      "problems_solved": 42,
      "easy_count": 20,
      "medium_count": 15,
      "hard_count": 7
    }
  },
  "meta": {
    "timestamp": "2025-03-11T12:34:56Z"
  },
  "errors": []
}
```

### Problems

#### GET /api/problems
Get a list of problems with filtering and pagination.

**Parameters:**
- `limit` (optional): Number of problems to return (default: 20)
- `offset` (optional): Number of problems to skip (default: 0)
- `difficulty` (optional): Filter by difficulty ("Easy", "Medium", "Hard")
- `order_by` (optional): Field to sort by (e.g., "frontend_id", "difficulty")
- `order_dir` (optional): Sort direction ("asc" or "desc")
- `search` (optional): Search keyword in problem titles
- `tags` (optional): Filter by problem tags
- `paid_only` (optional): Filter by paid status ("true" or "false")

**Response:**
```json
{
  "data": [
    {
      "id": 1,
      "frontend_id": 1,
      "title": "Two Sum",
      "title_slug": "two-sum",
      "difficulty": "Easy",
      "is_paid_only": false,
      "content": "<p>Problem description...</p>",
      "topic_tags": [
        {"name": "Array", "slug": "array"},
        {"name": "Hash Table", "slug": "hash-table"}
      ],
      "example_testcases": "[2,7,11,15]\n9",
      "similar_questions": [],
      "created_at": "2025-03-11T12:34:56Z"
    }
  ],
  "meta": {
    "pagination": {
      "total": 100,
      "page": 1,
      "per_page": 10
    },
    "timestamp": "2025-03-11T12:34:56Z"
  },
  "errors": []
}
```

#### GET /api/problems/by-id?id={id}
Get a problem by its internal ID.

#### GET /api/problems/by-frontend-id?frontend_id={frontend_id}
Get a problem by its LeetCode frontend ID.

#### GET /api/problems/by-slug?slug={slug}
Get a problem by its slug.

### Submissions

#### GET /api/submissions?user_id={user_id}
Get submissions for a user.

**Parameters:**
- `user_id` (required): The ID of the user to get submissions for
- `limit` (optional): Number of submissions to return (default: 20)
- `offset` (optional): Number of submissions to skip (default: 0)
- `problem_id` (optional): Filter by problem ID
- `from_date` (optional): Filter submissions after this date (ISO 8601 format)
- `to_date` (optional): Filter submissions before this date (ISO 8601 format)

**Response:**
```json
{
  "data": [
    {
      "id": "internal-user-123abc",
      "user_id": 1,
      "title": "Two Sum",
      "title_slug": "two-sum",
      "created_at": "2025-03-11T12:34:56Z",
      "submitted_at": "2025-03-11T12:30:00Z"
    }
  ],
  "meta": {
    "pagination": {
      "total": 42,
      "page": 1,
      "per_page": 20
    },
    "timestamp": "2025-03-11T12:34:56Z"
  },
  "errors": []
}
```

#### POST /api/submissions
Create a new submission.

**Request:**
```json
{
  "user_id": 1,
  "title": "Two Sum",
  "title_slug": "two-sum",
  "submitted_at": "2025-03-11T12:30:00Z"
}
```

**Response:**
```json
{
  "data": {
    "id": "internal-user-123abc",
    "user_id": 1,
    "title": "Two Sum",
    "title_slug": "two-sum",
    "created_at": "2025-03-11T12:34:56Z",
    "submitted_at": "2025-03-11T12:30:00Z"
  },
  "meta": {
    "timestamp": "2025-03-11T12:34:56Z"
  },
  "errors": []
}
```

### Reviews

#### GET /api/reviews?user_id={user_id}&status={status}&page={page}&per_page={per_page}
Get reviews for a user with optional status filter and pagination.

**Parameters:**
- `user_id` (required): User ID to get reviews for
- `status` (optional): Filter by review status
  - `due`: Return only reviews that are due or overdue
  - `upcoming`: Return only reviews scheduled for the future
  - Default (no status parameter): Return both due and upcoming reviews
- `page` (optional): Page number for pagination (default: 1)
- `per_page` (optional): Number of items per page (default: 10, max: 100)

**Response:**
```json
{
  "data": [
    {
      "id": 1,
      "submission_id": "internal-user-123abc",
      "next_review_at": "2025-03-10T12:34:56Z",
      "created_at": "2025-03-09T12:34:56Z",
      "stability": 3.0,
      "difficulty": 5.0,
      "elapsed_days": 0,
      "scheduled_days": 1,
      "reps": 1,
      "lapses": 0,
      "state": 2,
      "last_review": "2025-03-09T12:34:56Z",
      "problem": {
        "title": "Two Sum",
        "difficulty": "Easy"
      }
    }
  ],
  "meta": {
    "pagination": {
      "total": 25,
      "page": 1,
      "per_page": 10
    },
    "timestamp": "2025-03-11T12:34:56Z"
  },
  "errors": []
}
```

#### POST /api/reviews
Create a new review schedule for a submission.

**Request:**
```json
{
  "submission_id": "internal-user-123abc"
}
```

**Response:**
```json
{
  "data": {
    "id": 1,
    "next_review_at": "2025-03-12T12:34:56Z"
  },
  "meta": {
    "timestamp": "2025-03-11T12:34:56Z"
  },
  "errors": []
}
```

#### PUT /api/reviews
Update a review with a rating.

**Request:**
```json
{
  "review_id": 1,
  "rating": 3
}
```

**Rating Values:**
- 1: Again (difficult, needs to be reviewed again soon)
- 2: Hard (challenging but manageable)
- 3: Good (comfortable)
- 4: Easy (very comfortable)

**Response:**
```json
{
  "data": {
    "success": true,
    "next_review_at": "2025-03-14T12:34:56Z",
    "days_until_review": 3
  },
  "meta": {
    "timestamp": "2025-03-11T12:34:56Z"
  },
  "errors": []
}
```

## Error Codes

| Code | Description |
|------|-------------|
| `authentication_error` | Invalid or missing API key |
| `validation_error` | Invalid request parameters |
| `resource_not_found` | The requested resource was not found |
| `rate_limit_exceeded` | You have exceeded the allowed request rate |
| `internal_error` | An unexpected error occurred on the server |