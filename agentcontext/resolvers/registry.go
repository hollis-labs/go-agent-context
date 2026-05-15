package resolvers

import (
	"github.com/hollis-labs/go-agent-context/agentcontext"
)

// Default returns a map populated with the seven app-neutral
// resolvers shipped in this package:
//
//   - static_file  (StaticFileResolver)
//   - static_dir   (StaticDirResolver)
//   - inline       (InlineResolver)
//   - cmd          (CmdResolver)
//   - http_text    (HTTPTextResolver)
//   - http_json    (HTTPJSONResolver)
//   - role_summary (RoleSummaryResolver)
//
// The skill_index resolver is intentionally NOT included — that kind
// is owned by the skill-discovery layer (Subagent C / sibling
// subpackage), and the on-disk skill model is layered on top of the
// agentcontext contract. Callers that need skill_index should compose
// their own resolver map, e.g.:
//
//	res := resolvers.Default()
//	res[agentcontext.SlotSourceKindSkillIndex] = skill.NewIndexResolver(...)
//
// Each entry is constructed with its zero-config defaults. Callers
// who need to tune individual resolvers should build their own map
// using the per-resolver constructors (NewStaticDirResolver, etc.)
// and the functional options each exposes.
func Default() map[agentcontext.SlotSourceKind]agentcontext.Resolver {
	return map[agentcontext.SlotSourceKind]agentcontext.Resolver{
		agentcontext.SlotSourceKindStaticFile:  NewStaticFileResolver(),
		agentcontext.SlotSourceKindStaticDir:   NewStaticDirResolver(),
		agentcontext.SlotSourceKindInline:      NewInlineResolver(),
		agentcontext.SlotSourceKindCmd:         NewCmdResolver(),
		agentcontext.SlotSourceKindHTTPText:    NewHTTPTextResolver(),
		agentcontext.SlotSourceKindHTTPJSON:    NewHTTPJSONResolver(),
		agentcontext.SlotSourceKindRoleSummary: NewRoleSummaryResolver(),
	}
}
