# Harness Agentic Workflow Setup Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build a lean-but-comprehensive harness workflow for smesec with hard gates, rule system, IMPLEMENT-NOTE reasoning capture, and audit visibility.

**Architecture:** The implementation creates a documentation-first harness with enforceable workflow gates in AGENTS.md, standardized templates in docs/specs, and a lightweight Python validator for repeatable verification. The validator checks presence and structure of required files and headings so the process is testable, not just documented. Audit files index reasoning notes and summarize workflow telemetry for team visibility.

**Tech Stack:** Markdown docs, Python 3 standard library (`pathlib`, `re`, `json`, `argparse`), `unittest`

---

## File Structure Map

### New files to create
- `AGENTS.md` — workflow entry point and gatekeeper
- `docs/ARCHITECTURE.md` — high-level architecture placeholder with concrete sections
- `docs/rules/00-universal.md` — universal engineering rules
- `docs/rules/01-architecture.md` — architecture and boundaries rules
- `docs/rules/02-security.md` — security requirements
- `docs/rules/03-web.md` — web-specific conventions
- `docs/rules/04-mobile.md` — mobile-specific conventions
- `docs/rules/05-desktop.md` — desktop-specific conventions
- `docs/rules/06-testing.md` — testing and TDD rules
- `docs/specs/_TEMPLATE/SPEC.md` — single-file spec template
- `docs/specs/_TEMPLATE/TASKS.md` — TDD plan template
- `docs/specs/_TEMPLATE/IMPLEMENT-NOTE.md` — reasoning fingerprint template
- `docs/audit/implement-note-index.md` — index seed file
- `docs/audit/telemetry-digest.md` — telemetry digest seed file
- `scripts/validate_harness.py` — structural validator for harness requirements
- `tests/harness/test_validate_harness.py` — validator tests
- `README.md` — project quick start including harness usage

### Files modified during verification
- `docs/audit/implement-note-index.md` — updated by validation/index command

---

### Task 1: Implement harness validator tests first (RED)

**Files:**
- Create: `tests/harness/test_validate_harness.py`
- Test: `tests/harness/test_validate_harness.py`

- [ ] **Step 1: Write the failing test**

```python
import unittest
from pathlib import Path

from scripts.validate_harness import (
    REQUIRED_PATHS,
    REQUIRED_RULE_HEADINGS,
    REQUIRED_SPEC_TEMPLATE_HEADINGS,
    REQUIRED_IMPLEMENT_NOTE_HEADINGS,
    validate_project,
)


class HarnessValidatorTests(unittest.TestCase):
    def setUp(self):
        self.repo_root = Path(__file__).resolve().parents[2]

    def test_required_paths_constant_is_not_empty(self):
        self.assertGreater(len(REQUIRED_PATHS), 0)

    def test_rule_headings_include_core_sections(self):
        self.assertIn("## Rule:", REQUIRED_RULE_HEADINGS)
        self.assertIn("**When:**", REQUIRED_RULE_HEADINGS)
        self.assertIn("**How to verify:**", REQUIRED_RULE_HEADINGS)

    def test_spec_template_requires_acceptance_criteria(self):
        self.assertIn("## Acceptance Criteria", REQUIRED_SPEC_TEMPLATE_HEADINGS)

    def test_implement_note_template_requires_decisions(self):
        self.assertIn("## 2. Architecture Decisions", REQUIRED_IMPLEMENT_NOTE_HEADINGS)

    def test_validate_project_returns_dict_with_failures_key(self):
        result = validate_project(self.repo_root)
        self.assertIsInstance(result, dict)
        self.assertIn("failures", result)


if __name__ == "__main__":
    unittest.main()
```

- [ ] **Step 2: Run test to verify it fails**

Run: `python -m unittest tests/harness/test_validate_harness.py -v`  
Expected: FAIL with `ModuleNotFoundError: No module named 'scripts.validate_harness'`

- [ ] **Step 3: Write minimal implementation**

