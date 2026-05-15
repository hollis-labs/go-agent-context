package agentcontext

import "errors"

// Sentinel errors returned by Validate methods, the default
// DefaultProvider.Assemble dispatcher, and the default Renderer. All
// sentinels are errors.Is-comparable; callers branch on them with
// errors.Is rather than string match.
//
// The errors document which field or slot failed; they do not carry
// the offending value itself. Validation that needs to report a
// specific value wraps the sentinel with
// fmt.Errorf("%w: %s", sentinel, value).
var (
	// ErrMissingSlotName is returned by SlotSpec.Validate and by
	// ContextRequest.Validate when a slot has an empty Name. Slot names
	// are load-bearing for determinism (they appear in section headers,
	// in per-slot provenance, and in the hash canonicalization) so the
	// empty name is rejected up front.
	ErrMissingSlotName = errors.New("agentcontext: missing slot name")

	// ErrDuplicateSlotName is returned by ContextRequest.Validate when
	// two SlotSpecs share the same Name. Names must be unique within a
	// request so per-slot provenance and rendering output are
	// unambiguous.
	ErrDuplicateSlotName = errors.New("agentcontext: duplicate slot name")

	// ErrUnknownSlotKind is returned by SlotSpec.Validate when
	// SlotSource.Kind is not one of the declared SlotSourceKind values,
	// and by DefaultProvider.Assemble when no Resolver is registered
	// for a slot's declared kind. The wrap value names the offending
	// kind.
	ErrUnknownSlotKind = errors.New("agentcontext: unknown slot kind")

	// ErrUnsafeSlotPath is returned by SlotSpec.Validate when a
	// path-bearing SlotSource (static_file, static_dir) declares a
	// path containing ".." segments. Path-safety is enforced
	// defensively at validation time so consumers cannot trick a
	// resolver into reading outside the intended root. Absolute paths
	// and tilde-prefixed paths are permitted at the contract layer;
	// concrete resolvers MAY tighten this further.
	ErrUnsafeSlotPath = errors.New("agentcontext: unsafe slot path")

	// ErrSlotRequiredAndEmpty is returned by DefaultProvider.Assemble
	// when a SlotSpec marked Required=true resolves to an empty Content
	// AND its resolver returned no error. Resolver-returned errors on
	// required slots are surfaced separately as ErrRequiredSlotFailed.
	ErrSlotRequiredAndEmpty = errors.New("agentcontext: required slot resolved empty")

	// ErrRequiredSlotFailed is returned by DefaultProvider.Assemble
	// when a SlotSpec marked Required=true had its Resolver return a
	// non-nil error. The error is wrapped so callers can use
	// errors.Is(err, ErrRequiredSlotFailed) AND errors.Is(err,
	// resolver-specific-sentinel) to branch on either layer.
	ErrRequiredSlotFailed = errors.New("agentcontext: required slot resolver failed")

	// ErrBudgetExhausted is returned by the default Renderer (and
	// SHOULD be returned by alternative Renderers) when the rendered
	// output would exceed Limits.MaxBytes AND truncation is not
	// permitted. The default policy IS to permit truncation — see
	// LimitsApplied for which slots were dropped or shortened — so
	// this sentinel is reserved for future strict-mode callers.
	ErrBudgetExhausted = errors.New("agentcontext: budget exhausted")

	// ErrMissingResolver is returned by NewProvider when the supplied
	// resolver map is nil. An empty (but non-nil) map is permitted —
	// it merely means Assemble will fail on any non-zero-kind slot
	// with ErrUnknownSlotKind, which is the intended behaviour for
	// test fixtures.
	ErrMissingResolver = errors.New("agentcontext: nil resolver map")

	// ErrMissingRenderer is returned by NewProvider when the supplied
	// Renderer is nil. The package ships DefaultRenderer for the
	// common case; callers must opt out explicitly by passing a
	// non-nil custom Renderer.
	ErrMissingRenderer = errors.New("agentcontext: nil renderer")
)
