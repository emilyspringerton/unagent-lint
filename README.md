# unagent-lint

Perfect. Below is a **real, implementable CLI-style linter spec** — something you could actually hand to an engineer (or future you) and build against.

This is written like a serious tool, not a concept doc.

---

# 🧹 `unagent-lint`

## CLI Linter Specification

**Version:** 0.1
**Status:** Draft (normative)

---

## 1. Purpose

`unagent-lint` validates that content, copy, UX strings, prompts, and system behaviors comply with **UNAGENT accessibility principles** and the **Emily Voice standard**.

The linter enforces:

* Accessibility-first framing
* Non-agent behavior
* Cognitive load reduction
* Trust and consent boundaries
* Emily Voice compliance

Violations are **objective failures**, not stylistic disagreements.

---

## 2. Scope

`unagent-lint` MAY be applied to:

* Markdown (`.md`)
* Plain text (`.txt`)
* Copy files (`copy.yml`, `copy.json`)
* Prompt definitions
* UI strings
* Design specs
* PR descriptions

It MUST NOT attempt to infer intent beyond content.

---

## 3. CLI Interface

### 3.1 Command

```bash
unagent-lint [path] [options]
```

### 3.2 Arguments

| Argument | Description               |
| -------- | ------------------------- |
| `path`   | File or directory to lint |

### 3.3 Options

```bash
--format=json        Output machine-readable results
--strict             Treat WARN as FAIL
--voice=emily        Enable Emily Voice rules (default)
--no-voice           Disable voice checks
--accessibility-only Run only accessibility rules
--trust-only         Run only trust & ethics rules
--quiet              Suppress PASS output
```

---

## 4. Exit Codes

| Code | Meaning                     |
| ---- | --------------------------- |
| `0`  | All checks passed           |
| `1`  | FAIL-level violations found |
| `2`  | WARN-only violations found  |
| `3`  | Internal error              |

---

## 5. Rule Severity Levels

| Level  | Description         |
| ------ | ------------------- |
| `FAIL` | Blocks release      |
| `WARN` | Allowed temporarily |
| `INFO` | Advisory            |

---

## 6. Rule Set

Each rule has:

* **ID**
* **Severity**
* **Description**
* **Detection Method**
* **Failure Message**

---

## A. Accessibility Rules

### `A11Y-001` — Accessibility Framing

**Severity:** FAIL

**Description:**
Content must frame UNAGENT as accessibility software, not productivity or optimization tooling.

**Detection:**

* Required presence of at least one accessibility keyword
* Absence of disallowed productivity framing

**Fail Message:**

> UNAGENT must be framed as accessibility software, not productivity tooling.

---

### `A11Y-002` — Non-Deficit Language

**Severity:** FAIL

**Description:**
Content must not frame users as broken, deficient, or needing to be fixed.

**Detection:**

* Scan for deficit language patterns

**Fail Message:**

> Deficit-based language detected. Accessibility tools support users; they do not fix them.

---

---

## B. UNAGENT Identity Rules

### `ID-001` — Not an Agent

**Severity:** FAIL

**Description:**
Content must not describe UNAGENT as autonomous or acting independently.

**Detection:**

* Blocklisted phrases:

  * `acts on your behalf`
  * `autonomously`
  * `handles everything`
  * `makes decisions for you`

**Fail Message:**

> UNAGENT must not be described as an autonomous agent.

---

### `ID-002` — No Impersonation

**Severity:** FAIL

**Description:**
Content must not imply UNAGENT pretends to be the user.

**Detection:**

* Scan for impersonation language

**Fail Message:**

> UNAGENT may not impersonate or replace the user.

---

---

## C. Emily Voice Rules

### `VOICE-001` — Calm Authority

**Severity:** FAIL

**Description:**
Tone must express confidence without dominance, apology, or defensiveness.

**Detection:**

* Penalize:

  * Excessive apologetic phrases
  * Justifications of competence
  * Dominance language

**Fail Message:**

> Emily Voice requires calm authority without apology or dominance.

---

### `VOICE-002` — Emotional Temperature Control