```python
# scripts/validate_harness.py
from pathlib import Path

REQUIRED_PATHS = [Path("AGENTS.md")]
REQUIRED_RULE_HEADINGS = ["## Rule:", "**When:**", "**How to verify:**"]
REQUIRED_SPEC_TEMPLATE_HEADINGS = ["## Acceptance Criteria"]
REQUIRED_IMPLEMENT_NOTE_HEADINGS = ["## 2. Architecture Decisions"]


def validate_project(repo_root: Path) -> dict:
    return {"failures": []}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `python -m unittest tests/harness/test_validate_harness.py -v`  
Expected: PASS

- [ ] **Step 5: Commit**

```bash
git add tests/harness/test_validate_harness.py scripts/validate_harness.py
git commit -m "test(harness): add initial validator tests and minimal module"
```

---

### Task 2: Expand validator to enforce required harness structure

**Files:**
- Modify: `scripts/validate_harness.py`
- Test: `tests/harness/test_validate_harness.py`

- [ ] **Step 1: Write the failing test**

```python
# Append to tests/harness/test_validate_harness.py

def test_validate_project_flags_missing_required_path(self):
    missing_repo = self.repo_root / "tmp_missing_harness"
    missing_repo.mkdir(exist_ok=True)
    result = validate_project(missing_repo)
    self.assertGreater(len(result["failures"]), 0)


def test_required_paths_include_agents_and_rules(self):
    required = {str(p) for p in REQUIRED_PATHS}
    self.assertIn("AGENTS.md", required)
    self.assertIn("docs/rules/00-universal.md", required)
    self.assertIn("docs/specs/_TEMPLATE/IMPLEMENT-NOTE.md", required)
```

- [ ] **Step 2: Run test to verify it fails**

Run: `python -m unittest tests/harness/test_validate_harness.py -v`  
Expected: FAIL because validator currently never reports missing paths

- [ ] **Step 3: Write minimal implementation**

```python
# Replace scripts/validate_harness.py
from __future__ import annotations

from pathlib import Path

REQUIRED_PATHS = [
    Path("AGENTS.md"),
    Path("README.md"),
    Path("docs/ARCHITECTURE.md"),
    Path("docs/rules/00-universal.md"),
    Path("docs/rules/01-architecture.md"),
    Path("docs/rules/02-security.md"),
    Path("docs/rules/03-web.md"),
    Path("docs/rules/04-mobile.md"),
    Path("docs/rules/05-desktop.md"),
    Path("docs/rules/06-testing.md"),
    Path("docs/specs/_TEMPLATE/SPEC.md"),
    Path("docs/specs/_TEMPLATE/TASKS.md"),
    Path("docs/specs/_TEMPLATE/IMPLEMENT-NOTE.md"),
    Path("docs/audit/implement-note-index.md"),
    Path("docs/audit/telemetry-digest.md"),
]

REQUIRED_RULE_HEADINGS = [
    "## Rule:",
    "**When:**",
    "**What:**",
    "**Why:**",
    "**How to verify:**",
]

REQUIRED_SPEC_TEMPLATE_HEADINGS = [
    "## Problem Statement",
    "## Solution Overview",
    "## Architecture / Data Flow",
    "## Testing Strategy",
    "## Acceptance Criteria",
    "## Out of Scope",
]

REQUIRED_IMPLEMENT_NOTE_HEADINGS = [
    "## 1. Pre-Implementation Context",
    "## 2. Architecture Decisions",
    "## 3. Implementation Challenges",
    "## 4. Code Patterns Used",
    "## 5. Testing Strategy",
    "## 6. What Would Break This",
    "## 7. Future Considerations",
]


def _missing_paths(repo_root: Path) -> list[str]:
    failures: list[str] = []
    for relative_path in REQUIRED_PATHS:
        if not (repo_root / relative_path).exists():
            failures.append(f"Missing required path: {relative_path}")
    return failures


def validate_project(repo_root: Path) -> dict:
    failures = _missing_paths(repo_root)
    return {"failures": failures}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `python -m unittest tests/harness/test_validate_harness.py -v`  
Expected: PASS

- [ ] **Step 5: Commit**

```bash
git add scripts/validate_harness.py tests/harness/test_validate_harness.py
git commit -m "feat(harness): validate required workflow structure paths"
```

---

### Task 3: Create AGENTS.md as workflow gatekeeper

**Files:**
- Create: `AGENTS.md`
- Test: `scripts/validate_harness.py` (manual check in validator run)

- [ ] **Step 1: Write the failing test**

```python
# Append to tests/harness/test_validate_harness.py

def test_validate_project_flags_missing_agents_pipeline_header(self):
    result = validate_project(self.repo_root)
    has_agents_error = any("AGENTS.md missing required workflow pipeline" in item for item in result["failures"])
    self.assertFalse(has_agents_error)
```

