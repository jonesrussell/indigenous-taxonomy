"""Indigenous Taxonomy — generated from canonical YAML schemas."""

from .version import TAXONOMY_VERSION, SCHEMA_HASH
from .categories import Category
from .regions import Region
from .dialects import DialectCode, ALL_DIALECTS, dialect_by_code

__all__ = [
    "TAXONOMY_VERSION",
    "SCHEMA_HASH",
    "Category",
    "Region",
    "DialectCode",
    "ALL_DIALECTS",
    "dialect_by_code",
]
