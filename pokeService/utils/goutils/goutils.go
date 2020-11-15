package goutils

import (
	"runtime"

	"github.com/pkg/errors"
)

// GetFrame returns information about the calling frame
// Sourced from here: https://stackoverflow.com/questions/35212985/is-it-possible-get-information-about-caller-function-in-golang
func GetFrame(skipFrames int) runtime.Frame {
	// We need the frame at index skipFrames+2, since we never want runtime.Callers and getFrame
	targetFrameIndex := skipFrames + 2

	// Set size to targetFrameIndex+2 to ensure we have room for one more caller than we need
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)

	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			if frameIndex == targetFrameIndex {
				frame = frameCandidate
			}
		}
	}

	return frame
}

// ErrWrap returns a wrapped error with the function caller as a wrap
func ErrWrap(err error) error {
	return errors.Wrap(err, GetFrame(2).Function)
}
