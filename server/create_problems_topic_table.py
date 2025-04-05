import psycopg2
import json
from psycopg2.extras import execute_values

def main():
    # Connect to PostgreSQL database
    conn = psycopg2.connect(
        host="localhost",
        port=54322,
        database="postgres",
        user="postgres",
        password="postgres"
    )
    cur = conn.cursor()

    try:
        # Create problems_topic table if it doesn't exist
        cur.execute("""
        CREATE TABLE IF NOT EXISTS public.problems_topic (
            problem_id INTEGER NOT NULL,
            topic_slug TEXT NOT NULL,
            PRIMARY KEY (problem_id, topic_slug),
            FOREIGN KEY (problem_id) REFERENCES public.problems(id)
        )
        """)

        # Get all problems with their topic_tags
        cur.execute("SELECT id, topic_tags FROM public.problems")
        problems = cur.fetchall()

        # Prepare data for bulk insert
        data_to_insert = []
        for problem_id, topic_tags_json in problems:
            if topic_tags_json:
                for elem in topic_tags_json:
                    print(elem)
                    data_to_insert.append((problem_id, elem['slug']))

        # Bulk insert using execute_values for better performance
        execute_values(
            cur,
            "INSERT INTO problems_topic (problem_id, topic_slug) VALUES %s ON CONFLICT DO NOTHING",
            data_to_insert
        )

        conn.commit()
        print(f"Successfully inserted {len(data_to_insert)} topic mappings")

    except Exception as e:
        conn.rollback()
        print(f"Error occurred: {e}")
    finally:
        cur.close()
        conn.close()

if __name__ == "__main__":
    main()