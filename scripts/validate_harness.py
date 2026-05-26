from pathlib import Path

REQUIRED_PATHS = [Path('AGENTS.md')]

REQUIRED_RULE_HEADINGS = ['## Rule:', '**When:**', '**How to verify:**']

REQUIRED_SPEC_TEMPLATE_HEADINGS = ['## Acceptance Criteria']

REQUIRED_IMPLEMENT_NOTE_HEADINGS = ['## 2. Architecture Decisions']


def validate_project(repo_root: Path) -> dict:
    return {'failures': []}
