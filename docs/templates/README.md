# Documentation Templates

This directory contains templates for consistent project documentation.

## Available Templates

### Decision Record Template

**File:** [decision-record-template.md](decision-record-template.md)

**Purpose:** Comprehensive record of architectural decisions, trade-offs, and rationale for major projects.

**When to use:**
- Starting a new major feature or system
- Evaluating multiple architectural approaches
- Need to document decision-making process for stakeholders
- Want to capture context for future reference

**How to use:**
1. Copy the template to `docs/superpowers/specs/YYYY-MM-DD-[project-name]-decision-record.md`
2. Fill in all sections during brainstorming and design phase
3. Update as decisions evolve
4. Mark as "Approved" when finalized
5. Reference from implementation plans

**Key sections:**
- **Project Context:** Requirements, constraints, clarifying Q&A
- **Options Analysis:** 2-3 architectural approaches with pros/cons/economics
- **Decision Rationale:** Why the chosen option, trade-offs accepted
- **Tech Stack:** Technology choices with alternatives considered
- **Team & Delivery:** Structure, timeline, milestones, risk validation
- **Cost Model:** Pricing strategy and unit economics
- **Scope Boundaries:** What's in/out of current plan
- **Success Criteria:** Technical, business, and team metrics
- **Risks:** Open questions and mitigation strategies

**Best practices:**
- Fill in during brainstorming, not after implementation
- Capture actual Q&A dialogue, not sanitized version
- Include economics for each option (build cost, margins)
- Document WHY decisions were made, not just WHAT
- Update when assumptions change
- Link to related specs and plans

## Template Maintenance

When creating new templates:
1. Add template file to this directory
2. Update this README with description and usage
3. Follow naming convention: `[type]-template.md`
4. Include inline comments/instructions in template
5. Provide example if helpful

## Related Documentation

- **Specs:** `docs/superpowers/specs/` - Design specifications
- **Plans:** `docs/superpowers/plans/` - Implementation plans
- **Pilot:** `docs/pilot/` - Pilot-specific documentation
