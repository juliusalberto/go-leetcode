#!/usr/bin/env python3
import requests
import json
import time
from datetime import datetime, timezone

# Configuration
LEETCODE_USERNAME = "celanapelangi"
APP_USERNAME = "julius"
BASE_URL = "http://localhost:8080"
LIMIT = 20

def fetch_leetcode_submissions():
    """Fetch recent accepted submissions from LeetCode GraphQL API"""
    url = "https://leetcode.com/graphql"
    query = """
    query recentAcSubmissions($username: String!, $limit: Int!) {
      recentAcSubmissionList(username: $username, limit: $limit) {
        id
        title
        titleSlug
        timestamp
      }
    }
    """
    
    variables = {
        "username": LEETCODE_USERNAME,
        "limit": LIMIT
    }
    
    payload = {
        "query": query,
        "variables": variables
    }
    
    headers = {
        "Content-Type": "application/json"
    }
    
    response = requests.post(url, json=payload, headers=headers)
    
    if response.status_code != 200:
        print(f"Error fetching from LeetCode API: {response.status_code}")
        print(response.text)
        return []
    
    data = response.json()
    return data.get("data", {}).get("recentAcSubmissionList", [])

def get_user_id():
    """Get user ID for the APP_USERNAME"""
    response = requests.get(f"{BASE_URL}/api/users?username={APP_USERNAME}")
    
    if response.status_code != 200:
        print(f"Error fetching user: {response.status_code}")
        print(response.text)
        return None
    
    data = response.json()
    if "data" in data and data["data"]:
        return data["data"]["ID"]
    
    print(f"User {APP_USERNAME} not found. Creating...")
    
    # Create user if not found
    user_data = {
        "username": APP_USERNAME,
        "leetcode_username": LEETCODE_USERNAME
    }
    
    response = requests.post(f"{BASE_URL}/api/users", json=user_data)
    
    if response.status_code != 201:
        print(f"Error creating user: {response.status_code}")
        print(response.text)
        return None
    
    data = response.json()
    return data.get("data", {}).get("ID")

def process_leetcode_submission(user_id, submission_data):
    """Process a LeetCode submission through the unified endpoint"""
    # Convert LeetCode timestamp (Unix timestamp) to RFC3339 format
    unix_timestamp = int(submission_data["timestamp"])
    submitted_dt = datetime.fromtimestamp(unix_timestamp, tz=timezone.utc)
    submitted_rfc3339 = submitted_dt.isoformat().replace('+00:00', 'Z')
    
    # Current time in RFC3339
    now_rfc3339 = datetime.now(timezone.utc).isoformat().replace('+00:00', 'Z')
    
    # Create submission request matching the server's expected format
    submission = {
        "is_internal": False,
        "leetcode_submission_id": submission_data["id"],
        "user_id": user_id,
        "title": submission_data["title"],
        "title_slug": submission_data["titleSlug"],
        "submitted_at": submitted_rfc3339
    }
    
    # Send to the unified endpoint
    response = requests.post(f"{BASE_URL}/api/reviews/process-submission", json=submission)
    
    if response.status_code != 200:
        if response.status_code == 409 and "already exists" in response.text:
            print("submission already exists")
            return True
        print(f"Error processing submission: {response.status_code}")
        print(response.text)
        return False
    
    # Parse the response
    data = response.json().get("data", {})
    next_review = data.get("next_review_at", "unknown")
    days = data.get("days_until_review", "unknown")
    is_due = data.get("is_due", False)
    submission_id = data.get("submission_id", "unknown")
    
    # Format a nice output message
    due_status = "DUE NOW" if is_due else f"due in {days} days"
    print(f"Processed: {submission['title']} (LeetCode ID: {submission['leetcode_submission_id']})")
    print(f"Server ID: {submission_id}")
    print(f"Next review: {next_review} ({due_status})")
    
    return True

def main():
    print(f"Fetching submissions for LeetCode user: {LEETCODE_USERNAME}")
    submissions = fetch_leetcode_submissions()
    
    if not submissions:
        print("No submissions found.")
        return
    
    print(f"Found {len(submissions)} recent accepted submissions.")
    
    user_id = get_user_id()
    if not user_id:
        print("Failed to get or create user.")
        return
    
    print(f"Using user ID: {user_id}")
    
    successful = 0
    for submission in submissions:
        print(f"\nProcessing: {submission['title']} (ID: {submission['id']})")
        timestamp = datetime.fromtimestamp(int(submission['timestamp']), tz=timezone.utc)
        print(f"Submitted at: {timestamp} UTC")
        
        if process_leetcode_submission(user_id, submission):
            successful += 1
        
        # Small delay to avoid overwhelming the server
        time.sleep(0.5)
    
    print(f"\nSummary: Successfully processed {successful} out of {len(submissions)} submissions.")

if __name__ == "__main__":
    main()