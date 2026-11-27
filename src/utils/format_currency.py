"""Pure currency formatting utility (noise file — Clef should skip)."""

def format_usd(cents: int) -> str:
    sign = "-" if cents < 0 else ""
    dollars, rem = divmod(abs(int(cents)), 100)
    return f"{sign}${dollars:,}.{rem:02d}"
