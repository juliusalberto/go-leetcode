import os
import psycopg2
import glob
from dotenv import load_dotenv

def insert_leetcode_solutions():
    print("Starting insert_leetcode_solutions function")
    load_dotenv()
    print("Loaded .env file")

    # Database credentials from .env file
    db_host = os.environ.get("DB_HOST")
    db_port = os.environ.get("DB_PORT")
    db_user = os.environ.get("DB_USER")
    db_password = os.environ.get("DB_PASSWORD")
    db_name = os.environ.get("DB_NAME")
    print(f"DB credentials: host={db_host}, port={db_port}, user={db_user}, dbname={db_name}")

    # Construct the database connection string
    conn_string = f"host={db_host} port={db_port} user={db_user} password={db_password} dbname={db_name}"

    # Paths to the solution directories - adjusted to use parent directory
    python_solutions_dir = "../LeetCode-Solutions-master/Python"
    cpp_solutions_dir = "../LeetCode-Solutions-master/C++"
    
    # Check if directories exist
    print(f"Python solutions directory exists: {os.path.exists(python_solutions_dir)}")
    print(f"C++ solutions directory exists: {os.path.exists(cpp_solutions_dir)}")
    
    # List a few Python files to confirm the path works
    python_files = glob.glob(os.path.join(python_solutions_dir, "*.py"))
    print(f"First 5 Python files (if any): {python_files[:5]}")

    try:
        # Connect to the database
        print("Connecting to the database...")
        conn = psycopg2.connect(conn_string)
        cursor = conn.cursor()
        print("Database connection successful")

        # Process Python solutions
        print("Starting to process Python solutions...")
        python_files = glob.glob(os.path.join(python_solutions_dir, "*.py"))
        print(f"Found {len(python_files)} Python solution files")
        
        for filename in python_files:
            title_slug = os.path.splitext(os.path.basename(filename))[0]
            language = "Python"
            print(f"Processing Python solution: {title_slug}")

            # Read the solution code
            with open(filename, "r") as f:
                solution_code = f.read()

            # Get the problem_id from the problems table
            cursor.execute("SELECT id FROM problems WHERE title_slug = %s", (title_slug,))
            result = cursor.fetchone()

            if result:
                problem_id = result[0]

                # Insert the solution into the problem_solutions table
                try:
                    cursor.execute(
                        "INSERT INTO problem_solutions (problem_id, language, solution_code) VALUES (%s, %s, %s)",
                        (problem_id, language, solution_code),
                    )
                    conn.commit()
                    print(f"Inserted solution for {title_slug} ({language})")
                except psycopg2.errors.UniqueViolation:
                    conn.rollback()
                    print(f"Solution for {title_slug} ({language}) already exists")
                except Exception as e:
                    conn.rollback()
                    print(f"Error inserting solution for {title_slug} ({language}): {e}")
            else:
                print(f"Problem not found in problems table for {title_slug}")

        # Process C++ solutions (similar to Python)
        print("Starting to process C++ solutions...")
        cpp_files = glob.glob(os.path.join(cpp_solutions_dir, "*.cpp"))
        print(f"Found {len(cpp_files)} C++ solution files")
        
        for filename in cpp_files:
            title_slug = os.path.splitext(os.path.basename(filename))[0]
            language = "C++"
            print(f"Processing C++ solution: {title_slug}")

            # Read the solution code
            with open(filename, "r") as f:
                solution_code = f.read()

            # Get the problem_id from the problems table
            cursor.execute("SELECT id FROM problems WHERE title_slug = %s", (title_slug,))
            result = cursor.fetchone()

            if result:
                problem_id = result[0]

                # Insert the solution into the problem_solutions table
                try:
                    cursor.execute(
                        "INSERT INTO problem_solutions (problem_id, language, solution_code) VALUES (%s, %s, %s)",
                        (problem_id, language, solution_code),
                    )
                    conn.commit()
                    print(f"Inserted solution for {title_slug} ({language})")
                except psycopg2.errors.UniqueViolation:
                    conn.rollback()
                    print(f"Solution for {title_slug} ({language}) already exists")
                except Exception as e:
                    conn.rollback()
                    print(f"Error inserting solution for {title_slug} ({language}): {e}")
            else:
                print(f"Problem not found in problems table for {title_slug}")

    except psycopg2.Error as e:
        print(f"Database connection error: {e}")
    except Exception as e:
        print(f"An error occurred: {e}")
    finally:
        # Close the database connection
        if conn:
            cursor.close()
            conn.close()

if __name__ == "__main__":
    print("Starting the command!")
    print(f"Current working directory: {os.getcwd()}")
    insert_leetcode_solutions()
