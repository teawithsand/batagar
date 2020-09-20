package fuzz

import (
	"time"

	"github.com/teawithsand/batagar/events"
)


// FuzzerStatus represents current status of some fuzzer.
type FuzzerStatus struct {
	// TODO(teawithsand): move this to protobuf?
}

// Fuzzer represents single fuzzer, which is running and emits events.
type Fuzzer interface {
	// Fuzzer error if any
	Error() error
	Status() FuzzerStatus
	Listeners() events.Listeners
}

//go:generate stringer -type=ExitCause

// ExitCause tells how progarm exited.
type ExitCause int8

const (
	ExitCauseOK      ExitCause = 0
	ExitCasueTimeout ExitCause = 1
	ExitCauseNZEC    ExitCause = 2 // Here: all kinds of panics, ASAN failures, OOM Ram allocations, or regular SEGFAULT
)

// CrasherCheckResult contains result of checking given crasher.
type CrasherCheckResult struct {
	Runtime           time.Duration
	ExitCause         ExitCause
	PeakRAMUsageBytes uint64
}

// CoverageCheckResult contains result of coverage checking.
type CoverageCheckResult struct {
	Runtime   time.Duration
	ExitCause ExitCause // in fact we may have found crasher input

	FoundNewCoverage bool
}

// FuzzerInput is single input for fuzzer, either in memory or in some file.
type FuzzerInput interface{}

// FuzzerInputCollection is collection of inputs for fuzzer.
type FuzzerInputCollection interface{}

// Checker is responsible for checking if:
// 1. Given corpus provides new coverage.
// 2. Given crashers indeed crashes program, causes it to hang or to allocate huge chunk of RAM.
type Checker interface {
	CheckCrasher(crasher FuzzerInput) (CrasherCheckResult, error)
	CheckCoverage(fic1, fic2 FuzzerInputCollection) (CoverageCheckResult, error)
}
