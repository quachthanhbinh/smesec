import subprocess
import sys
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

    def test_validate_project_flags_missing_required_path(self):
        missing_repo = self.repo_root / "tmp_missing_harness"
        missing_repo.mkdir(exist_ok=True)
        self.addCleanup(lambda: missing_repo.rmdir() if missing_repo.exists() else None)
        result = validate_project(missing_repo)
        self.assertGreater(len(result["failures"]), 0)

    def test_required_paths_include_agents_and_rules(self):
        required = {p.as_posix() for p in REQUIRED_PATHS}
        self.assertIn("AGENTS.md", required)
        self.assertIn("docs/rules/00-universal.md", required)
        self.assertIn("docs/specs/_TEMPLATE/IMPLEMENT-NOTE.md", required)

    def test_validate_project_flags_missing_agents_pipeline_header(self):
        test_repo = self.repo_root / "tmp_test_agents"
        test_repo.mkdir(exist_ok=True)
        agents_file = test_repo / "AGENTS.md"
        agents_file.write_text("# AGENTS\n\nSome content without pipeline")
        self.addCleanup(lambda: test_repo.rmdir() if test_repo.exists() else None)
        self.addCleanup(lambda: agents_file.unlink() if agents_file.exists() else None)
        result = validate_project(test_repo)
        self.assertTrue(any("BRAINSTORM" in f for f in result["failures"]))

    def test_rule_files_contain_required_headings(self):
        test_repo = self.repo_root / "tmp_test_rules"
        test_repo.mkdir(exist_ok=True)
        rules_dir = test_repo / "docs" / "rules"
        rules_dir.mkdir(parents=True, exist_ok=True)

        # Create a rule file missing required headings
        rule_file = rules_dir / "00-universal.md"
        rule_file.write_text("# 00 Universal Rules\n\n## Rule: Some rule\n\nMissing other headings")

        def cleanup():
            if rule_file.exists():
                rule_file.unlink()
            if rules_dir.exists():
                rules_dir.rmdir()
            if (test_repo / "docs").exists():
                (test_repo / "docs").rmdir()
            if test_repo.exists():
                test_repo.rmdir()

        self.addCleanup(cleanup)

        result = validate_project(test_repo)
        # Should flag missing headings like **When:**, **What:**, etc.
        rule_failures = [f for f in result["failures"] if "00-universal.md" in f and "missing required heading" in f]
        self.assertGreater(len(rule_failures), 0)

    def test_spec_template_contains_required_sections(self):
        test_repo = self.repo_root / "tmp_test_spec_template"
        test_repo.mkdir(exist_ok=True)
        template_dir = test_repo / "docs" / "specs" / "_TEMPLATE"
        template_dir.mkdir(parents=True, exist_ok=True)

        # Create a SPEC.md template missing required sections
        spec_file = template_dir / "SPEC.md"
        spec_file.write_text("# Spec Template\n\n## Problem Statement\n\nMissing other sections")

        def cleanup():
            if spec_file.exists():
                spec_file.unlink()
            if template_dir.exists():
                template_dir.rmdir()
            if (test_repo / "docs" / "specs").exists():
                (test_repo / "docs" / "specs").rmdir()
            if (test_repo / "docs").exists():
                (test_repo / "docs").rmdir()
            if test_repo.exists():
                test_repo.rmdir()

        self.addCleanup(cleanup)

        result = validate_project(test_repo)
        # Should flag missing sections like ## Acceptance Criteria, ## Out of Scope, etc.
        spec_failures = [f for f in result["failures"] if "SPEC.md" in f and "missing required heading" in f]
        self.assertGreater(len(spec_failures), 0)

    def test_implement_note_template_contains_required_sections(self):
        test_repo = self.repo_root / "tmp_test_implement_note_template"
        test_repo.mkdir(exist_ok=True)
        template_dir = test_repo / "docs" / "specs" / "_TEMPLATE"
        template_dir.mkdir(parents=True, exist_ok=True)

        # Create an IMPLEMENT-NOTE.md template missing required sections
        implement_note_file = template_dir / "IMPLEMENT-NOTE.md"
        implement_note_file.write_text("# Implementation Note\n\n## 1. Pre-Implementation Context\n\nMissing other sections")

        def cleanup():
            if implement_note_file.exists():
                implement_note_file.unlink()
            if template_dir.exists():
                template_dir.rmdir()
            if (test_repo / "docs" / "specs").exists():
                (test_repo / "docs" / "specs").rmdir()
            if (test_repo / "docs").exists():
                (test_repo / "docs").rmdir()
            if test_repo.exists():
                test_repo.rmdir()

        self.addCleanup(cleanup)

        result = validate_project(test_repo)
        # Should flag missing sections like ## 2. Architecture Decisions, ## 7. Future Considerations, etc.
        implement_note_failures = [f for f in result["failures"] if "IMPLEMENT-NOTE.md" in f and "missing required heading" in f]
        self.assertGreater(len(implement_note_failures), 0)

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

    def test_validator_cli_returns_zero_on_valid_project(self):
        result = subprocess.run(
            [sys.executable, "-m", "scripts.validate_harness", "--repo-root", str(self.repo_root)],
            capture_output=True,
            text=True,
            cwd=self.repo_root
        )
        self.assertEqual(result.returncode, 0)
        self.assertIn("Harness validation passed", result.stdout)


if __name__ == '__main__':
    unittest.main()
