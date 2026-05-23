"""In-memory persistence stubs for accuracy validation scenarios."""

from typing import Any, Dict, List, Optional


INVOICES: Dict[str, Dict[str, Any]] = {
    "inv_1001": {
        "id": "inv_1001",
        "org_id": "org_acme",
        "customer_id": "cust_alice",
        "owner_user_id": "user_alice",
        "amount_cents": 125000,
        "status": "open",
        "line_items": [{"sku": "enterprise-seat", "qty": 25}],
    },
    "inv_1002": {
        "id": "inv_1002",
        "org_id": "org_acme",
        "customer_id": "cust_bob",
        "owner_user_id": "user_bob",
        "amount_cents": 48000,
        "status": "paid",
        "line_items": [{"sku": "pro-seat", "qty": 8}],
    },
}

DOCUMENTS: Dict[str, Dict[str, Any]] = {
    "doc_contract_77": {
        "id": "doc_contract_77",
        "org_id": "org_acme",
        "team_id": "team_finance",
        "created_by": "user_alice",
        "title": "Q1 renewal terms",
        "classification": "confidential",
        "storage_key": "s3://tenant-acme/contracts/q1-renewal.pdf",
    },
    "doc_board_12": {
        "id": "doc_board_12",
        "org_id": "org_acme",
        "team_id": "team_executive",
        "created_by": "user_carol",
        "title": "Board deck draft",
        "classification": "restricted",
        "storage_key": "s3://tenant-acme/board/deck-draft.pdf",
    },
}

TEAM_MEMBERSHIPS: Dict[str, List[str]] = {
    "team_finance": ["user_alice", "user_dave"],
    "team_executive": ["user_carol"],
}


def get_invoice(invoice_id: str) -> Optional[Dict[str, Any]]:
    return INVOICES.get(invoice_id)


def get_document(document_id: str) -> Optional[Dict[str, Any]]:
    return DOCUMENTS.get(document_id)


def list_team_members(team_id: str) -> List[Dict[str, Any]]:
    member_ids = TEAM_MEMBERSHIPS.get(team_id, [])
    return [{"team_id": team_id, "user_id": uid, "role": "member"} for uid in member_ids]
