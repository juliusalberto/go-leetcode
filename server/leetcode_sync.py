#!/usr/bin/env python3
import requests
import json
import time
from datetime import datetime, timezone

# Configuration
LEETCODE_USERNAME = "celanapelangi"
APP_USERNAME = "julius"
BASE_URL = "http://localhost:8080"
LIMIT = 2

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
    return data.get("data", {}).get("id")

def create_submission_and_review(user_id, submission_data):
    """Create submission in our database and trigger review creation"""
    # Convert LeetCode timestamp (Unix timestamp) to RFC3339 format
    unix_timestamp = int(submission_data["timestamp"])
    submitted_dt = datetime.fromtimestamp(unix_timestamp, tz=timezone.utc)
    submitted_rfc3339 = submitted_dt.isoformat().replace('+00:00', 'Z')
    
    # Current time in RFC3339
    now_rfc3339 = datetime.now(timezone.utc).isoformat().replace('+00:00', 'Z')
    
    submission = {
        "leetcode_submission_id": submission_data["id"],
        "is_internal": False,
        "user_id": user_id,
        "title": submission_data["title"],
        "title_slug": submission_data["titleSlug"],
        "submitted_at": submitted_rfc3339,
        "created_at": now_rfc3339
    }
    
    # First, create the submission
    submission_response = requests.post(f"{BASE_URL}/api/submissions", json=submission)
    
    if submission_response.status_code not in [201, 200]:
        if submission_response.status_code == 409 and "already exists" in submission_response.text:
            print(f"Submission {submission['leetcode_submission_id']} already exists, skipping...")
            return True
        
        print(f"Error creating submission: {submission_response.status_code}")
        print(submission_response.text)
        return False
    
    print(f"Created submission: {submission['leetcode_submission_id']} - {submission['title']}")
    
    # Then trigger the UpdateOrCreateReview endpoint
    submission["id"] = f"leetcode-{submission_data['id']}"
    submission.pop("is_internal")


    submission.pop("leetcode_submission_id")
    review_response = requests.post(f"{BASE_URL}/api/reviews/update-or-create", json=submission)
    
    if review_response.status_code != 200:
        print(f"Error creating/updating review: {review_response.status_code}")
        print(review_response.text)
        return False
    
    review_data = review_response.json().get("data", {})
    next_review = review_data.get("next_review_at", "unknown")
    days = review_data.get("days_until_review", "unknown")
    
    print(f"Created/updated review for {submission['title']}")
    print(f"Next review at: {next_review} (in {days} days)")
    
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
        
        if create_submission_and_review(user_id, submission):
            successful += 1
        
        # Small delay to avoid overwhelming the server
        time.sleep(0.5)
    
    print(f"\nSummary: Successfully processed {successful} out of {len(submissions)} submissions.")

if __name__ == "__main__":
    main()