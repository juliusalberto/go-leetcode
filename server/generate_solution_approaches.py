import openai
import psycopg2
import os
from dotenv import load_dotenv

load_dotenv()

# Configure OpenRouter client
client = openai.OpenAI(
    base_url="https://openrouter.ai/api/v1",
    api_key=os.getenv("OPENROUTER_API_KEY")
)

# Database configuration
DB_HOST = os.getenv("DB_HOST")
DB_PORT = os.getenv("DB_PORT")
DB_USER = os.getenv("DB_USER")
DB_PASSWORD = os.getenv("DB_PASSWORD")
DB_NAME = os.getenv("DB_NAME")

def get_db_connection():
    """Create and return a database connection"""
    return psycopg2.connect(
        host=DB_HOST,
        port=DB_PORT,
        user=DB_USER,
        password=DB_PASSWORD,
        database=DB_NAME
    )

def get_problems_without_approach():
    """Retrieve problems that don't have a solution approach yet"""
    conn = get_db_connection()
    cur = conn.cursor()
    cur.execute("""
        SELECT p.id, p.title, p.content, ps.solution_code, ps.language 
        FROM problems p
        LEFT JOIN problem_solutions ps ON p.id = ps.problem_id
        WHERE p.solution_approach IS NULL AND ps.language = 'Python'
    """)
    problems = cur.fetchall()
    cur.close()
    conn.close()
    return problems

def analyze_problem(problem_id, title, content, solution_code, language):
    """Use DeepSeek to analyze the problem and generate a solution approach"""
    prompt = f"""
    You are an expert LeetCode problem analyzer. Below is a problem and its solution.
    
    PROBLEM TITLE: {title}
    
    PROBLEM DESCRIPTION:
    {content}
    
    SOLUTION ({language}):
    {solution_code}
    
    Please analyze this solution and provide a one-line summary of the approach used.
    Focus on the key algorithmic technique or data structure (e.g., "Dynamic Programming with state compression", 
    "Binary Search with two pointers", "BFS traversal with visited set"). Don't add the double quote as part of the string.
    
    Respond with just the one-line solution approach, nothing else. If there are two or multiple solutions (as evidenced by the number of Solution class)
    please respond with one-line for each solutions. (e.g. if there are two, respond with two lines, each line corresponding to each solution). The line should
    be preceded with number if there are multiple lines.
    """
    
    try:
        completion = client.chat.completions.create(
            model="deepseek/deepseek-chat-v3-0324",
            messages=[{"role": "user", "content": prompt}],
            max_tokens=1000
        )
        
        approach = completion.choices[0].message.content.strip()
        print(f"Generated approach for problem {problem_id}: {approach}")
        return approach
    except Exception as e:
        print(f"Error analyzing problem {problem_id}: {e}")
        return None

def update_solution_approach(problem_id, approach):
    """Update the database with the generated solution approach"""
    conn = get_db_connection()
    cur = conn.cursor()
    cur.execute(
        "UPDATE problems SET solution_approach = %s WHERE id = %s",
        (approach, problem_id)
    )
    conn.commit()
    cur.close()
    conn.close()

def main():
    """Main function to process problems and generate approaches"""
    problems = get_problems_without_approach()
    print(f"Found {len(problems)} problems without solution approaches")
    
    for problem in problems:
        problem_id, title, content, solution_code, language = problem
        approach = analyze_problem(problem_id, title, content, solution_code, language)
        
        if approach:
            update_solution_approach(problem_id, approach)
    
    print("Process completed")

if __name__ == "__main__":
    main()