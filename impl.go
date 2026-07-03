package omnigent

import (
	"context"
	"strings"
	"time"

	"github.com/togo-framework/autopilot"
	"github.com/togo-framework/providers"
	"github.com/togo-framework/togo"
)

// omnigent implements autopilot.Implementer by running the `omni` CLI in the Env.
type omnigent struct {
	bin   string
	flags []string
	agent string
}

func newImpl(k *togo.Kernel) *omnigent {
	flags := strings.Fields(providers.Value(k, providers.CapImplement, Name, "flags", "-p", false))
	return &omnigent{
		bin:   providers.Value(k, providers.CapImplement, Name, "bin", "omni", false),
		flags: flags,
		agent: providers.Value(k, providers.CapImplement, Name, "agent", "", false),
	}
}

// Implement runs Omnigent on the issue inside the Env. It detects an explicit
// "BLOCKED:" signal (needs a human decision) and whether any files changed.
func (o *omnigent) Implement(ctx context.Context, env autopilot.Env, issue autopilot.Issue) autopilot.ImplementResult {
	cctx, cancel := context.WithTimeout(ctx, 20*time.Minute)
	defer cancel()

	args := append([]string{}, o.flags...)
	if o.agent != "" {
		args = append(args, "--agent", o.agent)
	}
	args = append(args, buildPrompt(issue))

	stdout, runErr := env.Run(cctx, o.bin, args...)
	changed := envDirty(cctx, env)

	if strings.Contains(stdout, "BLOCKED:") {
		reason := stdout[strings.Index(stdout, "BLOCKED:")+len("BLOCKED:"):]
		return autopilot.ImplementResult{NeedsInput: true, Reason: trim(strings.TrimSpace(reason), 1500), Changed: changed}
	}
	if runErr != nil && !changed {
		return autopilot.ImplementResult{Reason: "omnigent: " + trim(runErr.Error(), 600), Changed: false}
	}
	return autopilot.ImplementResult{Summary: trim(strings.TrimSpace(stdout), 2000), Changed: changed}
}

// envDirty reports whether the Env's working tree has uncommitted changes.
func envDirty(ctx context.Context, env autopilot.Env) bool {
	out, err := env.Run(ctx, "git", "status", "--porcelain")
	return err == nil && strings.TrimSpace(out) != ""
}

func buildPrompt(issue autopilot.Issue) string {
	return "You are an autonomous engineering agent working in this repository. " +
		"Implement the following issue end to end by editing files in the working tree. Keep changes minimal and focused.\n\n" +
		"Issue " + issue.ID + " — " + issue.Title + "\n" + issue.Body + "\n\n" +
		"Rules:\n" +
		"- Make the code changes directly; do NOT run git commit/push (the runner handles git).\n" +
		"- If you cannot proceed without a human decision, do NOT guess: make no changes and reply with one line starting 'BLOCKED: <your question>'.\n" +
		"- When done, briefly summarize what you changed."
}

func trim(s string, n int) string {
	if len(s) > n {
		return s[:n] + "…"
	}
	return s
}