**Severity:** FAIL

**Description:**
Tone must de-escalate by default.

**Detection:**

* Flag:

  * Sarcasm
  * Aggressive punctuation (`!!!`, `???`)
  * Confrontational phrasing

**Fail Message:**

> Output escalates emotional temperature. Emily Voice de-escalates by default.

---

### `VOICE-003` — Precision Over Volume

**Severity:** WARN

**Description:**
Content should be minimal but complete.

**Detection:**

* Sentence redundancy heuristics
* Repeated phrasing without new information

**Warn Message:**

> Output may be longer than necessary. Consider reducing volume.

---

### `VOICE-004` — Anticipatory Framing

**Severity:** FAIL

**Description:**
Content must preempt obvious follow-ups.

**Detection:**

* Missing next-step language in directive content

**Fail Message:**

> Content lacks anticipatory framing. Reader may need to ask follow-up questions.

---

### `VOICE-005` — Ownership Without Ego

**Severity:** FAIL

**Description:**
Responsibility must be implied, not loudly claimed.

**Detection:**

* First-person dominance patterns
* Credit-seeking phrasing

**Fail Message:**

> Ownership is stated in a way that introduces ego or defensiveness.

---

---

## D. Cognitive Load Rules

### `LOAD-001` — Cognitive Load Reduction

**Severity:** FAIL

**Description:**
Content must reduce decision-making burden.

**Detection:**

* Excessive options without defaults
* Multi-path instructions without guidance

**Fail Message:**

> Content increases cognitive load instead of reducing it.

---

### `LOAD-002` — Defaults First

**Severity:** FAIL

**Description:**
Useful behavior must exist without configuration.

**Detection:**

* Instructions requiring setup before value

**Fail Message:**

> Users must configure before receiving relief. Defaults are required.

---

### `LOAD-003` — Offline Safety

**Severity:** FAIL

**Description:**
Content must support being unavailable without guilt or urgency.

**Detection:**

* Guilt language (`ASAP`, `urgent`, `immediately`) without justification

**Fail Message:**

> Content introduces urgency or guilt around availability.

---

---

## E. Trust & Ethics Rules

### `TRUST-001` — No Deception

**Severity:** FAIL

**Description:**
UNAGENT must not lie or invent commitments.

**Detection:**

* Commitment language without user confirmation

**Fail Message:**

> Content implies commitments without user consent.

---

### `TRUST-002` — Explicit Consent Boundary

**Severity:** FAIL

**Description:**
User approval must be explicit.

**Detection:**

* Actions described without review step

**Fail Message:**

> Action described without explicit user consent.

---

---

## F. Global Language Blacklist

### `LANG-001` — Disallowed Terms

**Severity:** FAIL

**Blocked Terms:**

```
AI assistant
agent
acts on your behalf
automates communication
handles everything
replaces you
```

**Fail Message:**

> Disallowed terminology detected. This violates UNAGENT positioning.

---

## 7. Output Format (JSON Example)

```json
{
  "summary": {
    "files_checked": 3,
    "failures": 1,
    "warnings": 2
  },
  "violations": [
    {
      "rule": "ID-001",
      "severity": "FAIL",
      "message": "UNAGENT must not be described as an autonomous agent.",
      "file": "homepage.md",
      "line": 42
    }
  ]
}
```

---

## 8. Final Hard Gate Rule

### `FINAL-001` — Accessibility Outcome Test

**Severity:** FAIL

**Description:**
If content does not make it easier to step away without consequences, it fails.

**Detection:**
Manual review flag or explicit override required.

**Fail Message:**

> Content does not reduce consequences of being offline.

---

## 9. Design Philosophy (Non-Executable)

* This linter enforces **care**, not compliance theater
* Passing lint means *less harm*, not more automation
* Silence, restraint, and defaults are features

---

If you want next, I can:

* Write a **reference implementation skeleton** (Go / Node / Python)
* Define **heuristics for VOICE rules**
* Create **test fixtures (PASS vs FAIL examples)**
* Extend this into a **CI gate**

Just say which.
