# Project TODOs and Decisions

This file captures answers from the OPEN_QUESTIONS.md quizzing process. Each entry references the corresponding question number (Q#) and records decisions or follow-ups.

## Vision & Scope
- Q1 (Target genre and fidelity): Early Game Boy-era Pokémon feel; classic top-down adventure loop with that aesthetic and pacing.
- Q2 (Single vs. multi-session): Persist worlds across sessions; store generated elements/layouts to avoid regeneration cost.
- Q3 (Online requirements): Cloud-hosted LLM; game world stored locally. Multiplayer considered in later development.

## Agent Interfaces (JSON I/O)
- Q4 (Canonical schemas): Publish JSON Schema files; schemas need to be created.
- Q5 (Extensibility): Use MCP (Model Context Protocol) for capability negotiation/extension.
- Q6 (Validation strategy): Engine performs strict schema validation; agents negotiate schema via MCP. Define behavior for partial/invalid responses.

## Turn System & Orchestration
- Q7 (Authoritative clock): Needs further investigation. TODO: Propose options and trade-offs for turn advancement and concurrency.
- Q8 (Event queue): Each agent decides whether to emit an EVENT for its latest generation.
- Q9 (Conflict resolution): Object ownership model—creator agent owns the object and has final say in negotiations.

## Game Engine Integration
- Q10 (Target engines): Use Godot as the first-class engine.
- Q11 (Transport): Prefer MCP, ideally over gRPC. TODO: Define backpressure and timeout policies.
- Q12 (State authority): Agents may mutate state; engine validates all changes.

## Character Agent Behavior
- Q13 (Action set): Not final; agents may introduce new actions on the spot. TODO: Define a concept/protocol for dynamic action creation and negotiation.
- Q14 (Memory store): Store MEMORIES in a per-agent vector DB; supervisor may access all memories.
- Q15 (Needs model): A dedicated agent manages global NEEDS values and decay rates, relating needs across agents.
- Q16 (Safety rails): Engine validation plus a supervisor agent govern rule-breaking. Repeat offenders are deleted (including memories) and replaced by new agents with different starting behaviors.

## Object FSM Generation
- Q17 (Naming consistency): Yes. Enforce standardized naming via engine (e.g., UPPER_SNAKE states, verb-phrase transitions).
- Q18 (Requirements field): Deferred.
- Q19 (FSM limits/sanity): Engine must sanity-check graphs to prevent contradictions/degeneracy. TODO: Define concrete checks and thresholds.
- Q20 (Versioning/migration): Allow adding new states reachable from existing states. Deleting old states only if not active in any save; deleting transitions only if the resulting graph passes sanity checks.

## Interaction Mapping Tool
- Q21 (Time estimates): Skipped.
- Q22 (Cross-object planning): Skipped.
- Q23 (Error policy): Skipped.

## Art/Asset Generation
- Q24 (Style lock): Match early Game Boy Pokémon look and feel (isometric/pixel-art aesthetic as per repo vision).
- Q25 (Variant creation): Dedicated art-style agent manages object states and visual variants.
- Q26 (Caching & rights): Assets stored locally; approximate reproducibility via granular prompts; seed/model reproducibility not guaranteed.

## Multiplayer (PwPwA)
- Q27 (Authority model): Skipped.
- Q28 (Sync): Skipped.
- Q29 (Cheat prevention): Skipped.