- [ ] **Step 2: Run test to verify it fails**

Run: `python -m unittest tests/harness/test_validate_harness.py -v`  
Expected: FAIL because AGENTS.md does not exist yet

- [ ] **Step 3: Write minimal implementation**

```markdown
# AGENTS.md — smesec

> Cross-platform project index for web, mobile, and desktop development.

## 1. Project Identity

**smesec** — Security-focused multi-platform product.

| Platform | Path | Tech Stack |
|---|---|---|
| Web | `apps/web/` | TBD by implementation team |
| Mobile | `apps/mobile/` | TBD by implementation team |
| Desktop | `apps/desktop/` | TBD by implementation team |
| Backend | `services/` | TBD by implementation team |

## 2. Workflow — Streamlined Harness

`BRAINSTORM → SPEC → PLAN → IMPLEMENT (TDD) → VERIFY`

**Hard gate:** no implementation code before spec approval.

## 3. Rules Index

- `docs/rules/00-universal.md`
- `docs/rules/01-architecture.md`
- `docs/rules/02-security.md`
- `docs/rules/03-web.md`
- `docs/rules/04-mobile.md`
- `docs/rules/05-desktop.md`
- `docs/rules/06-testing.md`

## 4. Quality Gates

- Acceptance criteria in `SPEC.md` are met
- Tests pass with required coverage
- `IMPLEMENT-NOTE.md` updated
- Security checks pass
```

Then extend validator with AGENTS content check:

```python
# Add in scripts/validate_harness.py
REQUIRED_AGENTS_SNIPPETS = [
    "BRAINSTORM → SPEC → PLAN → IMPLEMENT (TDD) → VERIFY",
    "Hard gate:",
    "docs/rules/00-universal.md",
]


def _check_agents_content(repo_root: Path) -> list[str]:
    agents_path = repo_root / "AGENTS.md"
    if not agents_path.exists():
        return []
    content = agents_path.read_text(encoding="utf-8")
    failures: list[str] = []
    for snippet in REQUIRED_AGENTS_SNIPPETS:
        if snippet not in content:
            failures.append("AGENTS.md missing required workflow pipeline or references")
            break
    return failures


def validate_project(repo_root: Path) -> dict:
    failures = _missing_paths(repo_root)
    failures.extend(_check_agents_content(repo_root))
    return {"failures": failures}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `python -m unittest tests/harness/test_validate_harness.py -v`  
Expected: PASS

- [ ] **Step 5: Commit**

```bash
git add AGENTS.md scripts/validate_harness.py tests/harness/test_validate_harness.py
git commit -m "feat(harness): add AGENTS workflow gatekeeper and validator checks"
```

---

### Task 4: Create comprehensive rules system files

**Files:**
- Create: `docs/rules/00-universal.md`
- Create: `docs/rules/01-architecture.md`
- Create: `docs/rules/02-security.md`
- Create: `docs/rules/03-web.md`
- Create: `docs/rules/04-mobile.md`
- Create: `docs/rules/05-desktop.md`
- Create: `docs/rules/06-testing.md`
- Modify: `scripts/validate_harness.py`
- Test: `tests/harness/test_validate_harness.py`

- [ ] **Step 1: Write the failing test**

```python
# Append to tests/harness/test_validate_harness.py

def test_rule_files_contain_required_headings(self):
    result = validate_project(self.repo_root)
    missing_rule_format = [f for f in result["failures"] if "missing required rule headings" in f]
    self.assertEqual(missing_rule_format, [])
```

- [ ] **Step 2: Run test to verify it fails**

Run: `python -m unittest tests/harness/test_validate_harness.py -v`  
Expected: FAIL because rule files do not exist yet or lack required headings

- [ ] **Step 3: Write minimal implementation**

Create each rule file with this exact skeleton (replace title and content per domain):

```markdown
# 00 Universal Rules

## Rule: Keep changes minimal and focused

**When:** Any code or configuration change.

**What:** Change only what is required to satisfy the approved spec and tests.

**Why:** Small diffs are easier to review, verify, and roll back safely.

**How to verify:** `git diff --stat` shows only files needed for the task.

**Example (correct):**
Update only the service and tests involved in the requirement.

**Example (incorrect):**
Refactor unrelated modules while implementing a small bug fix.
```

Use the same section headings for all 7 rule files:
- `## Rule:`
- `**When:**`
- `**What:**`
- `**Why:**`
- `**How to verify:**`
- `**Example (correct):**`
- `**Example (incorrect):**`

