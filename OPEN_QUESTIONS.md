# Open Questions

These questions arise from README.md and agent_concepts/* and should be clarified before further implementation.

**Vision & Scope**
- **Target genre and fidelity**: What concrete game loop and level of simulation depth are we aiming for (life-sim, roguelike, narrative adventure)?
- **Single vs. multi-session**: Should worlds persist across sessions or reset per scenario/demo?
- **Online requirements**: Must the project run fully offline, or is cloud LLM/art generation acceptable?

**Agent Interfaces (JSON I/O)**
- **Canonical schemas**: What are the exact JSON schemas for inputs/outputs (versions, required/optional fields, enums)? Will we publish JSON Schematron/JSON Schema files?
- **Extensibility**: How do agents negotiate capabilities (e.g., additional actions for characters, extra metadata for objects)?
- **Validation strategy**: Where is schema validation enforced (engine, middleware, or at agent boundary)? What happens on partial/invalid responses?

**Turn System & Orchestration**
- **Authoritative clock**: Who advances turns (engine vs. coordinator agent)? How are concurrent agents sequenced deterministically?
- **Event queue**: How are `EVENT`s produced, ordered, and consumed across multiple agents in a turn?
- **Conflict resolution**: If two agents act on the same object, who arbitrates state transitions?

**Game Engine Integration**
- **Target engines**: Which engines are first-class (Unity, Godot, web)? Any reference adapter to prove JSON protocol end-to-end?
- **Transport**: How do engine and agents communicate (stdin/stdout, websockets, HTTP, gRPC)? Any backpressure/timeouts?
- **State authority**: Is the engine the source of truth for object state, or can agents mutate state directly?

**Character Agent Behavior**
- **Action set**: Are the six actions final? Do we foresee inventory, craft, pathfind, emote variants, or meta-actions (plan, wait)?
- **Memory store**: Where are `MEMORIES` and `RELATIONS` persisted and updated (engine, vector DB, agent-internal)? What size/eviction policy?
- **Needs model**: What schema and dynamics govern `NEEDS` (decay rates, death thresholds, recovery mechanics)?
- **Safety rails**: How do we prevent self-harm or rule-breaking actions beyond the allowed set?

**Object FSM Generation**
- **Naming consistency**: Are state and transition naming conventions standardized (e.g., UPPER_SNAKE for states, verb phrases for transitions)?
- **Requirements field**: How are `T_X_REQUIREMENT` items represented, validated, and acquired (objects by ID vs. by type)?
- **FSM limits**: Max states/transitions per object? How to prevent degenerate or contradictory graphs?
- **Versioning**: Can an object’s FSM evolve after creation without breaking saved games?

**Interaction Mapping Tool**
- **Time estimates**: How are `T_n_TIME` values calibrated and used (animation length vs. simulated time)? Units standardization?
- **Search/plan**: If intent requires multi-step across objects, can the mapper chain objects or only one at a time?
- **Error policy**: What qualifies as “bending the rules too much”? Do we suggest alternative intents or nearest valid transitions?

**Art/Asset Generation**
- **Style lock**: Are we fixed to isometric pixel art? Required canvas size, palette, and animation frames per state?
- **Variant creation**: Who generates additional states/angles/lighting for an object after the first image?
- **Caching & rights**: How are assets cached and licensed (DALL·E or local)? Reproducibility across seeds/models?

**Multiplayer (PwPwA)**
- **Authority model**: Who is authoritative in multiplayer (server, client, or agent)?
- **Sync**: How are turns synchronized across human players and agents with different latencies?
- **Cheat prevention**: Any anti-cheat or validation on client actions vs. agent/world rules?

**Persistence & IDs**
- **Identifiers**: How are objects/agents addressed (stable IDs, namespacing, lifecycle)?
- **Save format**: What is the canonical world save format (JSON bundle, DB)? Migration strategy?
- **Determinism**: How to ensure replayability given stochastic LLM outputs?

**Error Handling & Safety**
- **Sandboxing**: What constraints are enforced on agent outputs to avoid unsafe or offensive content?
- **Fallbacks**: Do we have automatic retries, self-healing prompts, or guard agents on malformed JSON?
- **Observability**: What logs/telemetry are captured to debug bad plans or mapping failures?

**Performance & Cost**
- **Model choices**: Which LLMs/vision models are targeted (GPT-4 class vs. open models)? Swappable backends?
- **Budgeting**: Any token/image generation budgets per turn/session? Throttling policies?
- **Caching**: Will we cache prompts/responses and FSM generations to cut cost/latency?

**Testing & Evaluation**
- **Golden paths**: What scenarios define correctness for mapping and character actions?
- **Sim fuzzing**: Will we fuzz intents and states to stress FSM validity and error handling?
- **Human-in-the-loop**: Any review tooling to approve generated FSMs/assets before use?

**Licensing & Governance**
- **Content policy**: Are user- and model-generated assets licensed under the repo license? Any attribution requirements?
- **Trademark/IP**: How do we prevent requests for protected content in art generation?
- **Contribution model**: Who approves new actions/schemas affecting compatibility?

**Roadmap & Deliverables**
- **MVP slice**: What is the smallest end-to-end demo we target (engine, one room, 3 objects, 1 character)?
- **Milestones**: In what order do we implement: schemas → engine adapter → mapping tool → FSM generator → character agent → assets?
- **Docs**: Will we include machine-readable schemas and a reference transcript of a few turns?

