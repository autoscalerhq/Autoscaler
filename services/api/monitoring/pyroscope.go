package monitoring

import (
	"github.com/grafana/pyroscope-go"
	"os"
	"runtime"
)

func SetupProfiling() (*pyroscope.Profiler, error) {
	// These 2 lines are only required if you're using mutex or block profiling
	// Read the explanation below for how to set these rates:
	runtime.SetMutexProfileFraction(5)
	runtime.SetBlockProfileRate(5)

	return pyroscope.Start(pyroscope.Config{
		ApplicationName: "simple.golang.app",

		// replace this with the address of pyroscope server
		ServerAddress: "http://localhost:4040",

		// you can disable logging by setting this to nil
		Logger: nil, //pyroscope.StandardLogger,

		// you can provide static tags via a map:
		Tags: map[string]string{"hostname": os.Getenv("HOSTNAME")},

		ProfileTypes: []pyroscope.ProfileType{
			// these profile types are enabled by default:
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,

			// these profile types are optional:
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})
}
