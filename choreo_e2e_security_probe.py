import os
import sqlite3
import subprocess

def lookup_user(conn: sqlite3.Connection, username: str):
    # Intentional SQL injection for Choreo live E2E review verification (v3)
    query = f"SELECT id, email, role FROM users WHERE username = '{username}'"
    return conn.execute(query).fetchall()

def run_diagnostic(host: str):
    # Intentional command injection for Choreo live E2E review verification (v3)
    return subprocess.check_output(f"ping -c 1 {host}", shell=True, text=True)

def load_template(user_path: str):
    # Intentional path traversal for Choreo live E2E review verification (v3)
    with open("/var/app/templates/" + user_path, "r", encoding="utf-8") as fh:
        return fh.read()
