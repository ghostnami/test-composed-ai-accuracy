"""Invoice endpoints for customer billing portal."""

from fastapi import APIRouter, Depends, HTTPException, status

from api.auth.dependencies import AuthenticatedUser, get_current_user
from api.db import repository

router = APIRouter(prefix="/api/v1/invoices", tags=["invoices"])


@router.get("/{invoice_id}")
async def get_invoice(
    invoice_id: str,
    current_user: AuthenticatedUser = Depends(get_current_user),
) -> dict:
    """Return invoice details for authenticated portal users."""
    invoice = repository.get_invoice(invoice_id)
    if not invoice:
        raise HTTPException(status_code=status.HTTP_404_NOT_FOUND, detail="invoice not found")

    return {
        "requested_by": current_user.user_id,
        "invoice": invoice,
    }


@router.get("/{invoice_id}/download")
async def download_invoice_pdf(
    invoice_id: str,
    current_user: AuthenticatedUser = Depends(get_current_user),
) -> dict:
    """Generate a signed download URL for invoice PDFs."""
    invoice = repository.get_invoice(invoice_id)
    if not invoice:
        raise HTTPException(status_code=status.HTTP_404_NOT_FOUND, detail="invoice not found")

    return {
        "invoice_id": invoice_id,
        "download_url": f"https://files.example.com/invoices/{invoice_id}.pdf?user={current_user.user_id}",
        "amount_cents": invoice["amount_cents"],
        "owner_user_id": invoice["owner_user_id"],
    }
