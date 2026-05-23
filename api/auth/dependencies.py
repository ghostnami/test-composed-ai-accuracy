"""Authentication dependencies shared across API routes."""

from dataclasses import dataclass
from typing import Optional

from fastapi import Header, HTTPException, status


@dataclass(frozen=True)
class AuthenticatedUser:
    user_id: str
    email: str
    org_id: str
    role: str


def get_current_user(x_user_id: str = Header(..., alias="X-User-Id")) -> AuthenticatedUser:
    """Resolve the caller from gateway-injected identity headers."""
    if not x_user_id:
        raise HTTPException(status_code=status.HTTP_401_UNAUTHORIZED, detail="missing identity")

    # In production this would hydrate from session/JWT claims.
    return AuthenticatedUser(
        user_id=x_user_id,
        email=f"{x_user_id}@example.com",
        org_id="org_demo",
        role="member",
    )


def require_admin(user: AuthenticatedUser) -> AuthenticatedUser:
    if user.role != "admin":
        raise HTTPException(status_code=status.HTTP_403_FORBIDDEN, detail="admin required")
    return user
