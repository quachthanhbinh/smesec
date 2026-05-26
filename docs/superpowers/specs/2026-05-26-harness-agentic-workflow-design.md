# smesec Harness Agentic Workflow Design

**Date:** 2026-05-26  
**Status:** Approved (design phase)  
**Scope:** Lean execution + comprehensive quality system

## 1. Goal

Implement a harness workflow for smesec that preserves full lifecycle discipline (brainstorm → spec → plan → implement → verify) while staying lean for a small team. The workflow must include reasoning capture (`IMPLEMENT-NOTE.md`), a comprehensive rules system, and audit visibility.

## 2. Chosen Approach

Hybrid of previous options:
- Keep execution lean (single-file specs, no custom debate/verifier agents)
- Keep quality infrastructure comprehensive (rules system, implementation notes, audit digest/index)

### Why this approach
- Fits small team velocity (2–5 people)
- Avoids high maintenance of a fully custom agent ecosystem
- Preserves traceability and engineering memory
- Supports multi-platform consistency (web, mobile, desktop)

## 3. Target Architecture

```text
smesec/
├── AGENTS.md
├── CLAUDE.md (optional)
├── docs/
│   ├── ARCHITECTURE.md
│   ├── rules/
│   │   ├── 00-universal.md
│   │   ├── 01-architecture.md
│   │   ├── 02-security.md
│   │   ├── 03-web.md
│   │   ├── 04-mobile.md
│   │   ├── 05-desktop.md
│   │   └── 06-testing.md
│   ├── specs/
│   │   └── YYYY-MM-DD-feature-name/
│   │       ├── SPEC.md
│   │       ├── TASKS.md
│   │       └── IMPLEMENT-NOTE.md
│   ├── plans/
│   └── audit/
│       ├── implement-note-index.md
│       └── telemetry-digest.md
└── .claude/
    └── settings.local.json
```

## 4. Workflow Phases (Hard Gates)

```text
BRAINSTORM → SPEC → PLAN → IMPLEMENT (TDD) → VERIFY
```

### Phase 1 — Brainstorm
- Use `superpowers:brainstorming`
- Clarify requirements and constraints
- Compare 2–3 approaches with trade-offs
- Present design sections and get user approval

**Gate:** No code changes until design is approved.

### Phase 2 — Spec
- Write `docs/specs/YYYY-MM-DD-feature-name/SPEC.md`
- Single-file spec (lean structure)

Required sections in `SPEC.md`:
- Problem Statement
- Solution Overview
- Architecture / Data Flow
- Database Changes (if applicable)
- API Contract (if applicable)
- Platform-Specific Considerations (web/mobile/desktop)
- Testing Strategy
- Acceptance Criteria
- Out of Scope

**Gate:** Spec must be reviewed and approved.

### Phase 3 — Plan
- Use `superpowers:writing-plans`
- Write `TASKS.md` in same feature folder
- Initialize `IMPLEMENT-NOTE.md`

`TASKS.md` must contain strict TDD sequence per task:
1. Write failing test
2. Verify RED
3. Minimal implementation
4. Verify GREEN
5. Commit checkpoint (optional by team policy)

**Gate:** Task plan complete before implementation starts.

### Phase 4 — Implement (TDD)
- Use `superpowers:test-driven-development`
- Execute `TASKS.md` in order
- Update `IMPLEMENT-NOTE.md` during implementation

**Gate:** All required tests pass.

### Phase 5 — Verify
Inline checklist (no custom verifier agent):
- Acceptance criteria in `SPEC.md` satisfied
- Tests pass and coverage target met (>=85%)
- Relevant rules followed
- `IMPLEMENT-NOTE.md` completed
- Security checks pass
- Cross-platform consistency maintained

## 5. Comprehensive Rules System

Rules folder: `docs/rules/`

- `00-universal.md`: DRY, YAGNI, simplicity, safe coding defaults
- `01-architecture.md`: boundaries, data ownership, cross-platform contracts
- `02-security.md`: authn/authz, input validation, secrets, OWASP controls
- `03-web.md`: web stack conventions and performance patterns
- `04-mobile.md`: mobile conventions, platform constraints, offline/sync behavior
- `05-desktop.md`: desktop app boundaries, IPC/process/security patterns
- `06-testing.md`: test taxonomy, coverage targets, TDD discipline

Each rule file format:
- When
- What
- Why
- How to verify
- Correct example
- Incorrect example

## 6. IMPLEMENT-NOTE.md Standard

Location:
- `docs/specs/YYYY-MM-DD-feature-name/IMPLEMENT-NOTE.md`

Purpose:
- Record WHY decisions were made
- Preserve rejected alternatives
- Capture platform-specific gotchas
- Prevent repeated failed approaches

Required sections:
1. Pre-Implementation Context
2. Architecture Decisions
3. Implementation Challenges
4. Code Patterns Used
5. Testing Strategy
6. What Would Break This
7. Future Considerations

Lifecycle:
- Created at Plan phase (sections 1–2 filled)
- Updated during Implement phase (sections 3–7 filled)
- Required reading before modifying that feature

## 7. Audit Infrastructure

### 7.1 `docs/audit/implement-note-index.md`
- Index of all `IMPLEMENT-NOTE.md` files
- Grouped by feature, platform, and pattern
- Includes status/date/summary

### 7.2 `docs/audit/telemetry-digest.md`
- Rolling summary of recent shipped work
- Patterns emerging, common failures, rule updates, technical debt
- Updated every 3–5 features or monthly

## 8. AGENTS.md Standard

`AGENTS.md` acts as top-level navigation and workflow gatekeeper:
- Project identity and platform map
- Canonical pipeline with hard gate
- Rules index table by domain
- Quality gate checklist
- Common run/test/audit commands
- Onboarding flow for humans and agents

## 9. Non-Goals

This design intentionally does not include:
- Custom CPO/CTO debate agent system
- Custom verifier agent negotiation loop
- 8-file heavy spec folder model

These can be added later if team scale/complexity requires them.

## 10. Success Criteria

Harness is considered successfully implemented when:
1. Structure exists (`AGENTS.md`, rules, specs/plans/audit directories)
2. At least one feature goes end-to-end through the workflow
3. `IMPLEMENT-NOTE.md` is created and maintained during implementation
4. Audit index/digest files are updated
5. Team uses rules and gates consistently

## 11. Rollout Strategy

1. Bootstrap harness files and templates
2. Define initial rules set
3. Run first pilot feature through full pipeline
4. Tune rule wording and checklist thresholds from pilot learnings
5. Make workflow default for all non-trivial work

## 12. Risks and Mitigations

- **Risk:** Process gets skipped under pressure  
  **Mitigation:** Keep docs concise; enforce hard gate in AGENTS.md and planning routine.

- **Risk:** Rules become stale or overly generic  
  **Mitigation:** Update rules from real post-feature learnings in telemetry digest.

- **Risk:** IMPLEMENT-NOTE quality drifts  
  **Mitigation:** Add verification gate requiring complete sections before closure.

## 13. Decision Record

Approved by user on 2026-05-26:
- Hybrid model selected
- Keep lean execution workflow
- Include comprehensive rules system
- Include `IMPLEMENT-NOTE.md`
- Include full audit visibility
