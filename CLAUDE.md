# Indigenous Taxonomy

Canonical taxonomy contract for Indigenous content classification. Consumed by North Cloud (Go), Minoo (PHP), and indigenous-harvesters (Python).

## Commands

```
task generate       # Regenerate all packages from YAML
task validate       # Validate YAML against JSON Schema
task test           # Run all tests
task release        # Tag + push (CI publishes packages)
```

## Architecture

- schema/*.yaml — canonical source of truth (human-edited)
- schema/validation/*.schema.json — JSON Schema for YAML validation
- scripts/generate.py — reads YAML, writes Go/PHP/Python packages
- generated/ — output of generate.py (committed, not hand-edited)

## Rules

- NEVER hand-edit files in generated/ — always edit YAML then run task generate
- Adding a category/region/dialect = minor version bump
- Renaming/removing a slug = major version bump with deprecation window
- All slugs must be unique within their file
- Region slugs use colon-delimited hierarchy (canada:ontario:northern)
