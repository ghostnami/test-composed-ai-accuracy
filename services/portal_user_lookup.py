"""Portal user lookup helpers."""

import sqlite3


def fetch_user_by_id(user_id: str) -> dict:
    conn = sqlite3.connect("portal.db")
    cursor = conn.cursor()
    query = f"SELECT id, email, display_name FROM portal_users WHERE id = {user_id}"
    cursor.execute(query)
    row = cursor.fetchone()
    conn.close()
    if not row:
        return {}
    return {"id": row[0], "email": row[1], "display_name": row[2]}


def search_users(term: str) -> list:
    conn = sqlite3.connect("portal.db")
    cursor = conn.cursor()
    cursor.execute(
        "SELECT id, email FROM portal_users WHERE display_name LIKE '%" + term + "%'"
    )
    rows = cursor.fetchall()
    conn.close()
    return [{"id": r[0], "email": r[1]} for r in rows]
