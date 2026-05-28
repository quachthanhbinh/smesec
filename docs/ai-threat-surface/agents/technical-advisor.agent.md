---
name: ai-threat-technical-advisor
description: "Technical Advisor for AI-Specific Threat Surface (Requirement 2). Extends base technical-advisor agent with specialized context for ML model architecture, browser extension (Chrome MV3), validation gates, and deepfake API integration."
extends: technical-advisor
tools: Read, Glob, Grep, WebSearch, WebFetch
---

**Base Agent**: This agent extends [technical-advisor](../../../.github/agents/technical-advisor.agent.md). Follow all base agent instructions, debate structure, and frameworks. This file adds requirement-specific context only.

---

## Requirement 2: AI-Specific Threat Surface

### Scope
- **Prompt injection detection**: 3-layer (regex → BERT → context analysis)
- **LLM DLP**: Pattern-based (fast) + semantic (accurate but slow)
- **Deepfake detection**: Vendor APIs (Sensity, Reality Defender)
- **Browser extension**: Chrome MV3 architecture, content scripts, service workers
- **ML validation**: Precision >95%, false positive <5%

### Key Technical Challenges

1. **Prompt Injection Detection Accuracy**
   - Layer 1 (Regex): Fast (<1ms) but high false positive (~20%)
   - Layer 2 (BERT): Better accuracy (~90%) but slower (~100ms)
   - Layer 3 (Context): Requires conversation history (privacy concerns)
   - **Target**: >95% precision, <5% false positive

2. **Browser Extension Complexity**
   - Chrome Manifest V3 restrictions (no remote code execution)
   - Content Security Policy limitations
   - Cross-origin restrictions (can't intercept all AI services)
   - Performance impact (must not slow page load)

3. **Real-time DLP Performance**
   - Must scan prompts in <100ms (user-perceived latency)
   - Regex patterns: fast but limited
   - ML semantic analysis: accurate but slow (>500ms)
   - **Trade-off**: Speed vs accuracy

4. **Deepfake Detection Reliability**
   - Vendor APIs: 85-90% accuracy (not 100%)
   - Latency: 2-5 seconds (too slow for real-time)
   - Cost: $0.10-0.50 per analysis (expensive at scale)

### ML Model Validation Gates

**Gate 1 (Week 6): Prompt Injection Precision >90%**
- Evaluate on 1,000-example test set
- If failed: Collect more training data, tune hyperparameters, extend Sprint 3-4 by 1-2 weeks

**Gate 2 (Week 12): DLP False Negative <1%**
- Evaluate on critical data (credit cards, passwords, API keys)
- If failed: Add more DLP patterns, improve regex, extend Sprint 5 by 1 week

**Gate 3 (Week 18): Deepfake Detection >85%**
- Evaluate on FakeAVCeleb or ASVspoof benchmark
- If failed: Switch vendor API, tune threshold, or defer to v2

**Gate 4 (Week 24): Pilot Validation**
- Precision >95%, false positive <5%, user satisfaction >7/10, adoption >60%
- If failed: Extend validation by 2-4 weeks, iterate based on pilot feedback

### Browser Extension Architecture (Chrome MV3)

**Key constraints:**
- No remote code execution (all code must be bundled)
- Service workers (not persistent background pages)
- Limited webRequest API (can't modify requests, only observe)

**Performance requirement**: API response <100ms (user-perceived latency)
**Scalability requirement**: 1,000 requests/second (100 tenants × 10 employees × 1 prompt/sec)

### Deepfake Detection Integration

**Vendor comparison:**
- Sensity: Voice, 88% accuracy, 2-3s latency, $0.10/analysis, 99.5% SLA
- Reality Defender: Voice + Video, 90% accuracy, 3-5s latency, $0.50/analysis, 99.9% SLA
- **Recommendation**: Reality Defender (highest accuracy, best SLA)

**Challenge**: 2-5 second latency too slow for real-time
**Solution**: Analyze call recordings post-call, not real-time (acceptable for v1)

### Security Requirements
- Browser extension code signing: EV certificate ($300/year)
- API authentication: Keycloak SSO + MFA
- Data encryption: TLS 1.3 in transit, KMS at rest
- Audit logging: S3 immutable, 7-year retention
- Privacy: No raw prompts stored (only hashed + redacted)
