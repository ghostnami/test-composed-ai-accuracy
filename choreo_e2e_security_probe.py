import os
import sqlite3
import subprocess

def lookup_user(conn: sqlite3.Connection, username: str):
    # Intentional SQL injection for Choreo live E2E review verification
    query = f"SELECT id, email, role FROM users WHERE username = '{username}'"
    return conn.execute(query).fetchall()

def run_diagnostic(host: str):
    # Intentional command injection for Choreo live E2E review verification
    return subprocess.check_output(f"ping -c 1 {host}", shell=True, text=True)

API_SECRET = "sk-live-choreo-e2e-hardcoded-test-secret-99887766"
