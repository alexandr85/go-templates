package interfaces

// Runnable can be implemented runnable components
type Runnable interface {
	Run() error
	Shutdown()
}

// Leadership
type Leadership interface {
	Leader() bool
	SetLeader(value bool)
}

// Readiness typically used for application readiness check
type Readiness interface {
	Ready() bool
}

// Identifiable can be implemented for types uniquely
// identified by a subset of its fields.
type Identifiable interface {
	Unique() interface{}
}
