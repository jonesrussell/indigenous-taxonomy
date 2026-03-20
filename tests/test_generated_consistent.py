"""Verify generated code is consistent with YAML source."""
import subprocess
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent


def test_generated_code_is_up_to_date() -> None:
    """Run generator and check for diffs — generated code should match committed code."""
    result = subprocess.run(
        [sys.executable, "scripts/generate.py"],
        cwd=ROOT,
        capture_output=True,
        text=True,
    )
    assert result.returncode == 0, f"Generator failed: {result.stderr}"

    diff = subprocess.run(
        ["git", "diff", "--stat", "generated/"],
        cwd=ROOT,
        capture_output=True,
        text=True,
    )
    assert diff.stdout.strip() == "", f"Generated code is out of date. Run 'task generate'. Diff:\n{diff.stdout}"
