package profiling

import "github.com/spf13/nitro"

var Timer *nitro.B

func Initialize() {
	Timer = nitro.Initialize()
}