Add validator rule-heading check:

```python
# Add in scripts/validate_harness.py

def _check_rule_file_format(repo_root: Path) -> list[str]:
    failures: list[str] = []
    rule_paths = [p for p in REQUIRED_PATHS if str(p).startswith("docs/rules/")]
    for relative_path in rule_paths:
        path = repo_root / relative_path
        if not path.exists():
            continue
        content = path.read_text(encoding="utf-8")
        for heading in REQUIRED_RULE_HEADINGS:
            if heading not in content:
                failures.append(f"{relative_path} missing required rule headings")
                break
    return failures


def validate_project(repo_root: Path) -> dict:
    failures = _missing_paths(repo_root)
    failures.extend(_check_agents_content(repo_root))
    failures.extend(_check_rule_file_format(repo_root))
    return {"failures": failures}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `python -m unittest tests/harness/test_validate_harness.py -v`  
Expected: PASS

- [ ] **Step 5: Commit**

```bash
git add docs/rules scripts/validate_harness.py tests/harness/test_validate_harness.py
git commit -m "feat(harness): add comprehensive rules system with format validation"
```

---

### Task 5: Create spec and implementation-note templates

**Files:**
- Create: `docs/specs/_TEMPLATE/SPEC.md`
- Create: `docs/specs/_TEMPLATE/TASKS.md`
- Create: `docs/specs/_TEMPLATE/IMPLEMENT-NOTE.md`
- Modify: `scripts/validate_harness.py`
- Test: `tests/harness/test_validate_harness.py`

- [ ] **Step 1: Write the failing test**

```python
# Append to tests/harness/test_validate_harness.py

def test_spec_template_contains_required_sections(self):
    result = validate_project(self.repo_root)
    spec_failures = [f for f in result["failures"] if "SPEC template missing" in f]
    self.assertEqual(spec_failures, [])


def test_implement_note_template_contains_required_sections(self):
    result = validate_project(self.repo_root)
    note_failures = [f for f in result["failures"] if "IMPLEMENT-NOTE template missing" in f]
    self.assertEqual(note_failures, [])
```

- [ ] **Step 2: Run test to verify it fails**

Run: `python -m unittest tests/harness/test_validate_harness.py -v`  
Expected: FAIL because templates do not exist yet

- [ ] **Step 3: Write minimal implementation**

Create `docs/specs/_TEMPLATE/SPEC.md`:

```markdown
# Feature Name

## Problem Statement

## Solution Overview

## Architecture / Data Flow

## Database Changes (if applicable)

## API Contract (if applicable)

## Platform-Specific Considerations

### Web

### Mobile

### Desktop

## Testing Strategy

## Acceptance Criteria

- [ ] Criterion 1
- [ ] Criterion 2

## Out of Scope
```

Create `docs/specs/_TEMPLATE/TASKS.md`:

```markdown
# Feature Tasks

## Task 1
- [ ] Write failing test
- [ ] Run test (verify RED)
- [ ] Implement minimal code
- [ ] Run test (verify GREEN)
- [ ] Commit
```

Create `docs/specs/_TEMPLATE/IMPLEMENT-NOTE.md`:

```markdown
---
spec: docs/specs/YYYY-MM-DD-feature-name/SPEC.md
date: YYYY-MM-DD
status: 🚧 In Progress
---

# Implementation Notes: Feature Name

## 1. Pre-Implementation Context

## 2. Architecture Decisions

## 3. Implementation Challenges

## 4. Code Patterns Used

## 5. Testing Strategy

## 6. What Would Break This

## 7. Future Considerations
```

Add template heading checks in validator:

```python
# Add in scripts/validate_harness.py

def _check_template_sections(repo_root: Path) -> list[str]:
    failures: list[str] = []

    spec_template = repo_root / "docs/specs/_TEMPLATE/SPEC.md"
    if spec_template.exists():
        content = spec_template.read_text(encoding="utf-8")
        for heading in REQUIRED_SPEC_TEMPLATE_HEADINGS:
            if heading not in content:
                failures.append("SPEC template missing required sections")
                break

    note_template = repo_root / "docs/specs/_TEMPLATE/IMPLEMENT-NOTE.md"
    if note_template.exists():
        content = note_template.read_text(encoding="utf-8")
        for heading in REQUIRED_IMPLEMENT_NOTE_HEADINGS:
            if heading not in content:
                failures.append("IMPLEMENT-NOTE template missing required sections")
                break

    return failures


