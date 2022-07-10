package pvi

type PVIEntity struct {
	ID    string
	State EntityState
}

type EntityState int

const (
	Available = iota
	Consumed
	Deleted
	Unknown
)

type Query struct {
	ID    string
	State EntityState
}

type UpdateValues struct {
	State EntityState
}
