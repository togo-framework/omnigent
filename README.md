<!-- togo-header -->
<div align="center">
  <picture><source media="(prefers-color-scheme: dark)" srcset=".github/assets/togo-mark-dark.svg" /><img src=".github/assets/togo-mark.svg" alt="ToGO" height="64" /></picture>
  <h1>togo-framework/omnigent</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1F8A99" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/omnigent"><img src="https://pkg.go.dev/badge/github.com/togo-framework/omnigent.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/omnigent
```

<!-- /togo-header -->

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
<!-- togo-sponsors -->
---

<div align="center">
  <h3>💎 Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><img src=".github/assets/id8media.svg" height="44" alt="ID8 Media" /></a>
    &nbsp;&nbsp;&nbsp;&nbsp;
    <a href="https://one-studio.co"><img src=".github/assets/one-studio.jpeg" height="44" alt="One Studio" /></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
