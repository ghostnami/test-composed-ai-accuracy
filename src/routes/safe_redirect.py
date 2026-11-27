"""Safe relative redirect helper (hard negative — properly mitigated)."""
from urllib.parse import urlparse
from flask import Blueprint, redirect, request

bp = Blueprint("safe_redirect", __name__)

@bp.route("/auth/continue", methods=["GET"])
def continue_redirect():
    target = request.args.get("next", "/dashboard")
    parsed = urlparse(target)
    if parsed.scheme or parsed.netloc or not target.startswith("/") or target.startswith("//"):
        return redirect("/dashboard")
    return redirect(target)
