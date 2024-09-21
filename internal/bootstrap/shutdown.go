package bootstrap

// CleanupFunc Define a type for cleanup functions
type CleanupFunc func()

var cleanupFuncs []CleanupFunc
var shuttingDown bool = false

// RegisterCleanup adds a function to the list of cleanup functions
func RegisterCleanup(fn CleanupFunc) {
	cleanupFuncs = append(cleanupFuncs, fn)
}

// Shutdown executes all registered cleanup functions
func Shutdown() {
	shuttingDown = true
	for _, fn := range cleanupFuncs {
		fn()
	}
}
