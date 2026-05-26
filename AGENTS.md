# AGENTS.md - Workflow Gatekeeper

## Project Identity

| Platform | Status |
|----------|--------|
| Web | ✓ |
| Mobile | ✓ |
| Desktop | ✓ |

## Workflow

All feature development follows this pipeline:

**BRAINSTORM → SPEC → PLAN → IMPLEMENT (TDD) → VERIFY**

### Hard Gate

No implementation without an approved spec and plan. No exceptions.

## Rules Index

All agents must follow these rules in order:

1. `docs/rules/00-universal.md` - Universal development principles
2. `docs/rules/01-architecture.md` - Architecture patterns and decisions
3. `docs/rules/02-security.md` - Security requirements and practices
4. `docs/rules/03-web.md` - Web platform specifics
5. `docs/rules/04-mobile.md` - Mobile platform specifics
6. `docs/rules/05-desktop.md` - Desktop platform specifics
7. `docs/rules/06-testing.md` - Testing standards and practices

## Quality Gates

Before marking any task complete, verify:

- [ ] All tests pass
- [ ] Code follows architecture patterns from rules
- [ ] Security checklist completed (if applicable)
- [ ] Platform-specific guidelines followed
- [ ] IMPLEMENT-NOTE.md created and indexed
- [ ] Changes verified in running application
