// Code generated by "stringer -type=ExitCause"; DO NOT EDIT.

package fuzz

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ExitCauseOK-0]
	_ = x[ExitCasueTimeout-1]
	_ = x[ExitCauseNZEC-2]
}

const _ExitCause_name = "ExitCauseOKExitCasueTimeoutExitCauseNZEC"

var _ExitCause_index = [...]uint8{0, 11, 27, 40}

func (i ExitCause) String() string {
	if i < 0 || i >= ExitCause(len(_ExitCause_index)-1) {
		return "ExitCause(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ExitCause_name[_ExitCause_index[i]:_ExitCause_index[i+1]]
}
