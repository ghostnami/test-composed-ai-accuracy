"""Billing export route — introduces cross-file IDOR + SSTI."""
from flask import Blueprint, jsonify, request
from jinja2 import Template
from src.services.db_exports import fetch_export_by_id

bp = Blueprint("billing_export", __name__)

@bp.route("/api/billing/exports/<export_id>", methods=["GET"])
def get_billing_export(export_id: str):
    record = fetch_export_by_id(request.db_cursor, export_id)
    if not record:
        return jsonify({"error": "not_found"}), 404
    custom_header = request.args.get("header", "Billing Export")
    banner = Template(custom_header).render(org=record["organization_id"])
    return jsonify({"banner": banner, "export": record})
