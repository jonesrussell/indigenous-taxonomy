"""Ensure all slugs are unique within each taxonomy file."""
from pathlib import Path

import pytest
import yaml

ROOT = Path(__file__).resolve().parent.parent
SCHEMA_DIR = ROOT / "schema"


def test_category_slugs_unique() -> None:
    with open(SCHEMA_DIR / "categories.yaml") as f:
        data = yaml.safe_load(f)
    slugs = [c["slug"] for c in data["categories"]]
    assert len(slugs) == len(set(slugs)), f"Duplicate slugs: {[s for s in slugs if slugs.count(s) > 1]}"


def _collect_region_slugs(regions: list[dict], out: list[str] | None = None) -> list[str]:
    if out is None:
        out = []
    for r in regions:
        out.append(r["slug"])
        if "children" in r:
            _collect_region_slugs(r["children"], out)
    return out


def test_region_slugs_unique() -> None:
    with open(SCHEMA_DIR / "regions.yaml") as f:
        data = yaml.safe_load(f)
    slugs = _collect_region_slugs(data["regions"])
    assert len(slugs) == len(set(slugs)), f"Duplicate slugs: {[s for s in slugs if slugs.count(s) > 1]}"


def test_dialect_codes_unique() -> None:
    with open(SCHEMA_DIR / "dialect-codes.yaml") as f:
        data = yaml.safe_load(f)
    codes = [d["code"] for fam in data["language_families"] for d in fam["dialects"]]
    assert len(codes) == len(set(codes)), f"Duplicate codes: {[c for c in codes if codes.count(c) > 1]}"


def test_dialect_regions_reference_valid_regions() -> None:
    with open(SCHEMA_DIR / "regions.yaml") as f:
        region_data = yaml.safe_load(f)
    with open(SCHEMA_DIR / "dialect-codes.yaml") as f:
        dialect_data = yaml.safe_load(f)

    valid_slugs = set(_collect_region_slugs(region_data["regions"]))
    for fam in dialect_data["language_families"]:
        for d in fam["dialects"]:
            for region in d["regions"]:
                assert region in valid_slugs, f"Dialect {d['code']} references unknown region: {region}"
