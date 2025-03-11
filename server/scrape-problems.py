import requests
import json
import time
import psycopg2
import os
from psycopg2.extras import Json

# Database connection parameters
db_params = {
    'host': os.environ.get('DB_HOST', 'localhost'),
    'port': os.environ.get('DB_PORT', '5432'),
    'user': os.environ.get('DB_USER', 'leetcode_app'),
    'password': os.environ.get('DB_PASSWORD', 'ilovel33t'),
    'database': os.environ.get('DB_NAME', 'leetcode_app_test')
}

# Base URLs
PROBLEMS_LIST_URL = "https://leetcode.com/api/problems/algorithms/"
PROBLEM_DETAIL_URL = "http://localhost:3000/select?titleSlug="


# Insert problem into database
def insert_problem(conn, problem_data):
    with conn.cursor() as cur:
        # Map difficulty level to text
        difficulty_map = {1: "Easy", 2: "Medium", 3: "Hard"}
        
        # Extract data from the problem detail response
        problem = {
            'id': problem_data.get('questionId'),
            'frontend_id': problem_data.get('questionFrontendId'),
            'title': problem_data.get('questionTitle'),
            'title_slug': problem_data.get('titleSlug'),
            'difficulty': problem_data.get('difficulty'),
            'is_paid_only': problem_data.get('isPaidOnly', False),
            'content': problem_data.get('question'),
            'topic_tags': Json(problem_data.get('topicTags', [])),
            'example_testcases': problem_data.get('exampleTestcases'),
            'similar_questions': Json(problem_data.get('similarQuestions', '[]'))
        }

        if not problem['content']:
            problem['content'] = ""
        
        cur.execute("""
        INSERT INTO problems 
        (id, frontend_id, title, title_slug, difficulty, is_paid_only, content, topic_tags, example_testcases, similar_questions)
        VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)
        ON CONFLICT (title_slug) DO UPDATE SET
        title = EXCLUDED.title,
        difficulty = EXCLUDED.difficulty,
        is_paid_only = EXCLUDED.is_paid_only,
        content = EXCLUDED.content,
        topic_tags = EXCLUDED.topic_tags,
        example_testcases = EXCLUDED.example_testcases,
        similar_questions = EXCLUDED.similar_questions
        """, (
            problem['id'],
            problem['frontend_id'],
            problem['title'],
            problem['title_slug'],
            problem['difficulty'],
            problem['is_paid_only'],
            problem['content'],
            problem['topic_tags'],
            problem['example_testcases'],
            problem['similar_questions']
        ))
    conn.commit()

def problem_exists(conn, problem_id):
    with conn.cursor() as cur:
        cur.execute("SELECT 1 FROM problems WHERE id = %s", (problem_id,))
        return cur.fetchone() is not None

def main():
    # Connect to the database
    try:
        conn = psycopg2.connect(**db_params)
        print("Connected to database successfully!")
        
        # Get the list of all problems
        print("Fetching problem list...")
        response = requests.get(PROBLEMS_LIST_URL)
        problems_list = response.json()
        
        # Process each problem
        total_problems = len(problems_list['stat_status_pairs'])
        print(f"Found {total_problems} problems. Starting download...")
        
        for i, problem in enumerate(problems_list['stat_status_pairs']):
            title_slug = problem['stat']['question__title_slug']
            paid_only = problem['paid_only']
            problem_id = str(problem['stat']['question_id'])

            if problem_exists(conn=conn, problem_id=problem_id):
                continue
            
            print(f"Processing problem {i+1}/{total_problems}: {title_slug}")
            
            # Get detailed problem data
            try:
                while True:
                    detail_url = f"{PROBLEM_DETAIL_URL}{title_slug}"
                    detail_response = requests.get(detail_url)
                    if detail_response.status_code == 200:
                        problem_detail = detail_response.json()
                        
                        if not problem_detail.get('questionId'):
                            continue

                        insert_problem(conn, problem_detail)
                        print(f"Successfully saved: {title_slug}")
                        break
                    else:
                        print(f"Failed to fetch details for {title_slug}: HTTP {detail_response.status_code}")
            except Exception as e:
                print(f"Error processing {title_slug}: {str(e)}")
            
            # Wait to avoid rate limiting
            time.sleep(10)
        
        print("All problems processed!")
        
    except Exception as e:
        print(f"Error: {str(e)}")
    finally:
        if conn:
            conn.close()
            print("Database connection closed.")

if __name__ == "__main__":
    main()