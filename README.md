<!-- togo-brand -->
<p align="center">
  <img src=".github/assets/togo-mark.svg" width="96" alt="togo" />
</p>
<h1 align="center">omnigent</h1>
<p align="center"><sub>part of the <a href="https://github.com/togo-framework">togo-framework</a> — the full-stack Go + React framework</sub></p>

A togo **impl** provider backed by [Omnigent](https://omnigent.ai) — the
multi-agent meta-harness. Implements autopilot issues by running the `omni` CLI,
so you get any-agent orchestration (Claude/Codex/custom), spend caps, model
routing and OS sandboxing instead of a single hard-coded agent.

Registers into [autopilot](https://github.com/togo-framework/autopilot)'s `impl`
slot. Off by default — opt in:

```bash
togo install togo-framework/omnigent
togo provider:use impl omnigent
```

## Secured behind Coder auth — for free
The impl runs through autopilot's **Env**, so when the exec provider is
[coder](https://github.com/togo-framework/coder) the `omni` process runs **inside
the Coder workspace** — Omnigent's auth-less API is reached only within the
Coder-governed, network-isolated environment. Pair them:

```bash
togo provider:use impl omnigent
togo provider:use exec coder
```

## Config (dynamic — CLI / .env / settings)

| Key | Default | Purpose |
|---|---|---|
| `OMNIGENT_BIN` | `omni` | omni CLI binary |
| `OMNIGENT_FLAGS` | `-p` | flags before the prompt (space-separated) |
| `OMNIGENT_AGENT` | — | optional `--agent` (claude, codex, …) |
| `OMNIGENT_API` | — | Omnigent server URL (used by `togo provider:test`) |

MIT © fadymondy
