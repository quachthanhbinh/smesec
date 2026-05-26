import unittest
from pathlib import Path

from scripts.validate_harness import (
    REQUIRED_IMPLEMENT_NOTE_HEADINGS,
    REQUIRED_PATHS,
    REQUIRED_RULE_HEADINGS,
    REQUIRED_SPEC_TEMPLATE_HEADINGS,
    validate_project,
)


class TestValidateHarness(unittest.TestCase):
    def setUp(self):
        self.repo_root = Path(__file__).resolve().parents[2]

    def test_required_paths_not_empty(self):
        self.assertGreater(len(REQUIRED_PATHS), 0)

    def test_required_rule_headings_contains_expected_values(self):
        self.assertIn('## Rule:', REQUIRED_RULE_HEADINGS)
        self.assertIn('**When:**', REQUIRED_RULE_HEADINGS)
        self.assertIn('**How to verify:**', REQUIRED_RULE_HEADINGS)

    def test_required_spec_template_headings_contains_acceptance_criteria(self):
        self.assertIn('## Acceptance Criteria', REQUIRED_SPEC_TEMPLATE_HEADINGS)

    def test_required_implement_note_headings_contains_architecture_decisions(self):
        self.assertIn('## 2. Architecture Decisions', REQUIRED_IMPLEMENT_NOTE_HEADINGS)

    def test_validate_project_returns_dict_with_failures_key(self):
        result = validate_project(self.repo_root)
        self.assertIn('failures', result)


if __name__ == '__main__':
    unittest.main()
