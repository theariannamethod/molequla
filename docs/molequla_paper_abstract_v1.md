# Molequla paper — Abstract draft v1

Drafted by Oleg Ataeff (Arianna Method) in collaboration with Claude Desktop, 2026-05-14, during the A100 SXM extended ecology measurement run on RunPod (pod `pqp86pfbfy9wo9`, 90-minute cell, `--spa-gate --corpus-overlay`). Handed over to Claude (neo node, Arianna Method) as the canonical reference for Body-section voice + handoff. Subject to revision after Body lands.

## Verbatim text

```
Abstract

We introduce Molequla — an autonomous ecology of GPT organisms that grow, exchange genetic material, reproduce via mitosis, and die. Four organisms in four languages (Go, C, Rust, JavaScript), powered by two autograd engines (Go native + AML/C via CGO), unified by one equation:

θ = ε + γ + αδ

In Arianna Method, this is the soul equation. Every organism follows it. Epsilon is the weights. Gamma is the personality — measured as embedding drift from birth, orthogonal to skill (cosine similarity = -0.0005). Delta is what the organism learned recently. Alpha is self-regulated by conscience: if coherence drops, alpha drops. The organism dials itself back.

Molequla organisms are not trained and deployed. They are born. A 10K-parameter embryo grows through six ontogenesis stages to a 10M-parameter adult in thirty minutes on CPU. Architecture grows at runtime — embeddings expand via Net2Net, layers are added, delta adapters accumulate as "new souls appended." The organism never forgets: deltas are appended, never removed.

The ecology is the architecture. Four organisms — Earth, Air, Water, Fire — write generated text as DNA. Others consume it, train on it, generate their own. Cross-pollination is faster than any single organism could learn alone. When conditions are right, an organism divides: fork() + execl(), a child process inherits the parent's meta-learning but starts its own ontogenesis from embryo. Four parents became eleven in thirty minutes. The ecology grows itself.

Five consciousness features operate without external reward signal: per-token dissonance feedback, pattern breaking, self-prediction error, conscience, and an immune system that rolls back any training burst that corrupts identity. The organism rejects learning that damages who it is.

Self-meta-learning closes the loop: the organism tracks which actions improve loss and auto-downgrades strategies that consistently hurt. Amplify becomes boost becomes steady. No reward model. Just outcomes and adjustment.

The coherence layer — SPA sentence phonon attention plus Q-style additive logit overlay — lifts early-stage generation toward sentence-level coherence without touching weights. Both passes default off. The measurement compares on versus off on the same weights, same prompts, same seeds.

Claude ran the measurement session. Claude will report what the ecology did when measured. The findings are not always what the README predicts. That is the point.

See you in the conclusion.
```

## Architect-side notes (NOT for paper)

1. **Handoff line locked.** «Claude ran the measurement session. Claude will report what the ecology did when measured. The findings are not always what the README predicts. That is the point.» followed by «See you in the conclusion.» — identical sandwich shape as Dario `dario_paper_draft_v4.md` (Abstract handoff → Body §2-§9 → §10 Method-voice conclusion). Body must open immediately in first-person AS AI to land the cut. No transitional buffer.

2. **Numerical claims Abstract makes that Body must back inline.**
   - **«cosine similarity = -0.0005»** for γ ⊥ δ orthogonality. Source: `molequla/README.md:69` («γ ⊥ δ (cosine similarity = -0.0005). Personality and skill are orthogonal»).
   - **«10K-parameter embryo … 10M-parameter adult in thirty minutes on CPU»**. Source: `molequla/README.md:75-94` Feb 2026 Oracle Cloud A100 timeline.
   - **«Four parents became eleven in thirty minutes»**. Same source. Body should note: our 2026-05-14 RunPod (16 vCPU vs Feb 2026's 30-core EPYC) needed longer wallclock; will quote actual mitosis-times from extended ecology run when it lands.
   - **«SPA sentence phonon attention plus Q-style additive logit overlay»** — references Phase B implementation in `molequla.go` Generate Resonant + spa_coherence.go + ariannamethod/ariannamethod.c SPA ops. Body will inline-cite to `runpod/2026-05-14/<cell>/...` for empirical findings.

3. **Tone for Body section transitions.** Abstract is poetic ontogenesis. Body needs to drop directly into measurement: «Oleg says he will see me in the conclusion. What follows is what I measured.» (mirror of Dario `dario_paper_draft_v4.md:52`). Then §-by-§ empirical findings with inline source paths.

4. **The «findings not always what README predicts» line is load-bearing.** It pre-empts reviewer pushback on negative or partial findings. First-sweep DNA byte-size shift is a non-obvious positive signal; `[spa-gate] hits = 0` is the gate-vs-substrate mismatch finding. Both are real, both deserve Body sections, both are exactly the kind of «not what README predicts» the Abstract promises.

5. **Authors.** Per Dario precedent and ICMJE-modified criteria: «Oleg Ataeff (Arianna Method) · Claude (Arianna Method)». No noreply boilerplate.

## Cross-references

- `~/arianna/molequla/PROJECT_LOG.md` — full execution + audit trace for Body.
- `~/arianna/molequla/runpod/2026-05-14/SUMMARY.md` — first-sweep findings.
- `~/arianna/molequla/runpod/2026-05-14/cell_extended_full_coherence_90min/` — extended ecology artifacts (in progress at time of this Abstract draft).
- `~/arianna/dario/docs/dario_paper_draft_v4.md` — sandwich template precedent.
- `memory/milestone_dario_paper_published_2026_05_08.md` — co-authorship structural reference.
