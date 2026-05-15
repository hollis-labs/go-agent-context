# go-agent-context

`go-agent-context` is a small, standalone Go library that provides a portfolio-shared **mechanical context assembly substrate** for agent runtimes — the layer that takes a declarative `ContextRequest` (an ordered list of typed slot specifications, byte/token budgets, and caller-supplied provenance), resolves each slot through a registered `Resolver` (static files, inline strings, shell commands, HTTP fetches, role summaries, skill indexes), and renders the resolved slots into a deterministic boot-prompt body suitable for injection into a [`go-agent-launch`](https://github.com/hollis-labs/go-agent-launch) `PreparedLaunch`.

It sits **above** [`go-agent-launch`](https://github.com/hollis-labs/go-agent-launch) — the launch substrate that materializes the bootdir and prepares the spawn `argv` — and **below** the app-specific orchestrators (Tether, Torque, Nanite) that each previously grew their own near-identical context-assembly pipelines. The library defines the shared `ContextRequest` → `ContextResult` flow, the typed `SlotSource` vocabulary, the `Resolver` and `Renderer` extension points, the deterministic content hash, and the provenance + budget enforcement contract so a single context spec can drive every consumer.

## Status

v0 — foundation. Public API in flux; expect breaking changes before the v0.1.0 tag. This repository currently ships the package scaffold and the public context-assembly contracts only; the concrete slot resolvers, skill discovery, and `go-agent-launch` integration land in subsequent Phase 2 subagents under epic `EP-20260514-0001` / sprint `SP-20260514-0004`.

## Scope

**In scope.** Mechanical, deterministic context assembly:

- Typed slot vocabulary (`static_file`, `static_dir`, `inline`, `cmd`, `http_text`, `http_json`, `role_summary`, `skill_index`) — the `SlotSource` tagged union.
- The `Resolver` interface that pluggable backends implement for each kind.
- A default section-headered `Renderer` that composes slot results in input order (deterministic, NOT alphabetical).
- Byte and token budget enforcement (`Limits` / `LimitsApplied`) — token estimate uses the documented char/4 heuristic; precise tokenization is out of scope.
- Caller-supplied provenance pass-through (`ProvenanceInput`) plus per-slot resolver-emitted attribution (`SlotProvenance`).
- A stable content hash over the canonical request (`HashRequest`) for caching.
- Validation sentinels — every error path is an `errors.Is`-comparable sentinel.

**Out of scope.** Anything that interprets the assembled context or reaches into an app's runtime:

- No Nanite-specific chat, classifier, reminder, or UI logic.
- No Vanta `memory_recall` / `conduit_lookup` calls. (Callers can hand resolved Vanta output into an `inline` slot, but the library does not call Vanta.)
- No MCP transport, no Anthropic SDK, no provider adapters.
- No app business logic, no project-specific conventions, no opinionated defaults beyond "section headers + input order".

The library is intentionally app-neutral. It imports only the Go standard library and `gopkg.in/yaml.v3` (for the boot-profile YAML shape next subagents will consume). It does not import `go-agent-launch`, `go-agent-sessions`, `go-providers`, Tether, Torque, or Nanite. Consumers configure the pipeline through caller-supplied resolvers, renderers, and provenance rather than direct dependencies on any orchestrator.

## Install

```bash
go get github.com/hollis-labs/go-agent-context
```

## Usage

```go
// TODO(phase-2): runnable example after Subagent B lands the concrete
// file/inline/cmd resolvers and Subagent D wires Assemble into a
// go-agent-launch ContextHook. Caller pattern (placeholder API):
//
//     provider := agentcontext.NewProvider(resolvers, renderer)
//     result, err := provider.Assemble(ctx, agentcontext.ContextRequest{
//         Slots:   []agentcontext.SlotSpec{ /* role_summary, cmd, etc. */ },
//         Limits:  agentcontext.Limits{MaxBytes: 32_000},
//         Workdir: workdir,
//         Provenance: agentcontext.ProvenanceInput{LineageAlias: "nanite.backend.main"},
//     })
//     if err != nil { /* ... */ }
//     // result.Rendered is the deterministic boot-prompt body.
```

A runnable end-to-end example will live under [`examples/`](./examples) once
the concrete resolvers and the `go-agent-launch` `ContextHook` wrapper are
committed.

## License

MIT — see [LICENSE](./LICENSE).
