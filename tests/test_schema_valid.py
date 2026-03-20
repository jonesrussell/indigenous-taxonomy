"""Validate YAML files against their JSON Schema definitions."""
import json
from pathlib import Path

import jsonschema
import pytest
import yaml

ROOT = Path(__file__).resolve().parent.parent
SCHEMA_DIR = ROOT / "schema"
VALIDATION_DIR = SCHEMA_DIR / "validation"

YAML_SCHEMAS = [
    ("categories", "categories.schema.json"),
    ("regions", "regions.schema.json"),
    ("dialect-codes", "dialect-codes.schema.json"),
]


@pytest.mark.parametrize("yaml_name,schema_file", YAML_SCHEMAS)
def test_yaml_validates_against_schema(yaml_name: str, schema_file: str) -> None:
    with open(SCHEMA_DIR / f"{yaml_name}.yaml") as f:
        data = yaml.safe_load(f)
    with open(VALIDATION_DIR / schema_file) as f:
        schema = json.load(f)
    jsonschema.validate(data, schema)
