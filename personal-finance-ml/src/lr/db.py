import psycopg2
import pandas as pd

def get_connection():
    return psycopg2.connect(
        host= "ep-nameless-lab-a19gn5v2-pooler.ap-southeast-1.aws.neon.tech",
        database="neondb",
        user="neondb_owner",
        password="npg_v78EMtPslUoA",
        port="5432"
    )

def fetch_expense_data(user_id):
    conn = get_connection()

    query = f"""
    SELECT amount, expense_date
    FROM expenses
    WHERE user_id = {user_id}
    ORDER BY expense_date;
    """

    df = pd.read_sql(query, conn)
    conn.close()

    return df
