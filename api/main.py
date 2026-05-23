"""Minimal FastAPI app wiring for billing + document routes."""

from fastapi import FastAPI

from api.routes import documents, invoices

app = FastAPI(title="Accuracy Validation API", version="0.2.0")
app.include_router(invoices.router)
app.include_router(documents.router)


@app.get("/health")
async def health() -> dict:
    return {"status": "ok"}