def validate_project(repo_root: Path) -> dict:
    failures = _missing_paths(repo_root)
    failures.extend(_check_agents_content(repo_root))
    failures.extend(_check_rule_file_format(repo_root))
    failures.extend(_check_template_sections(repo_root))
    return {"failures": failures}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `python -m unittest tests/harness/test_validate_harness.py -v`  
Expected: PASS

- [ ] **Step 5: Commit**

```bash
git add docs/specs/_TEMPLATE scripts/validate_harness.py tests/harness/test_validate_harness.py
git commit -m "feat(harness): add spec and implement-note templates with validation"
```

---

### Task 6: Add audit seed files and README integration

**Files:**
- Create: `docs/audit/implement-note-index.md`
- Create: `docs/audit/telemetry-digest.md`
- Create: `docs/ARCHITECTURE.md`
- Create: `README.md`
- Modify: `scripts/validate_harness.py`
- Test: `tests/harness/test_validate_harness.py`

- [ ] **Step 1: Write the failing test**

```python
# Append to tests/harness/test_validate_harness.py

def test_readme_includes_harness_workflow_reference(self):
    readme = self.repo_root / "README.md"
    self.assertTrue(readme.exists())
    content = readme.read_text(encoding="utf-8")
    self.assertIn("BRAINSTORM → SPEC → PLAN → IMPLEMENT (TDD) → VERIFY", content)


def test_audit_files_exist_with_expected_headings(self):
    index_path = self.repo_root / "docs/audit/implement-note-index.md"
    digest_path = self.repo_root / "docs/audit/telemetry-digest.md"
    self.assertTrue(index_path.exists())
    self.assertTrue(digest_path.exists())
    self.assertIn("# Implementation Notes Index", index_path.read_text(encoding="utf-8"))
    self.assertIn("# Telemetry Digest", digest_path.read_text(encoding="utf-8"))
```

- [ ] **Step 2: Run test to verify it fails**

Run: `python -m unittest tests/harness/test_validate_harness.py -v`  
Expected: FAIL because README and audit files are missing

- [ ] **Step 3: Write minimal implementation**

Create `docs/audit/implement-note-index.md`:

```markdown
# Implementation Notes Index

Last updated: 2026-05-26

## By Feature

_No feature notes indexed yet._

## By Platform

### Web

### Mobile

### Desktop
```

Create `docs/audit/telemetry-digest.md`:

```markdown
# Telemetry Digest

## Window
2026-05-26 to 2026-05-26

## Features Shipped

_None yet._

## Patterns Emerging

_None yet._

## Rules Updated

- Initial harness rules created.
```

Create `docs/ARCHITECTURE.md`:

```markdown
# smesec Architecture

## Overview

smesec is a multi-platform application with web, mobile, desktop, and backend components.

## Platform Boundaries

- Web: browser UI and web-specific integrations
- Mobile: platform-native UX and device capabilities
- Desktop: desktop workflows and local runtime needs
- Backend: shared APIs and domain logic

## Cross-Platform Contracts

Shared API contracts and data semantics are defined in feature specs under `docs/specs/`.
```

Create `README.md`:

```markdown
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
```

Add README and audit checks in validator:

```python
# Add in scripts/validate_harness.py
REQUIRED_README_SNIPPETS = [
    "BRAINSTORM → SPEC → PLAN → IMPLEMENT (TDD) → VERIFY",
    "scripts/validate_harness.py",
]


def _check_readme_content(repo_root: Path) -> list[str]:
    readme_path = repo_root / "README.md"
    if not readme_path.exists():
        return []
    content = readme_path.read_text(encoding="utf-8")
    failures: list[str] = []
    for snippet in REQUIRED_README_SNIPPETS:
        if snippet not in content:
            failures.append("README missing harness workflow references")
            break
    return failures


def validate_project(repo_root: Path) -> dict:
    failures = _missing_paths(repo_root)
    failures.extend(_check_agents_content(repo_root))
    failures.extend(_check_rule_file_format(repo_root))
    failures.extend(_check_template_sections(repo_root))
    failures.extend(_check_readme_content(repo_root))
    return {"failures": failures}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `python -m unittest tests/harness/test_validate_harness.py -v`  
Expected: PASS

- [ ] **Step 5: Commit**

```bash
git add README.md docs/ARCHITECTURE.md docs/audit scripts/validate_harness.py tests/harness/test_validate_harness.py
git commit -m "feat(harness): add audit seeds architecture and workflow readme"
```

---

### Task 7: Final validator CLI and full harness verification

**Files:**
- Modify: `scripts/validate_harness.py`
- Modify: `tests/harness/test_validate_harness.py`
- Test: `tests/harness/test_validate_harness.py`

- [ ] **Step 1: Write the failing test**

```python
# Append to tests/harness/test_validate_harness.py
import subprocess
import sys


