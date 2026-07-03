package omnigent

import (
	"context"
	"strings"
	"testing"

	"github.com/togo-framework/autopilot"
	"github.com/togo-framework/togo"
)

// fakeEnv implements autopilot.Env for tests.
type fakeEnv struct {
	dir   string
	dirty bool
	omni  string
	runs  [][]string
}

func (f *fakeEnv) Dir() string { return f.dir }
func (f *fakeEnv) Run(_ context.Context, name string, args ...string) (string, error) {
	f.runs = append(f.runs, append([]string{name}, args...))
	if name == "git" {
		if f.dirty {
			return " M file.go", nil
		}
		return "", nil
	}
	return f.omni, nil // the omni CLI stdout
}

func (f *fakeEnv) omniArgs() []string {
	for _, r := range f.runs {
		if r[0] == "omni" {
			return r[1:]
		}
	}
	return nil
}

func has(ss []string, want string) bool {
	for _, s := range ss {
		if s == want {
			return true
		}
	}
	return false
}

func TestBuildPromptCarriesIssue(t *testing.T) {
	p := buildPrompt(autopilot.Issue{ID: "42", Title: "add /health", Body: "return 200"})
	if !strings.Contains(p, "add /health") || !strings.Contains(p, "return 200") || !strings.Contains(p, "BLOCKED:") {
		t.Fatalf("prompt missing content: %s", p)
	}
}

func TestImplementBlocked(t *testing.T) {
	env := &fakeEnv{dir: "/r", omni: "thinking…\nBLOCKED: which database?"}
	o := &omnigent{bin: "omni", flags: []string{"-p"}}
	res := o.Implement(context.Background(), env, autopilot.Issue{ID: "x", Title: "t"})
	if !res.NeedsInput || !strings.Contains(res.Reason, "which database?") {
		t.Fatalf("expected blocked with question, got %+v", res)
	}
	if a := env.omniArgs(); !has(a, "-p") {
		t.Fatalf("omni not invoked with -p: %v", a)
	}
}

func TestImplementChangedWithAgent(t *testing.T) {
	env := &fakeEnv{dir: "/r", omni: "done", dirty: true}
	o := &omnigent{bin: "omni", flags: []string{"-p"}, agent: "claude"}
	res := o.Implement(context.Background(), env, autopilot.Issue{ID: "x", Title: "t"})
	if res.NeedsInput || !res.Changed {
		t.Fatalf("expected changed, not blocked: %+v", res)
	}
	a := env.omniArgs()
	if !has(a, "--agent") || !has(a, "claude") {
		t.Fatalf("--agent claude missing: %v", a)
	}
}

func TestNewImplReadsConfig(t *testing.T) {
	t.Setenv("OMNIGENT_BIN", "omni2")
	t.Setenv("OMNIGENT_FLAGS", "run -p")
	t.Setenv("OMNIGENT_AGENT", "codex")
	o := newImpl(togo.New())
	if o.bin != "omni2" || len(o.flags) != 2 || o.agent != "codex" {
		t.Fatalf("config not read: %+v", o)
	}
}
