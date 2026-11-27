"""Tenant export repository helper (parameterized SQL, no org filter)."""

def fetch_export_by_id(cursor, export_id: str):
    cursor.execute(
        "SELECT id, organization_id, file_url, created_at FROM billing_exports WHERE id = %s",
        (export_id,),
    )
    return cursor.fetchone()