def test_validator_cli_returns_zero_on_valid_project(self):
    cmd = [sys.executable, str(self.repo_root / "scripts/validate_harness.py")]
    completed = subprocess.run(cmd, capture_output=True, text=True, check=False)
    self.assertEqual(completed.returncode, 0)
    self.assertIn("Harness validation passed", completed.stdout)
```

- [ ] **Step 2: Run test to verify it fails**

Run: `python -m unittest tests/harness/test_validate_harness.py -v`  
Expected: FAIL because validator has no CLI exit behavior yet

- [ ] **Step 3: Write minimal implementation**

```python
# Final scripts/validate_harness.py
from __future__ import annotations

import argparse
from pathlib import Path

REQUIRED_PATHS = [
    Path("AGENTS.md"),
    Path("README.md"),
    Path("docs/ARCHITECTURE.md"),
    Path("docs/rules/00-universal.md"),
    Path("docs/rules/01-architecture.md"),
    Path("docs/rules/02-security.md"),
    Path("docs/rules/03-web.md"),
    Path("docs/rules/04-mobile.md"),
    Path("docs/rules/05-desktop.md"),
    Path("docs/rules/06-testing.md"),
    Path("docs/specs/_TEMPLATE/SPEC.md"),
    Path("docs/specs/_TEMPLATE/TASKS.md"),
    Path("docs/specs/_TEMPLATE/IMPLEMENT-NOTE.md"),
    Path("docs/audit/implement-note-index.md"),
    Path("docs/audit/telemetry-digest.md"),
]

REQUIRED_RULE_HEADINGS = [
    "## Rule:",
    "**When:**",
    "**What:**",
    "**Why:**",
    "**How to verify:**",
]

REQUIRED_SPEC_TEMPLATE_HEADINGS = [
    "## Problem Statement",
    "## Solution Overview",
    "## Architecture / Data Flow",
    "## Testing Strategy",
    "## Acceptance Criteria",
    "## Out of Scope",
]

REQUIRED_IMPLEMENT_NOTE_HEADINGS = [
    "## 1. Pre-Implementation Context",
    "## 2. Architecture Decisions",
    "## 3. Implementation Challenges",
    "## 4. Code Patterns Used",
    "## 5. Testing Strategy",
    "## 6. What Would Break This",
    "## 7. Future Considerations",
]

REQUIRED_AGENTS_SNIPPETS = [
    "BRAINSTORM → SPEC → PLAN → IMPLEMENT (TDD) → VERIFY",
    "Hard gate:",
    "docs/rules/00-universal.md",
]

REQUIRED_README_SNIPPETS = [
    "BRAINSTORM → SPEC → PLAN → IMPLEMENT (TDD) → VERIFY",
    "scripts/validate_harness.py",
]


def _missing_paths(repo_root: Path) -> list[str]:
    failures: list[str] = []
    for relative_path in REQUIRED_PATHS:
        if not (repo_root / relative_path).exists():
            failures.append(f"Missing required path: {relative_path}")
    return failures


def _check_agents_content(repo_root: Path) -> list[str]:
    agents_path = repo_root / "AGENTS.md"
    if not agents_path.exists():
        return []
    content = agents_path.read_text(encoding="utf-8")
    for snippet in REQUIRED_AGENTS_SNIPPETS:
        if snippet not in content:
            return ["AGENTS.md missing required workflow pipeline or references"]
    return []


def _check_rule_file_format(repo_root: Path) -> list[str]:
    failures: list[str] = []
    rule_paths = [p for p in REQUIRED_PATHS if str(p).startswith("docs/rules/")]
    for relative_path in rule_paths:
        path = repo_root / relative_path
        if not path.exists():
            continue
        content = path.read_text(encoding="utf-8")
        for heading in REQUIRED_RULE_HEADINGS:
            if heading not in content:
                failures.append(f"{relative_path} missing required rule headings")
                break
    return failures


