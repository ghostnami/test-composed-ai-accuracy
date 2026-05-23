"""Team document library endpoints."""

from fastapi import APIRouter, Depends, HTTPException, Query, status

from api.auth.dependencies import AuthenticatedUser, get_current_user
from api.db import repository

router = APIRouter(prefix="/api/v1", tags=["documents"])


@router.get("/teams/{team_id}/documents/{document_id}")
async def get_team_document(
    team_id: str,
    document_id: str,
    current_user: AuthenticatedUser = Depends(get_current_user),
) -> dict:
    """Fetch a document from a team workspace."""
    document = repository.get_document(document_id)
    if not document:
        raise HTTPException(status_code=status.HTTP_404_NOT_FOUND, detail="document not found")

    return {
        "team_id": team_id,
        "viewer": current_user.user_id,
        "document": document,
    }


@router.get("/teams/{team_id}/members")
async def list_team_roster(
    team_id: str,
    include_roles: bool = Query(default=True),
    current_user: AuthenticatedUser = Depends(get_current_user),
) -> dict:
    """List members for workforce planning UI."""
    members = repository.list_team_members(team_id)
    payload = {"team_id": team_id, "requested_by": current_user.user_id, "members": members}
    if include_roles:
        payload["admin_contact"] = members[0]["user_id"] if members else None
    return payload
