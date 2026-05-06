# Changelog

## [v1.1.0] — 2026-05-06

### Added
- **Treaty namespace** — 11 numbered-treaty constants (`treaty:1` through `treaty:11`) covering Canadian numbered treaties. New `Treaty` Go type alias, `AllTreaties` slice, `IsValidTreaty(string) bool` helper.
- **City-level region children** under provinces: `canada:manitoba:winnipeg`, `canada:ontario:toronto`, `canada:ontario:ottawa`, `canada:british-columbia:vancouver`, `canada:alberta:calgary`, `canada:saskatchewan:saskatoon`.
- **First Nations community samples** (with `-fn` suffix) — at least one per active treaty area, including `canada:manitoba:sagkeeng-fn`, `canada:alberta:siksika-fn`, `canada:saskatchewan:beardy-okimasis-fn`, `canada:ontario:mississaugas-of-the-credit-fn`, plus partial-marker entries for British Columbia, Quebec, Atlantic, and Yukon.
- **`ParentRegion(Region) (Region, bool)`** function — walks the region hierarchy upward via a generated `regionParents` lookup table sourced from the YAML `children:` structure.

### Changed
- Generator (`scripts/generate.py`) now emits `treaties.go` and a sorted `regionParents` map in `regions.go`.

### Removed
- Nothing (additive release).

### Notes
- Released to support `alert-crawler` (north-cloud monorepo, mission `community-alert-pipeline-01KQZC7A`) for sovereignty-aware scope resolution in community safety alerts.

## [1.0.0] — 2026-03-19

### Added
- Initial taxonomy: 10 categories, 13 Canadian regions, 10 dialect codes
- JSON Schema validation for all YAML files
- Python generator producing Go, PHP, and Python packages
- CI workflows for validation and release
