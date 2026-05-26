from __future__ import annotations

import argparse
import sys
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
    "**Example (correct):**",
    "**Example (incorrect):**",
]

REQUIRED_SPEC_TEMPLATE_HEADINGS = [
    "## Problem Statement",
    "## Solution Overview",
    "## Architecture / Data Flow",
    "## Database Changes (if applicable)",
    "## API Contract (if applicable)",
    "## Platform-Specific Considerations",
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
    "## Workflow",
    "BRAINSTORM → SPEC → PLAN → IMPLEMENT (TDD) → VERIFY",
    "## Rules Index",
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
    failures: list[str] = []
    agents_path = repo_root / "AGENTS.md"
    if not agents_path.exists():
        return failures

    content = agents_path.read_text(encoding="utf-8")
    for snippet in REQUIRED_AGENTS_SNIPPETS:
        if snippet not in content:
            failures.append(f"AGENTS.md missing required snippet: {snippet}")
    return failures


def _check_rule_file_format(repo_root: Path) -> list[str]:
    failures: list[str] = []
    rule_files = [
        Path("docs/rules/00-universal.md"),
        Path("docs/rules/01-architecture.md"),
        Path("docs/rules/02-security.md"),
        Path("docs/rules/03-web.md"),
        Path("docs/rules/04-mobile.md"),
        Path("docs/rules/05-desktop.md"),
        Path("docs/rules/06-testing.md"),
    ]

    for rule_file in rule_files:
        file_path = repo_root / rule_file
        if not file_path.exists():
            continue

        content = file_path.read_text(encoding="utf-8")
        for heading in REQUIRED_RULE_HEADINGS:
            if heading not in content:
                failures.append(f"{rule_file} missing required heading: {heading}")

    return failures


def _check_template_sections(repo_root: Path) -> list[str]:
    failures: list[str] = []

    # Check SPEC.md template
    spec_template_path = repo_root / "docs/specs/_TEMPLATE/SPEC.md"
    if spec_template_path.exists():
        content = spec_template_path.read_text(encoding="utf-8")
        for heading in REQUIRED_SPEC_TEMPLATE_HEADINGS:
            if heading not in content:
                failures.append(f"docs/specs/_TEMPLATE/SPEC.md missing required heading: {heading}")

    # Check IMPLEMENT-NOTE.md template
    implement_note_path = repo_root / "docs/specs/_TEMPLATE/IMPLEMENT-NOTE.md"
    if implement_note_path.exists():
        content = implement_note_path.read_text(encoding="utf-8")
        for heading in REQUIRED_IMPLEMENT_NOTE_HEADINGS:
            if heading not in content:
                failures.append(f"docs/specs/_TEMPLATE/IMPLEMENT-NOTE.md missing required heading: {heading}")

    return failures


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


def main() -> int:
    parser = argparse.ArgumentParser(
        description="Validate harness workflow structure and content"
    )
    parser.add_argument(
        "--repo-root",
        type=Path,
        default=Path.cwd(),
        help="Path to repository root (defaults to current directory)",
    )
    args = parser.parse_args()

    result = validate_project(args.repo_root)
    failures = result["failures"]

    if not failures:
        print("Harness validation passed")
        return 0
    else:
        print("Harness validation failed:")
        for failure in failures:
            print(f"  - {failure}")
        return 1


if __name__ == "__main__":
    sys.exit(main())
