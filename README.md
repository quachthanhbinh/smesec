# smesec

## Development Workflow

`BRAINSTORM → SPEC → PLAN → IMPLEMENT (TDD) → VERIFY`

## Harness Files

- `AGENTS.md`
- `docs/rules/`
- `docs/specs/_TEMPLATE/`
- `docs/audit/`
- `scripts/validate_harness.py`

## Validate Harness

Run:

```bash
python scripts/validate_harness.py
python -m unittest tests/harness/test_validate_harness.py -v
```