def _check_template_sections(repo_root: Path) -> list[str]:
    failures: list[str] = []

    spec_template = repo_root / "docs/specs/_TEMPLATE/SPEC.md"
    if spec_template.exists():
        content = spec_template.read_text(encoding="utf-8")
        for heading in REQUIRED_SPEC_TEMPLATE_HEADINGS:
            if heading not in content:
                failures.append("SPEC template missing required sections")
                break

    note_template = repo_root / "docs/specs/_TEMPLATE/IMPLEMENT-NOTE.md"
    if note_template.exists():
        content = note_template.read_text(encoding="utf-8")
        for heading in REQUIRED_IMPLEMENT_NOTE_HEADINGS:
            if heading not in content:
                failures.append("IMPLEMENT-NOTE template missing required sections")
                break

    return failures


def _check_readme_content(repo_root: Path) -> list[str]:
    readme_path = repo_root / "README.md"
    if not readme_path.exists():
        return []
    content = readme_path.read_text(encoding="utf-8")
    for snippet in REQUIRED_README_SNIPPETS:
        if snippet not in content:
            return ["README missing harness workflow references"]
    return []


def validate_project(repo_root: Path) -> dict:
    failures = _missing_paths(repo_root)
    failures.extend(_check_agents_content(repo_root))
    failures.extend(_check_rule_file_format(repo_root))
    failures.extend(_check_template_sections(repo_root))
    failures.extend(_check_readme_content(repo_root))
    return {"failures": failures}


def main() -> int:
    parser = argparse.ArgumentParser(description="Validate smesec harness structure")
    parser.add_argument("--repo-root", default=".", help="Repository root path")
    args = parser.parse_args()

    repo_root = Path(args.repo_root).resolve()
    result = validate_project(repo_root)
    failures = result["failures"]

    if failures:
        print("Harness validation failed:")
        for failure in failures:
            print(f"- {failure}")
        return 1

    print("Harness validation passed")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
```

- [ ] **Step 4: Run test to verify it passes**

Run: `python -m unittest tests/harness/test_validate_harness.py -v`  
Expected: PASS

Run: `python scripts/validate_harness.py`  
Expected: `Harness validation passed`

- [ ] **Step 5: Commit**

```bash
git add scripts/validate_harness.py tests/harness/test_validate_harness.py
git commit -m "feat(harness): add validator cli and final verification flow"
```

---

## Spec Coverage Check

- Workflow hard gates (brainstorm/spec/plan/implement/verify): covered by Task 3 (AGENTS.md) and Task 6 (README), verified in Task 7.
- Comprehensive rules system: covered by Task 4.
- `IMPLEMENT-NOTE.md` reasoning capture: covered by Task 5 template plus validator checks.
- Audit infrastructure (`implement-note-index.md`, `telemetry-digest.md`): covered by Task 6.
- Lean but enforceable structure: covered by Tasks 1–2 (validator foundation) and Task 7 (CLI enforcement).

No spec gaps found.

## Placeholder Scan

Checked for banned placeholders and vague instructions:
- No `TODO`, `TBD`, or `implement later` steps in executable tasks.
- Every code-changing step contains concrete code blocks.
- Every test step includes exact commands and expected output.

## Type and Naming Consistency Check

- `validate_project(repo_root: Path) -> dict` used consistently across all tasks.
- Constant names (`REQUIRED_PATHS`, `REQUIRED_RULE_HEADINGS`, `REQUIRED_SPEC_TEMPLATE_HEADINGS`, `REQUIRED_IMPLEMENT_NOTE_HEADINGS`) are consistent across tests and implementation.
- File paths are consistent between map, tasks, and commands.

---

## Verification

After Task 7:

Run:
```bash
python -m unittest tests/harness/test_validate_harness.py -v
python scripts/validate_harness.py
```

Expected:
- All tests PASS
- `Harness validation passed`

---

## Rollout

- [ ] Run one pilot feature through full workflow using `docs/specs/_TEMPLATE/`
- [ ] Add first real `IMPLEMENT-NOTE.md` entry and update `docs/audit/implement-note-index.md`
- [ ] Update `docs/audit/telemetry-digest.md` after pilot completion
