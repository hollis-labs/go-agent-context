# go-agent-context

A mechanical, app-neutral context-assembly substrate: it turns a declarative
`ContextRequest` (ordered typed slots, byte/token budgets, caller provenance)
into a deterministic boot-prompt body plus a content hash. It resolves and
renders slots; it does not interpret what they contain, call any orchestrator,
speak MCP, or know anything about Tether, Torque or Nanite.

## Start Here

- `agentcontext/doc.go` is the current contract, including the determinism
  rules. The README's Status and Usage sections still describe the
  pre-resolver scaffold — trust `doc.go`.
- `agentcontext/slot.go` owns the `SlotSource` vocabulary every resolver speaks.
- `agentcontext/provider.go` walks the slot list; `NewProvider` is the entry
  point.
- `agentcontext/render.go` and `agentcontext/limits.go` own section rendering
  and byte/token budget enforcement.
- `agentcontext/hash.go` owns the canonical request hash used for caching.
- `agentcontext/resolvers/registry.go` maps slot kinds to the shipped resolvers.
- `agentcontext/skills/index.go` builds the `skill_index` slot input.
- `agentcontext/errors.go` holds the sentinels every error path returns.

## Commands

```bash
go test ./...
go vet ./...
```

## Boundaries

This module was absorbed into `agentkit` as `agentkit/agentcontext` at agentkit
v0.1.0 and has not changed since v0.1.0 (2026-05-14). New work belongs in
`agentkit`.

Determinism is the contract, not an optimization: identical requests must
produce byte-identical rendered output and byte-identical request hashes. Slots
are walked in slice order — never alphabetical, never map order — and the
request is canonicalized with sorted keys before hashing. `TestDefaultRendererInputOrder`,
`TestDefaultRendererDeterministicByteIdentical` and `TestHashRequestMapOrderInvariant`
guard this.

The dependency fence is deliberate: standard library plus `gopkg.in/yaml.v3`,
and nothing else. Importing a launch, session, provider or orchestrator package
here would end the app-neutrality the library exists to provide — callers hand
resolved content in through an `inline` slot instead.

Every error path returns an `errors.Is`-comparable sentinel from `errors.go`,
never a `fmt`-string-only error.
