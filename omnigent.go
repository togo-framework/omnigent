// Package omnigent is a togo impl provider backed by Omnigent
// (https://omnigent.ai) — the multi-agent meta-harness. It implements autopilot
// issues by running the `omni` CLI, so you get any-agent orchestration, spend
// caps, model routing and OS sandboxing instead of a single hard-coded agent.
//
// It registers into autopilot's `impl` capability slot. Opt in with
// `togo provider:use impl omnigent` (or TOGO_IMPL_PROVIDER=omnigent); it is off
// by default so `claude` stays the default.
//
// Because the impl runs THROUGH autopilot's Env, when the exec provider is
// togo-framework/coder the `omni` process runs INSIDE the Coder workspace — so
// Omnigent's API (which has no auth of its own) is reached only within the
// Coder-governed, network-isolated environment. That is "Omnigent behind Coder
// auth", achieved for free by the Env.
//
// Fully dynamic over CLI / .env / settings:
//
//	OMNIGENT_BIN     omni CLI binary        (default "omni")
//	OMNIGENT_FLAGS   flags before the prompt (default "-p"; space-separated)
//	OMNIGENT_AGENT   optional --agent value  (e.g. claude, codex)
//	OMNIGENT_API     Omnigent server URL     (used by `togo provider:test`)
package omnigent

import (
	"github.com/togo-framework/providers"
	"github.com/togo-framework/togo"
)

// Name is the plugin's stable identifier and its `impl` backend name.
const Name = "omnigent"

func init() {
	togo.RegisterProviderFunc(Name, togo.PriorityService, func(k *togo.Kernel) error {
		providers.Use(k, providers.CapImplement, Name, newImpl(k), false)
		if k.Log != nil {
			k.Log.Info("plugin active", "plugin", Name)
		}
		return nil
	})
}
